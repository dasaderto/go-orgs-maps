package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/utils"
	"time"
)

type OrganizationFormRepository struct {
	BaseRepository
}

func NewOrganizationFormRepository(db *sqlx.DB) *OrganizationFormRepository {
	var model models.OrganizationFormDB
	return &OrganizationFormRepository{BaseRepository{
		db:    db,
		model: model,
	}}
}

func (r OrganizationFormRepository) BulkCreate(forms []*models.OrganizationFormDB) error {
	query := `INSERT INTO organization_forms (created_at, updated_at, text_answer, organization_id, question_id) VALUES (:created_at, :updated_at, :text_answer, :organization_id, :question_id) RETURNING id;`
	tx := r.db.MustBegin()
	for _, form := range forms {
		form.CreatedAt = time.Now()
		rows, err := tx.NamedQuery(query, form)
		if err != nil {
			utils.LogError(err)
			return err
		}
		if rows.Next() {
			err = rows.Scan(&form.ID)
			if err != nil {
				utils.LogError(err)
				return err
			}
		}
		err = rows.Close()
		if err != nil {
			utils.LogError(err)
			return err
		}
	}
	err := tx.Commit()
	if err != nil {
		utils.LogError(err)
	}
	return err
}

func (r OrganizationFormRepository) FindQuestionAnswers(question models.OrganizationFormDB) ([]*models.QuestionTemplateDB, error) {
	query := `SELECT *
				FROM organization_question_templates
				WHERE id IN (
						SELECT answer_id
						FROM organization_forms_answers
						WHERE organization_form_id = $1
						)`
	var answers []*models.QuestionTemplateDB
	err := r.db.Select(&answers, query, question.ID)
	if err != nil {
		utils.LogError(err)
		return []*models.QuestionTemplateDB{}, err
	}
	return answers, nil
}

func (r OrganizationFormRepository) RemoveOrganizationQuestionAnswers(questionId int64, answersIds []int64) error {
	if len(answersIds) == 0 {
		return nil
	}
	query := `DELETE FROM organization_forms_answers
				WHERE organization_form_id = $1 AND organization_forms_answers.answer_id IN ($2);`
	query, args, err := sqlx.In(query, questionId, answersIds)
	if err != nil {
		utils.LogError(err)
		return err
	}
	query = r.db.Rebind(query)
	_, err = r.db.Exec(query, args...)
	if err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}

func (r OrganizationFormRepository) SetupQuestionAnswers(
	form *models.OrganizationFormDB, answers []*models.QuestionTemplateDB,
) error {
	questionExistsAnswers, err := r.FindQuestionAnswers(*form)
	if err != nil {
		return err
	}
	var existsAnswersIds []int64
	for _, s := range questionExistsAnswers {
		existsAnswersIds = append(existsAnswersIds, s.ID)
	}
	var answersIds []int64
	for _, answer := range answers {
		answersIds = append(answersIds, answer.ID)
	}
	answersToDeleteIDs := utils.Difference(existsAnswersIds, answersIds)
	err = r.RemoveOrganizationQuestionAnswers(form.ID, answersToDeleteIDs)
	if err != nil {
		return err
	}
	var formAnswers []*models.OrganizationFormsAnswersDB
	for _, answer := range utils.UnifySlice(answers) {
		if utils.Contains(existsAnswersIds, answer.ID) {
			continue
		}
		formAnswers = append(formAnswers, &models.OrganizationFormsAnswersDB{
			OrganizationFormID: form.ID,
			AnswerID:           answer.ID,
		})
	}
	query := `INSERT INTO organization_forms_answers (organization_form_id, answer_id) VALUES (:organization_form_id, :answer_id)`
	_, err = r.db.NamedExec(query, formAnswers)
	if err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}
