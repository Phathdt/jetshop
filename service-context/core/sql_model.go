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

func (m *SQLModel) FullFill() {
	t := time.Now()

	if m.UpdatedAt == nil {
		m.UpdatedAt = &t
	}
}

func NewUpsertSQLModel(id int) *SQLModel {
	t := time.Now()

	return &SQLModel{
		Id:        id,
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}

func NewUpsertWithoutIdSQLModel() *SQLModel {
	t := time.Now()

	return &SQLModel{
		CreatedAt: &t,
		UpdatedAt: &t,
	}
}
