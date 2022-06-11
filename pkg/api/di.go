package api

import (
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/handlers"
	"go-gr-maps/pkg/repositories"
	"go-gr-maps/pkg/services"
)

type Container struct {
	db *sqlx.DB
}

func NewContainer(db *sqlx.DB) *Container {
	return &Container{db: db}
}

func (c *Container) InitOrganizationRepository() services.IOrganizationRepository {
	var repository services.IOrganizationRepository = repositories.NewOrganizationRepository(c.db)
	return repository
}

func (c *Container) InitOrganizationService() handlers.IOrganizationService {
	var service handlers.IOrganizationService = services.NewOrganizationService(
		c.InitOrganizationRepository(),
	)
	return service
}
