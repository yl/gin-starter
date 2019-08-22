package repositories

import (
	"github.com/jinzhu/gorm"
	"go-trading/models"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *models.User) *models.User {
	err := r.db.Save(user).Error
	if err != nil {
		log.Fatalln(err)
	}

	return user
}
