package service

import (
	"context"
	"errors"

	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     auth.Authentication
}

//go:generate mockgen -source=service.go -destination=service_mocks.go -package=service
type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserLogin(ctx context.Context, userData models.NewUser) (string, error)

	AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error)

	AddJobDetails(ctx context.Context, jobData models.Jobs, cid uint64) (models.Jobs, error)
	ViewAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error)
	ViewJobByCid(ctx context.Context, cid uint64) ([]models.Jobs, error)
}

func NewService(userRepo repository.UserRepo, a auth.Authentication) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
	}, nil
}
