package services

import "go-gr-maps/pkg/models"

type IQuestionTemplateRepository interface {
	BulkCreate(questionTemplates []*models.QuestionTemplateDB) error
	FindQuestions() []*models.QuestionTemplateDB
	FindQuestionAnswers(question models.QuestionTemplateDB) []*models.QuestionTemplateDB
	FindQuestionsAnswers(questions []*models.QuestionTemplateDB) []*models.QuestionTemplateDB
}
