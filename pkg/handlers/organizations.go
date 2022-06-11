package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go-gr-maps/pkg/models"
	"net/http"
)

type IOrganizationService interface {
	AllOrganizations() []models.OrganizationDB
}

type OrganizationHandler struct {
	db      *sqlx.DB
	service IOrganizationService
}

func NewOrganizationHandler(db *sqlx.DB, service IOrganizationService) *OrganizationHandler {
	return &OrganizationHandler{db: db, service: service}
}

// ListAccounts godoc
// @Accept       json
// @Produce      json
// @Success 200 {object} []models.OrganizationDB
// @Router       /api/organizations/ [get]
func (h *OrganizationHandler) OrganizationsList(c echo.Context) error {
	organizations := h.service.AllOrganizations()
	return c.JSON(http.StatusOK, organizations)

}
