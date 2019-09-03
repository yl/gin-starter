package models

import (
	"time"

	"github.com/yangliulnn/gin-starter/httpd/responses"
	"github.com/yangliulnn/gin-starter/httpd/utils"
	"github.com/yangliulnn/gin-starter/services/database"
)

type User struct {
	ID        int    `gorm:"primary_key"`
	Mobile    string `gorm:"unique_index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func NewUser() *User {
	return &User{}
}

type Users []User

func (u *User) Transformer() responses.Item {
	return &map[string]interface{}{
		"id":         u.ID,
		"mobile":     u.Mobile,
		"created_at": utils.NewTime().Format(u.CreatedAt),
	}
}

func (us *Users) Transformer() responses.Collection {
	collection := make(responses.Collection, len(*us))
	for index, user := range *us {
		collection[index] = user.Transformer()
	}
	return collection
}

func (u *User) Save() error {
	err := database.DB.Save(u).Error
	return err
}

func (u *User) FirstBy(field string, value interface{}) error {
	err := database.DB.Where(field+" = ?", value).First(u).Error
	return err
}
