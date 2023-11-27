package pkg

import (
	"errors"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password could not be empty")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Info().Msg("error in hashing the password")
		return "", err
	}
	return string(hashedPass), nil

}

func CheckHashedPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
