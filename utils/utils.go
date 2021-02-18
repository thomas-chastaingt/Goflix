package utils

import "golang.org/x/crypto/bcrypt"

/**************************** crypto *************************/

//HashPassword encrypt the password before saved it in database
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash check if the password given exist in database
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
