package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"jeevan/jobportal/internal/models"

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

func (s *Service) FilterJob(ctx context.Context, jobApplication []models.RespondJobApplicant) ([]models.RespondJobApplicant, error) {

	var FilterJobData []models.RespondJobApplicant
	jobdetail, err := s.UserRepo.ViewJobDetailsBy(ctx, uint64(4))

	if err != nil {
		return nil, errors.New("not able to get  jobs from database")
	}

	ch := make(chan models.RespondJobApplicant)
	wg := new(sync.WaitGroup)

	for _, v := range jobApplication {
		wg.Add(1)
		go func(v models.RespondJobApplicant) {
			defer wg.Done()
			bool, singleApplication := checkApplicantsCriteria(v, jobdetail)
			if bool {
				// FilterJobData = append(FilterJobData, singleApplication)
				ch <- singleApplication
			}

		}(v)
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

func checkApplicantsCriteria(v models.RespondJobApplicant, jobdetail models.Jobs) (bool, models.RespondJobApplicant) {
	////////////////////////////////filter budget
	applicantBudget, err := strconv.Atoi(v.Jobs.Budget)
	if err != nil {
		panic("error while conversion budget data from applicants")
	}
	hrBudget, err := strconv.Atoi(jobdetail.Budget)
	if err != nil {
		panic("error while conversion budget data from hr posting")
	}
	if applicantBudget > hrBudget {
		return false, models.RespondJobApplicant{}
	}
	/////////////////////////////////////////////filter notice period
	hrMinNoticePeriod, err := strconv.Atoi(jobdetail.MinNoticePeriod)
	fmt.Println(hrMinNoticePeriod)
	if err != nil {
		panic("error while conversion min notice  period data from hr posting")
	}
	hrMaxNoticePeriod, err := strconv.Atoi(jobdetail.MaxNoticePeriod)
	fmt.Println(hrMaxNoticePeriod)
	if err != nil {
		panic("error while conversion max notice period data from hr posting")
	}
	applicantNoticePeriod, err := strconv.Atoi(v.Jobs.Np)
	if err != nil {
		panic("error while conversion notice period from applicant")
	}

	if (applicantNoticePeriod < hrMinNoticePeriod) || (applicantNoticePeriod > hrMaxNoticePeriod) {
		return false, models.RespondJobApplicant{}
	}
	if v.Jobs.JobDescription != jobdetail.Jobdescription {
		return false, models.RespondJobApplicant{}
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
		return false, models.RespondJobApplicant{}
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
		return false, models.RespondJobApplicant{}
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
		return false, models.RespondJobApplicant{}
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
		return false, models.RespondJobApplicant{}
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
		return false, models.RespondJobApplicant{}
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
		return false, models.RespondJobApplicant{}
	}
	return true, v
}
