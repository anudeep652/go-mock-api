package api

import (
	"log"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashedP, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedP)
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GenerateApiKey() string {

	k := make([]byte, API_KEY_LENGTH)
	for i, _ := range k {
		k[i] = CHARSET[rand.Intn(len(CHARSET))]

	}
	return string(k)
}
