package repository

import (
	"context"
	"errors"

	"jeevan/jobportal/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) UpdatePassword(ctx context.Context, email string, resetPassword string) error {
	// var userDetails models.User
	result := r.DB.Model(&models.User{}).Where("email=?", email).Update("PasswordHash", "resetPassword")
	if result.Error != nil {
		// Handle the error
		log.Error().Err(result.Error).Msg("failed to update password in database")
		return errors.New("password reset failed")
	} else {
		// Update successful
		log.Info().Msg("password reset successfully")
	}
	return nil
}

func (r *Repo) CreateUser(ctx context.Context, UserDetails models.User) (models.User, error) {
	result := r.DB.Create(&UserDetails)
	if result.Error != nil {
		//handle error
		log.Error().Err(result.Error).Msg("failed to user data in database")
		return models.User{}, errors.New("login falied")
	}
	return UserDetails, nil
}

func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	var userDetails models.User
	result := r.DB.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to check the email in database")
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil

}
