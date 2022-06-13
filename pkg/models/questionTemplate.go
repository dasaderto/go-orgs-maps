package models

type OrganizationFormAnswerType string

const (
	TEXT     = "TEXT"
	NUMBER   = "NUMBER"
	RADIO    = "RADIO"
	CHECKBOX = "CHECKBOX"
)

type QuestionTemplateDB struct {
	BaseModel
	Text           string  `db:"text" json:"text"`
	ParentId       *int64  `db:"parent_id" json:"parent_id"`
	GrPoints       float64 `db:"gr_points" json:"gr_points"`
	OpenedAnswerId *int64  `db:"opened_answer_id" json:"opened_answer_id"`
	AnswerType     string  `db:"answer_type" json:"answer_type"`
	IsAnswer       bool    `db:"is_answer" json:"is_answer"`
	IsRequired     bool    `db:"is_required" json:"is_required"`
	IsPrivate      bool    `db:"is_private" json:"is_private"`
}

func (m QuestionTemplateDB) FillDefault() {
}

func (m QuestionTemplateDB) TableName() string {
	return "organization_question_templates"
}
