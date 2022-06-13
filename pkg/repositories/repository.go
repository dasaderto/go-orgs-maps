package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-gr-maps/pkg/models"
	"go-gr-maps/pkg/utils"
	"time"
)

type BaseRepository struct {
	db    *sqlx.DB
	model models.IBaseModel
}

func (r BaseRepository) BulkCreate(query string, models []*models.BaseModel) error {
	tx := r.db.MustBegin()
	for _, item := range models {
		item.CreatedAt = time.Now()
		rows, err := tx.NamedQuery(query, item)
		if err != nil {
			utils.LogError(err)
			return err
		}
		if rows.Next() {
			err = rows.Scan(&item.ID)
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
