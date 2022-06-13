package fixtures

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/api"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/services"
	"go-gr-maps/pkg/utils"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type CustomFaker struct {
	faker faker.Faker
}

func NewCustomFaker() *CustomFaker {
	return &CustomFaker{faker: faker.NewWithSeed(rand.NewSource(time.Now().UnixNano()))}
}

type Seeder struct {
	db                             *sqlx.DB
	organizationsRepository        services.IOrganizationRepository
	organizationsSectorsRepository services.IOrganizationSectorRepository
	questionTemplateRepository     services.IQuestionTemplateRepository
	organizationFormRepository     services.IOrganizationFormRepository
}

func NewSeeder(db *sqlx.DB) *Seeder {
	diContainer := api.NewContainer(db)

	return &Seeder{
		db:                             db,
		organizationsRepository:        diContainer.InitOrganizationRepository(),
		organizationsSectorsRepository: diContainer.InitOrganizationSectorsRepository(),
		questionTemplateRepository:     diContainer.InitQuestionTemplateRepository(),
		organizationFormRepository:     diContainer.InitOrganizationFormRepository(),
	}
}

func (s Seeder) seedOrganizationsSectors() []*models.OrganizationSectorDB {
	var organizationsSectors []*models.OrganizationSectorDB
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		organizationSector := &models.OrganizationSectorDB{}
		organizationSector.FillDefault()
		fake := NewCustomFaker()
		organizationSector.Name = fake.faker.Lorem().Word()
		organizationsSectors = append(organizationsSectors, organizationSector)
	}
	err := s.organizationsSectorsRepository.BulkCreate(organizationsSectors)
	if err != nil {
		utils.LogError(err)
		return []*models.OrganizationSectorDB{}
	}
	return organizationsSectors
}

func (s Seeder) seedOrganizations() []*models.OrganizationDB {
	var organizations []*models.OrganizationDB
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		organization := &models.OrganizationDB{}
		organization.FillDefault()
		fake := faker.NewWithSeed(rand.NewSource(time.Now().UnixNano()))
		organization.FullName = sql.NullString{String: fake.Company().Name(), Valid: true}
		organization.Name = sql.NullString{String: fake.Company().Name(), Valid: true}
		organization.About = sql.NullString{String: fake.Lorem().Sentence(30), Valid: true}
		organization.Logo = sql.NullString{String: fake.Internet().URL(), Valid: true}
		organization.Color = []string{fake.Color().Hex(), fake.Color().Hex()}
		organization.Url = sql.NullString{String: fake.Internet().URL(), Valid: true}
		organization.EmployeesAmount = rand.Intn(1000)
		organization.Revenue = rand.Intn(10000000)
		inn := ""
		for num := 0; num < 12; num++ {
			rand.Seed(time.Now().UnixNano())
			inn += strconv.Itoa(rand.Intn(10))
		}
		organization.OrganizationInn = sql.NullString{String: inn, Valid: true}
		organization.Phone = sql.NullString{String: fake.Phone().Number(), Valid: true}
		organization.Email = sql.NullString{String: fake.Internet().Email(), Valid: true}
		organization.Status = models.OrganizationDBStatus(
			fake.RandomStringElement([]string{string(models.MODERATION), string(models.ACCEPTED)}),
		)
		organization.GrPoints = rand.Intn(16)
		organizations = append(organizations, organization)
	}
	err := s.organizationsRepository.BulkCreate(organizations)
	if err != nil {
		return organizations
	}
	return organizations
}

func (s Seeder) seedOrganizationsRels(organizations []*models.OrganizationDB, sectors []*models.OrganizationSectorDB) {
	for _, organization := range organizations {
		var orgSectors []*models.OrganizationSectorDB
		for i := 0; i < 3; i++ {
			rand.Seed(time.Now().UnixNano())
			orgSectors = append(orgSectors, sectors[rand.Intn(len(sectors))])
		}
		err := s.organizationsRepository.SetSectors(organization, orgSectors)
		if err != nil {
			return
		}
	}
}

func (s Seeder) seedQuestionTemplates() {
	file, _ := ioutil.ReadFile("pkg/fixtures/questiontemplatedb.json")
	var templates []*models.QuestionTemplateDB
	err := json.Unmarshal(file, &templates)
	if err != nil {
		return
	}

	err = s.questionTemplateRepository.BulkCreate(templates)
	if err != nil {
		return
	}
}

func (s Seeder) seedOrganizationsForms(organizations []*models.OrganizationDB) {
	questionsTemplates := s.questionTemplateRepository.FindQuestions()
	var preparedQuestionsAnswers = make(map[int64][]*models.QuestionTemplateDB, len(questionsTemplates))
	questionsAnswers := s.questionTemplateRepository.FindQuestionsAnswers(questionsTemplates)
	for _, ans := range questionsAnswers {
		preparedQuestionsAnswers[*ans.ParentId] = append(preparedQuestionsAnswers[*ans.ParentId], ans)
	}
	for _, organization := range organizations {
		var organizationForms []*models.OrganizationFormDB
		for _, question := range questionsTemplates {
			fake := NewCustomFaker()
			form := &models.OrganizationFormDB{
				OrganizationId:     organization.ID,
				TemplateQuestionId: question.ID,
				TextAnswer: sql.NullString{
					String: fake.faker.Lorem().Sentence(5),
					Valid:  true,
				},
			}
			organizationForms = append(organizationForms, form)
		}
		err := s.organizationFormRepository.BulkCreate(organizationForms)
		if err != nil {
			return
		}
		for _, form := range organizationForms {
			questionFormAnswers := preparedQuestionsAnswers[form.TemplateQuestionId]
			if len(questionFormAnswers) == 0 {
				utils.LogError(errors.New(fmt.Sprintf("Undefined question answers question_id %d", form.TemplateQuestionId)))
			}
			err = s.organizationFormRepository.SetupQuestionAnswers(
				form, []*models.QuestionTemplateDB{questionFormAnswers[0]},
			)
			if err != nil {
				return
			}
		}
	}
}

func (s Seeder) SeedAll() {
	organizations := s.seedOrganizations()
	organizationSectors := s.seedOrganizationsSectors()
	if len(organizationSectors) != 0 && len(organizations) != 0 {
		s.seedOrganizationsRels(organizations, organizationSectors)
	}
	s.seedQuestionTemplates()
	if len(organizations) != 0 {
		s.seedOrganizationsForms(organizations)
	}
}
