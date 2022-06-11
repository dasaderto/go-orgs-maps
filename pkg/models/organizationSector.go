package models

type OrganizationSectorDB struct {
	BaseModel
	Name string `db:"name" json:"name"`
}

func (m OrganizationSectorDB) TableName() string {
	return "organization_sectors"
}
