package demo

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TestBcrypt(t *testing.T) {
	myPwd := "shubham"
	providedHash, _ := HashPassword(myPwd)
	fmt.Println("Password :", myPwd)
	fmt.Println("Hash :", providedHash)

	isMatch := CheckPasswordHash(myPwd, providedHash)
	fmt.Println("Matched ?:", isMatch)
}
