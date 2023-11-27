package repository

import (
	"context"
	"errors"

	"jeevan/jobportal/internal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Jobs, error) {
	var jobData models.Jobs
	result := r.DB.Preload("JobLocation").
		Preload("Technology").
		Preload("WorkMode").
		Preload("Qualification").
		Preload("Shift").
		Preload("JobType").
		Where("id = ?", jid).Find(&jobData)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to find job details in database")
		return models.Jobs{}, errors.New("job data not found")
	}
	return jobData, nil
}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Jobs) (models.ResponseJobId, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to create job details in database")
		return models.ResponseJobId{}, errors.New("could not create the jobs")
	}
	return models.ResponseJobId{
		ID: jobData.ID,
	}, nil
}

func (r *Repo) FindAllJobs(ctx context.Context) ([]models.Jobs, error) {
	var jobDatas []models.Jobs
	result := r.DB.Find(&jobDatas)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to find all the jobs in database")
		return nil, errors.New("jobs data not found")
	}
	return jobDatas, nil
}

func (r *Repo) FindJob(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	var jobData []models.Jobs
	result := r.DB.Where("cid = ?", cid).Find(&jobData)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("failed to find job detail in database")
		return nil, errors.New("job data not found")
	}
	return jobData, nil
}
