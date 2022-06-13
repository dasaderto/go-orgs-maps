package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/utils"
	"time"
)

type QuestionTemplateRepository struct {
	BaseRepository
}

func NewQuestionTemplateRepository(db *sqlx.DB) *QuestionTemplateRepository {
	var model models.QuestionTemplateDB
	return &QuestionTemplateRepository{BaseRepository{db: db, model: &model}}
}

func (r QuestionTemplateRepository) BulkCreate(questionTemplates []*models.QuestionTemplateDB) error {
	query := `INSERT INTO organization_question_templates (created_at, updated_at, text, answer_type, is_answer, is_required, is_private, opened_answer_id, parent_id, gr_points) 
VALUES (:created_at, :updated_at, :text, :answer_type, :is_answer, :is_required, :is_private, :opened_answer_id, :parent_id, :gr_points) RETURNING id`
	tx := r.db.MustBegin()
	for _, template := range questionTemplates {
		template.CreatedAt = time.Now()
		rows, err := tx.NamedQuery(query, template)
		if err != nil {
			utils.LogError(err)
			return err
		}
		if rows.Next() {
			err = rows.Scan(&template.ID)
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
		return err
	}
	return nil
}

func (r QuestionTemplateRepository) FindQuestions() []*models.QuestionTemplateDB {
	var questionTemplates []*models.QuestionTemplateDB
	query := `SELECT * FROM organization_question_templates WHERE is_answer=false`
	err := r.db.Select(&questionTemplates, query)
	if err != nil {
		utils.LogError(err)
		return []*models.QuestionTemplateDB{}
	}
	return questionTemplates
}

func (r QuestionTemplateRepository) FindQuestionAnswers(question models.QuestionTemplateDB) []*models.QuestionTemplateDB {
	var answersTemplates []*models.QuestionTemplateDB
	query := `SELECT * FROM organization_question_templates WHERE parent_id=$1`
	err := r.db.Select(&answersTemplates, query, question.ID)
	if err != nil {
		utils.LogError(err)
		return []*models.QuestionTemplateDB{}
	}
	return answersTemplates
}

func (r QuestionTemplateRepository) FindQuestionsAnswers(questions []*models.QuestionTemplateDB) []*models.QuestionTemplateDB {
	var defaultReturn []*models.QuestionTemplateDB
	if len(questions) == 0 {
		return defaultReturn
	}
	var answersTemplates []*models.QuestionTemplateDB
	var questionsIds []int64
	for _, q := range questions {
		questionsIds = append(questionsIds, q.ID)
	}
	query := `SELECT * FROM organization_question_templates WHERE parent_id IN (?);`
	query, args, err := sqlx.In(query, questionsIds)
	if err != nil {
		utils.LogError(err)
		return defaultReturn
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&answersTemplates, query, args...)
	if err != nil {
		utils.LogError(err)
		return defaultReturn
	}
	return answersTemplates
}
