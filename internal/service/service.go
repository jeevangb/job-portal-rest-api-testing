package service

import (
	"context"
	"errors"

	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/cache"
	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     auth.Authentication
	rdb      cache.Caching
}

//go:generate mockgen -source=service.go -destination=service_mocks.go -package=service
type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserLogin(ctx context.Context, userData models.NewUser) (string, error)

	AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error)

	AddJobDetails(ctx context.Context, jobData models.Hr, cid uint64) (models.ResponseJobId, error)
	FilterJob(ctx context.Context, jobApplication []models.RespondJobApplicant) ([]models.RespondJobApplicant, error)
	ViewAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error)
	ViewJobByCid(ctx context.Context, cid uint64) ([]models.Jobs, error)

	CheckUserDataAndSendOtp(ctx context.Context, userData models.ForgotPasswod) error
}

func NewService(userRepo repository.UserRepo, a auth.Authentication, redis cache.Caching) (UserService, error) {

	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
		rdb:      redis,
	}, nil
}
