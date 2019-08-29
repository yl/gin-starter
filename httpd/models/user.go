package models

import (
	"time"

	"github.com/yangliulnn/gin-starter/httpd/responses"
	"github.com/yangliulnn/gin-starter/httpd/utils"
	"github.com/yangliulnn/gin-starter/services/database"
)

type User struct {
	ID        int   `gorm:"primary_key"`
	Mobile    string `gorm:"unique_index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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

func (u *User) Insert(mobile string, password string) (*User, error) {
	user := &User{Mobile: mobile, Password: password}
	err := database.DB.Save(user).Error
	return user, err
}

func (u *User) All() (*Users, error) {
	users := &Users{}
	err := database.DB.Find(&users).Error
	return users, err
}

func (u *User) FindById(id uint) (*User, error) {
	user := &User{}
	err := database.DB.Find(&user, id).Error
	return user, err
}

func (u *User) FindOne(condition map[string]interface{}) (*User, error) {
	user := &User{}
	err := database.DB.Where(condition).First(user).Error
	return user, err
}
