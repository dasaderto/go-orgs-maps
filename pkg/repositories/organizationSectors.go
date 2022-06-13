package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/utils"
	"time"
)

type OrganizationSectorsRepository struct {
	BaseRepository
}

func NewOrganizationSectorsRepository(db *sqlx.DB) *OrganizationSectorsRepository {
	var model models.OrganizationSectorDB
	return &OrganizationSectorsRepository{BaseRepository{db: db, model: &model}}
}

func (r OrganizationSectorsRepository) BulkCreate(organizationsSectors []*models.OrganizationSectorDB) error {
	query := `INSERT INTO organization_sectors (created_at, updated_at, name) VALUES (:created_at, :updated_at, :name) RETURNING id`
	tx := r.db.MustBegin()
	for _, sector := range organizationsSectors {
		sector.CreatedAt = time.Now()
		rows, err := tx.NamedQuery(query, sector)
		if err != nil {
			utils.LogError(err)
			return err
		}
		if rows.Next() {
			err = rows.Scan(&sector.ID)
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
