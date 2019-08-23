package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash[:]), err
}

func (p *Password) Check(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
