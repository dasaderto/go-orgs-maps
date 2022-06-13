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
	return repositories.NewOrganizationRepository(c.db)
}

func (c *Container) InitOrganizationSectorsRepository() services.IOrganizationSectorRepository {
	return repositories.NewOrganizationSectorsRepository(c.db)
}

func (c *Container) InitQuestionTemplateRepository() services.IQuestionTemplateRepository {
	return repositories.NewQuestionTemplateRepository(c.db)
}

func (c *Container) InitOrganizationFormRepository() services.IOrganizationFormRepository {
	return repositories.NewOrganizationFormRepository(c.db)
}

func (c *Container) InitOrganizationService() handlers.IOrganizationService {
	return services.NewOrganizationService(
		c.InitOrganizationRepository(),
	)
}
