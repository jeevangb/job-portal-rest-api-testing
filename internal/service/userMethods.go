package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/pkg"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) ResetPassword(ctx context.Context, resetData models.ResetPassword) error {
	if resetData.NewPassword != resetData.ConfirmPassword {
		return errors.New("new and confirm password is not match")
	}

	otp, err := s.rdb.CheckCacheOtp(ctx, resetData.Email)
	if err != nil {
		return errors.New("otp in cache not there")
	}
	if otp != resetData.Otp {
		return errors.New("otp mismatch")
	}
	HashPassword, err := pkg.HashPassword(resetData.ConfirmPassword)
	if err != nil {
		return errors.New("failed to hash the password")
	}

	s.UserRepo.UpdatePassword(ctx, resetData.Email, HashPassword)
	return nil
}

func (s *Service) CheckUserDataAndSendOtp(ctx context.Context, userData models.ForgotPasswod) error {
	resetData, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return err
	}
	if resetData.Dob != userData.Dob {
		return errors.New("date of birth not exist")
	}

	otpData, err := pkg.GenerateOneTimePassword(userData.Email)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate otp")
	}
	err = s.rdb.AddToCacheRedis(ctx, userData.Email, otpData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UserLogin(ctx context.Context, userData models.NewUser) (string, error) {
	// checcking the email in the db
	var userDetails models.User
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return "", err
	}

	// comaparing the password and hashed password
	// err = pkg.CheckHashedPassword(userData.Password, userDetails.PasswordHash)
	// if err != nil {
	// 	log.Info().Err(err).Send()
	// 	return "", errors.New("entered password is wrong")
	// }

	// setting up the claims
	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token, err := s.auth.GenerateAuthToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *Service) UserSignup(ctx context.Context, userData models.NewUser) (models.User, error) {
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userDetails := models.User{
		Username:     userData.Username,
		Email:        userData.Email,
		PasswordHash: hashedPass,
		Dob:          userData.Dob,
	}
	userDetails, err = s.UserRepo.CreateUser(ctx, userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}
