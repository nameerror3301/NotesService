package tools

import "golang.org/x/crypto/bcrypt"

// From hashing password
func HashingPassword(pass string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	return bytes, err
}

// From validation password
func CheckControlSum(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
