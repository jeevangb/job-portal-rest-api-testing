package service

import (
	"context"
	"encoding/json"
	"fmt"
	"jeevan/jobportal/internal/models"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {
	jobData, err := s.UserRepo.ViewJobDetailsBy(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil
}

func (s *Service) AddJobDetails(ctx context.Context, jobData models.Hr, cid uint64) (models.ResponseJobId, error) {
	createjobDetails := models.Jobs{
		Cid:             uint(cid),
		Title:           jobData.Title,
		MinNoticePeriod: jobData.Minnp,
		MaxNoticePeriod: jobData.Maxnp,
		Budget:          jobData.Budget,
		Jobdescription:  jobData.JobDescription,
	}
	for _, v := range jobData.JobLocation {
		jobLoc := models.JobLocation{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.JobLocation = append(createjobDetails.JobLocation, jobLoc)
	}
	for _, v := range jobData.Technology {
		jobTech := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.Technology = append(createjobDetails.Technology, jobTech)
	}
	for _, v := range jobData.WorkMode {
		jobWorkMode := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.WorkMode = append(createjobDetails.WorkMode, jobWorkMode)
	}
	for _, v := range jobData.Qualification {
		jobWorkQualification := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.Qualification = append(createjobDetails.Qualification, jobWorkQualification)
	}
	for _, v := range jobData.Shift {
		jobShift := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.Shift = append(createjobDetails.Shift, jobShift)
	}
	for _, v := range jobData.JobType {
		jobtype := models.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.JobType = append(createjobDetails.JobType, jobtype)
	}
	job, err := s.UserRepo.CreateJob(ctx, createjobDetails)
	if err != nil {
		return models.ResponseJobId{}, err
	}
	return job, nil
}

func (s *Service) ViewJobByCid(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FindJob(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}

// var mapCaching = make(map[uint]models.Jobs)

// func (s *Service) FilterJob(ctx context.Context, jobApplications []models.RespondJobApplicant) ([]models.RespondJobApplicant, error) {

// 	var FilterJobData []models.RespondJobApplicant

// 	ch := make(chan models.RespondJobApplicant)
// 	wg := new(sync.WaitGroup)

// 	for _, jobApplication := range jobApplications {
// 		wg.Add(1)
// 		go func(jobApplication models.RespondJobApplicant) {
// 			defer wg.Done()
// 			jobdetail, ok := mapCaching[jobApplication.Jid]
// 			if !ok {
// 				val, err := s.UserRepo.ViewJobDetailsBy(ctx, uint64(jobApplication.Jid))
// 				if err != nil {
// 					return
// 				}
// 				mapCaching[jobApplication.Jid] = val
// 				jobdetail = val
// 			}

// 			b := checkApplicantsCriteria(jobApplication, jobdetail)
// 			if b {

// 				ch <- jobApplication
// 			}

// 		}(jobApplication)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for data := range ch {
// 		FilterJobData = append(FilterJobData, data)
// 	}

//		return FilterJobData, nil
//	}
func (s *Service) FilterJob(ctx context.Context, jobApplications []models.RespondJobApplicant) ([]models.RespondJobApplicant, error) {
	var FilterJobData []models.RespondJobApplicant
	ch := make(chan models.RespondJobApplicant)
	wg := new(sync.WaitGroup)
	for _, jobApplication := range jobApplications {
		wg.Add(1)
		go func(jobApplication models.RespondJobApplicant) {
			defer wg.Done()
			var jobdata models.Jobs
			data, err := s.rdb.GetCahceData(ctx, jobApplication.Jid)
			if err == redis.Nil {
				databasedata, err := s.UserRepo.ViewJobDetailsBy(ctx, uint64(jobApplication.Jid))
				if err != nil {
					return
				}
				err = s.rdb.AddToCache(ctx, jobApplication.Jid, databasedata)
				if err != nil {
					return
				}
			} else {

				err = json.Unmarshal([]byte(data), &jobdata)
				if err != nil {
					log.Info().Err(err).Msg("failed to unmarshal the jobdata")
					return
				}
			}
			// Deserialize job details
			b := checkApplicantsCriteria(jobApplication, jobdata)
			if b {
				ch <- jobApplication
			}
		}(jobApplication)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for data := range ch {
		FilterJobData = append(FilterJobData, data)
	}
	return FilterJobData, nil
}

func checkApplicantsCriteria(v models.RespondJobApplicant, jobdetail models.Jobs) bool {
	applicantBudget, err := strconv.Atoi(v.Jobs.Budget)
	if err != nil {
		return false
	}
	hrBudget, err := strconv.Atoi(jobdetail.Budget)
	if err != nil {
		panic("error while conversion budget data from hr posting")
	}
	if applicantBudget > hrBudget {
		return false
	}
	hrMaxNoticePeriod, err := strconv.Atoi(jobdetail.MaxNoticePeriod)
	if err != nil {
		return false
	}
	applicantNoticePeriod, err := strconv.Atoi(v.Jobs.Np)
	if err != nil {
		return false
	}
	if (applicantNoticePeriod < 0) || (applicantNoticePeriod > hrMaxNoticePeriod) {
		return false
	}
	count := 0
	for _, v1 := range v.Jobs.JobLocation {
		count = 0
		for _, v2 := range jobdetail.JobLocation {
			if v1 == v2.ID {
				fmt.Println(v1, v2)
				count++
			}
		}
	}
	if count == 0 {
		return false
	}
	count = 0
	for _, v1 := range v.Jobs.JobType {
		count = 0
		for _, v2 := range jobdetail.JobType {
			if v1 == v2.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false
	}
	count = 0
	for _, v1 := range v.Jobs.Qualification {
		count = 0
		for _, v2 := range jobdetail.Qualification {
			if v1 == v2.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false
	}
	count = 0
	for _, v1 := range v.Jobs.Shift {
		count = 0
		for _, v2 := range jobdetail.Shift {
			if v1 == v2.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false
	}
	count = 0
	for _, v1 := range v.Jobs.Technology {
		count = 0
		for _, v2 := range jobdetail.Technology {
			if v1 == v2.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false
	}
	count = 0
	for _, v1 := range v.Jobs.WorkMode {
		count = 0
		for _, v2 := range jobdetail.WorkMode {
			if v1 == v2.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false
	}
	return true
}
