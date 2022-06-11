package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"go-gr-maps/pkg/handlers"
)

type Router struct {
	Router      *echo.Echo
	db          *sqlx.DB
	diContainer *Container
}

func NewRouter(db *sqlx.DB) *Router {
	diContainer := NewContainer(db)
	return &Router{
		Router:      echo.New(),
		db:          db,
		diContainer: diContainer,
	}
}

func (r *Router) InitRouter() {
	apiGroup := r.Router.Group("/api")
	r.Router.GET("/swagger/*", echoSwagger.WrapHandler)
	r.InitOrgsRouter(apiGroup)
}

func (r *Router) InitOrgsRouter(routerGroup *echo.Group) {
	organizationsHandler := handlers.NewOrganizationHandler(
		r.db,
		r.diContainer.InitOrganizationService(),
	)
	routerGroup.GET("/organizations/", organizationsHandler.OrganizationsList)
}
