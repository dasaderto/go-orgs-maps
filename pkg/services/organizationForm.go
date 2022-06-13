package services

import "go-gr-maps/pkg/models"

type IOrganizationFormRepository interface {
	BulkCreate(forms []*models.OrganizationFormDB) error
	FindQuestionAnswers(question models.OrganizationFormDB) ([]*models.QuestionTemplateDB, error)
	SetupQuestionAnswers(form *models.OrganizationFormDB, answers []*models.QuestionTemplateDB) error
}
