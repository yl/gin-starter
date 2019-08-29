package helpers

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash[:])
}
