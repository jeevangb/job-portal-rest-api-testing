package service

import (
	"context"
	"encoding/json"
	"fmt"
	"jeevan/jobportal/internal/models"
	"log"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"

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
	/////////////////////////////
	for _, v := range jobData.JobLocation {
		jobLoc := models.JobLocation{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.JobLocation = append(createjobDetails.JobLocation, jobLoc)
	}
	////////////////////////////
	for _, v := range jobData.Technology {
		jobTech := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.Technology = append(createjobDetails.Technology, jobTech)
	}
	////////////////////////////
	for _, v := range jobData.WorkMode {
		jobWorkMode := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.WorkMode = append(createjobDetails.WorkMode, jobWorkMode)
	}

	////////////////////////////
	for _, v := range jobData.Qualification {
		jobWorkQualification := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.Qualification = append(createjobDetails.Qualification, jobWorkQualification)
	}

	///////////////////////////////////////////////
	for _, v := range jobData.Shift {
		jobShift := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		createjobDetails.Shift = append(createjobDetails.Shift, jobShift)
	}
	///////////////////////////////////////////////
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
	fmt.Println(jobApplications)

	ch := make(chan models.RespondJobApplicant)
	wg := new(sync.WaitGroup)

	for _, jobApplication := range jobApplications {
		wg.Add(1)
		go func(jobApplication models.RespondJobApplicant) {
			defer wg.Done()

			// Create a new Redis client for each goroutine
			redisClient := redis.NewClient(&redis.Options{
				Addr:     "localhost:6379", // Redis server address
				Password: "",               // No password by default
				DB:       0,                // Default DB
			})

			// Check if the Redis client is connected successfully
			_, err := redisClient.Ping(ctx).Result()
			if err != nil {
				log.Fatalf("Error connecting to Redis: %v", err)
				return
			}

			// Use Redis to retrieve job details
			jobDetailKey := fmt.Sprintf("job:%d", jobApplication.Jid)

			jobDetailStr, err := redisClient.Get(ctx, jobDetailKey).Result()

			if err != nil {

				// If the key is not in Redis, fetch it from the database
				val, err := s.UserRepo.ViewJobDetailsBy(ctx, uint64(jobApplication.Jid))
				fmt.Println("//////////////////////////////////////////////////", val)
				if err != nil {
					return
				}
				// Serialize the job details to JSON before storing in Redis
				valJSON, err := json.Marshal(val)
				if err != nil {
					return
				}

				// Store the job details in Redis
				err = redisClient.Set(ctx, jobDetailKey, valJSON, 0).Err()
				fmt.Println(err)
				if err != nil {
					return
				}
				jobDetailStr = string(valJSON)

			}

			// Deserialize job details
			var jobDetail models.Jobs
			err = json.Unmarshal([]byte(jobDetailStr), &jobDetail)
			if err != nil {
				return
			}

			b := checkApplicantsCriteria(jobApplication, jobDetail)
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
	////////////////////////////////filter budget
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
	/////////////////////////////////////////////filter notice period
	// hrMinNoticePeriod, err := strconv.Atoi(jobdetail.MinNoticePeriod)
	// fmt.Println(hrMinNoticePeriod)
	// if err != nil {

	// 	panic("error while conversion min notice  period data from hr posting")
	// }
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

	////////////////////////////////////////////////////////////////////////////
	for _, v1 := range v.Jobs.JobLocation {
		count = 0
		for _, v2 := range jobdetail.JobLocation {
			if v1 == v2.ID {
				count++

			}
		}
	}
	if count == 0 {
		return false
	}

	//////////////////////////////////////////////////////////////////////
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

	//////////////////////////////////////////////////////////////////////////
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

	////////////////////////////////////////////////////////////////////////
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

	////////////////////////////////////////////////////////////////////////
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
	////////////////////////////////////////////////////////////////////////
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
