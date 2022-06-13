package services

import (
	"go-gr-maps/pkg/models"
)

type IOrganizationRepository interface {
	FindAll() []models.OrganizationDB
	BulkCreate(organizations []*models.OrganizationDB) error
	SetSectors(organization *models.OrganizationDB, organizationSector []*models.OrganizationSectorDB) error
}

type OrganizationService struct {
	repository IOrganizationRepository
}

func NewOrganizationService(repository IOrganizationRepository) *OrganizationService {
	return &OrganizationService{repository: repository}
}

func (s OrganizationService) AllOrganizations() []models.OrganizationDB {
	return s.repository.FindAll()
}
