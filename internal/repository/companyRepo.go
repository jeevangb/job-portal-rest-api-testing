package repository

import (
	"context"
	"errors"

	"jeevan/jobportal/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateCompany(ctx context.Context, companyData models.Company) (models.Company, error) {
	result := r.DB.Create(&companyData)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to create company data into database")
		return models.Company{}, errors.New("failure in creat comapany details")
	}
	return companyData, nil
}

func (r *Repo) ViewCompanies(ctx context.Context) ([]models.Company, error) {
	var userDetails []models.Company
	result := r.DB.Find(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Msg("failure to find company data")
		return nil, errors.New("company data not found")
	}
	return userDetails, nil
}

func (r *Repo) ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error) {
	var companyData models.Company
	result := r.DB.Where("id = ?", cid).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Msg("failure to find all company data")
		return models.Company{}, errors.New("companies data not found")
	}
	return companyData, nil
}
