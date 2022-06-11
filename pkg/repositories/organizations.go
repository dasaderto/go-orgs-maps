package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go-gr-maps/pkg/models"
	"time"
)

type OrganizationRepository struct {
	db    *sqlx.DB
	model models.IBaseModel
}

func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	var model models.OrganizationDB
	return &OrganizationRepository{db: db, model: &model}
}

func (r OrganizationRepository) FindAll() []models.OrganizationDB {
	var organizations []models.OrganizationDB
	err := r.db.Select(
		&organizations,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", r.model.TableName()),
	)
	if err != nil {
		logrus.Error(err.Error())
		return []models.OrganizationDB{}
	}
	return organizations
}

func (r OrganizationRepository) BulkCreate(organizations []*models.OrganizationDB) error {
	tx := r.db.MustBegin()
	query := `
INSERT INTO organizations (created_at, updated_at, name, full_name, about, logo, color, url, employees_amount, revenue, organization_inn, phone, email, status, gr_points) 
VALUES (:created_at, :updated_at, :name, :full_name, :about, :logo, :color, :url, :employees_amount, :revenue, :organization_inn, :phone, :email, :status, :gr_points) 
RETURNING id`
	for _, organization := range organizations {
		organization.CreatedAt = time.Now()
		rows, err := tx.NamedQuery(query, organization)
		if err != nil {
			return err
		}
		if rows.Next() {
			err := rows.Scan(&organization.ID)
			if err != nil {
				return err
			}
		}
		err = rows.Close()
		if err != nil {
			return err
		}
	}
	err := tx.Commit()
	return err
}
