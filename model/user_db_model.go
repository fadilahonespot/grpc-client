package model

type UserDB struct {
	Id       string `gorm:"primary_key"`
	Name     string
	Email    string
	Alamat   string
	Password string
}

func (e *UserDB) TableName() string {
	return "user"
}