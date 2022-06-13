package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/utils"
	"time"
)

type OrganizationRepository struct {
	BaseRepository
}

func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	var model models.OrganizationDB
	return &OrganizationRepository{BaseRepository{db: db, model: &model}}
}

func (r OrganizationRepository) FindAll() []models.OrganizationDB {
	var organizations []models.OrganizationDB
	err := r.db.Select(
		&organizations,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id ASC", r.model.TableName()),
	)
	if err != nil {
		utils.LogError(err)
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
			utils.LogError(err)
			return err
		}
		if rows.Next() {
			err = rows.Scan(&organization.ID)
			if err != nil {
				utils.LogError(err)
				return err
			}
		}
		err = rows.Close()
		if err != nil {
			utils.LogError(err)
			return err
		}
	}
	err := tx.Commit()
	if err != nil {
		utils.LogError(err)
	}
	return err
}

func (r OrganizationRepository) FindSectors(organizationID int64) ([]*models.OrganizationSectorDB, error) {
	var sectors []*models.OrganizationSectorDB
	err := r.db.Select(
		&sectors,
		`SELECT * FROM organization_sectors WHERE id IN (SELECT sector_id from organization_sectors_rel WHERE organization_id=$1)`,
		organizationID,
	)
	if err != nil {
		utils.LogError(err)
	}
	return sectors, err
}

func (r OrganizationRepository) RemoveSectors(organizationId int64, sectorsIds []int64) error {
	if len(sectorsIds) == 0 {
		return nil
	}
	query := `DELETE FROM organization_sectors_rel where organization_id=$1 and sector_id IN ($2)`
	query, args, err := sqlx.In(query, organizationId, sectorsIds)
	if err != nil {
		utils.LogError(err)
		return err
	}
	query = r.db.Rebind(query)
	_, err = r.db.Exec(query, args...)
	if err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}

func (r OrganizationRepository) SetSectors(organization *models.OrganizationDB, organizationSectors []*models.OrganizationSectorDB) error {
	organizationExistsSectors, err := r.FindSectors(organization.ID)
	if err != nil {
		return err
	}
	var existsSectorsIds []int64
	for _, s := range organizationExistsSectors {
		existsSectorsIds = append(existsSectorsIds, s.ID)
	}
	var sectorsIds []int64
	for _, sector := range organizationSectors {
		sectorsIds = append(sectorsIds, sector.ID)
	}
	sectorsToDeleteIDs := utils.Difference(existsSectorsIds, sectorsIds)
	err = r.RemoveSectors(organization.ID, sectorsToDeleteIDs)
	if err != nil {
		return err
	}
	var orgSectorsRelations []*models.OrganizationSectorsRelDB
	for _, sector := range utils.UnifySlice(organizationSectors) {
		if utils.Contains(existsSectorsIds, sector.ID) {
			continue
		}
		orgSectorsRelations = append(orgSectorsRelations, &models.OrganizationSectorsRelDB{
			OrganizationID: organization.ID,
			SectorID:       sector.ID,
		})
	}
	query := `INSERT INTO organization_sectors_rel (organization_id, sector_id) VALUES (:organization_id, :sector_id)`
	_, err = r.db.NamedExec(query, orgSectorsRelations)
	if err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}
