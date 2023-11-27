package service

import (
	"context"
	"errors"
	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/pkg"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) ResetPassword(ctx context.Context, resetData models.ResetPassword) error {
	if resetData.NewPassword != resetData.ConfirmPassword {
		log.Info().Msg("failed to comapre password")
		return errors.New("password does not match")
	}
	otp, err := s.rdb.CheckCacheOtp(ctx, resetData.Email)
	if err != nil {
		return err
	}
	if otp != resetData.Otp {
		log.Info().Msg("failed to compare the otp")
		return errors.New("otp mismatch")
	}
	HashPassword, err := pkg.HashPassword(resetData.ConfirmPassword)
	if err != nil {
		return err
	}
	err = s.UserRepo.UpdatePassword(ctx, resetData.Email, HashPassword)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CheckUserDataAndSendOtp(ctx context.Context, userData models.ForgotPasswod) error {
	resetData, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return err
	}
	if resetData.Dob != userData.Dob {
		log.Info().Msg("failed to compare date of birth")
		return errors.New("invalid dob")
	}
	otpData, err := pkg.GenerateOneTimePassword(userData.Email)
	if err != nil {
		return err
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
