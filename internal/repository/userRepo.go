package repository

import (
	"context"
	"errors"
	"fmt"

	"jeevan/jobportal/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) UpdatePassword(ctx context.Context, email string, resetPassword string) error {
	// var userDetails models.User
	result := r.DB.Model(&models.User{}).Where("email=?", email).Update("PasswordHash", "resetPassword")
	if result.Error != nil {
		// Handle the error
		fmt.Println("Error updating password:", result.Error)
	} else {
		// Update successful
		fmt.Println("Password updated successfully")
	}
	return nil
}

func (r *Repo) CreateUser(ctx context.Context, UserDetails models.User) (models.User, error) {
	result := r.DB.Create(&UserDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("could not create the user")
	}
	return UserDetails, nil
}

func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	var userDetails models.User
	result := r.DB.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil

}
