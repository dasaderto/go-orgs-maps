package models

type OrganizationFormAnswersQuestionTemplateDB struct {
	AnswerId int
}

type OrganizationFormDB struct {
	BaseModel
	OrganizationId int    `db:"organization_id" json:"organization_id"`
	QuestionId     int    `db:"question_id" json:"question_id"`
	TextAnswer     string `db:"text_answer" json:"text_answer"`
}

func (m OrganizationFormDB) TableName() string {
	return "organization_forms"
}

type OrganizationFormsAnswers struct {
	OrganizationFormID     int `db:"organization_form_id" json:"organization_form_id"`
	OrganizationTemplateID int `db:"organization_template_id" json:"organization_template_id"`
}

func (m OrganizationFormsAnswers) TableName() string {
	return "organization_forms_answers"
}
