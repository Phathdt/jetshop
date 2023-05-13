package core

// SimpleUser only contains public infos
type SimpleUser struct {
	SQLModel
	LastName  string `json:"last_name" gorm:"column:last_name;" db:"last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name;" db:"first_name"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func NewSimpleUser(id int, firstName, lastName string) SimpleUser {
	return SimpleUser{
		SQLModel:  SQLModel{Id: id},
		LastName:  lastName,
		FirstName: firstName,
	}
}
