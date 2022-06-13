package services

import "go-gr-maps/pkg/models"

type IOrganizationSectorRepository interface {
	BulkCreate(organizationsSectors []*models.OrganizationSectorDB) error
}
