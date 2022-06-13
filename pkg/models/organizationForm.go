package models

import "database/sql"

type OrganizationFormDB struct {
	BaseModel
	OrganizationId     int64          `db:"organization_id" json:"organization_id"`
	TemplateQuestionId int64          `db:"question_id" json:"question_id"`
	TextAnswer         sql.NullString `db:"text_answer" json:"text_answer"`
}

func (m OrganizationFormDB) FillDefault() {
}

func (m OrganizationFormDB) TableName() string {
	return "organization_forms"
}

type OrganizationFormsAnswersDB struct {
	OrganizationFormID int64 `db:"organization_form_id" json:"organization_form_id"`
	AnswerID           int64 `db:"answer_id" json:"answer_id"`
}

func (m OrganizationFormsAnswersDB) TableName() string {
	return "organization_forms_answers"
}
