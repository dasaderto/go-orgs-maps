package models

import (
	"database/sql"
	"github.com/lib/pq"
)

type OrganizationDBStatus string

const (
	MODERATION OrganizationDBStatus = "MODERATION"
	ACCEPTED   OrganizationDBStatus = "ACCEPTED"
)

var DEFAULT_ORGANIZATION_COLOR = []string{"#000000", "#000000"}

type OrganizationDB struct {
	BaseModel
	Name            sql.NullString       `db:"name" json:"name"`
	FullName        sql.NullString       `db:"full_name" json:"full_name"`
	About           sql.NullString       `db:"about" json:"about"`
	Logo            sql.NullString       `db:"logo" json:"logo"`
	Color           pq.StringArray       `db:"color" json:"color"`
	Url             sql.NullString       `db:"url" json:"url"`
	EmployeesAmount int                  `db:"employees_amount" json:"employees_amount"`
	Revenue         int                  `db:"revenue" json:"revenue"`
	OrganizationInn sql.NullString       `db:"organization_inn" json:"organization_inn"`
	Phone           sql.NullString       `db:"phone" json:"phone"`
	Email           sql.NullString       `db:"email" json:"email"`
	Status          OrganizationDBStatus `db:"status" json:"status"`
	GrPoints        int                  `db:"gr_points" json:"gr_points"`
}

type OrganizationSectorsRelDB struct {
	OrganizationID int64 `db:"organization_id"`
	SectorID       int64 `db:"sector_id"`
}

func (m OrganizationDB) TableName() string {
	return "organizations"
}

func (m *OrganizationDB) FillDefault() {

	if len(m.Color) == 0 {
		m.Color = DEFAULT_ORGANIZATION_COLOR
	}

	if m.Status != ACCEPTED {
		m.Status = MODERATION
	}
}

func (m OrganizationSectorsRelDB) TableName() string {
	return "organization_sectors_rel"
}
