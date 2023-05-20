package core

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;" db:"id"`
	CreatedAt *time.Time `json:"inserted_at,omitempty" gorm:"column:inserted_at;"  db:"inserted_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"  db:"updated_at"`
}

func NewSQLModel() SQLModel {
	now := time.Now().UTC()

	return SQLModel{
		Id:        0,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
