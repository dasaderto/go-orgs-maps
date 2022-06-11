package pkg

import (
	"database/sql"
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/api"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/services"
	"math/rand"
	"strconv"
	"time"
)

type Seeder struct {
	db                      *sqlx.DB
	organizationsRepository services.IOrganizationRepository
}

func NewSeeder(db *sqlx.DB) *Seeder {
	diContainer := api.NewContainer(db)

	return &Seeder{
		db:                      db,
		organizationsRepository: diContainer.InitOrganizationRepository(),
	}
}

func (s Seeder) SeedOrganizations() {
	var organizations []*models.OrganizationDB
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		organization := &models.OrganizationDB{}
		organization.FillDefault()
		orgsFaker := faker.NewWithSeed(rand.NewSource(time.Now().UnixNano()))
		organization.FullName = sql.NullString{String: orgsFaker.Company().Name(), Valid: true}
		organization.Name = sql.NullString{String: orgsFaker.Company().Name(), Valid: true}
		organization.About = sql.NullString{String: orgsFaker.Lorem().Sentence(30), Valid: true}
		organization.Logo = sql.NullString{String: orgsFaker.Internet().URL(), Valid: true}
		organization.Color = []string{orgsFaker.Color().Hex(), orgsFaker.Color().Hex()}
		organization.Url = sql.NullString{String: orgsFaker.Internet().URL(), Valid: true}
		organization.EmployeesAmount = rand.Intn(1000)
		organization.Revenue = rand.Intn(10000000)
		inn := ""
		for num := 0; num < 12; num++ {
			rand.Seed(time.Now().UnixNano())
			inn += strconv.Itoa(rand.Intn(10))
		}
		organization.OrganizationInn = sql.NullString{String: inn, Valid: true}
		organization.Phone = sql.NullString{String: orgsFaker.Phone().Number(), Valid: true}
		organization.Email = sql.NullString{String: orgsFaker.Internet().Email(), Valid: true}
		organization.Status = models.OrganizationDBStatus(
			orgsFaker.RandomStringElement([]string{string(models.MODERATION), string(models.ACCEPTED)}),
		)
		organization.GrPoints = rand.Intn(16)
		organizations = append(organizations, organization)
	}
	err := s.organizationsRepository.BulkCreate(organizations)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, organization := range organizations {
		fmt.Println(organization.ID)
	}
}

func (s Seeder) SeedAll() {
	s.SeedOrganizations()
}
