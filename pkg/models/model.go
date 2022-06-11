package models

import "time"

type IBaseModel interface {
	TableName() string
	FillDefault()
}

type BaseModel struct {
	ID        int64      `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
