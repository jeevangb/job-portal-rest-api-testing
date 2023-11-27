package service

// import (
// 	"context"
// 	"errors"
// 	"jeevan/jobportal/internal/auth"
// 	"jeevan/jobportal/internal/cache"
// 	"jeevan/jobportal/internal/models"
// 	"jeevan/jobportal/internal/repository"
// 	"reflect"
// 	"testing"

// 	"github.com/go-redis/redis/v8"
// 	"go.uber.org/mock/gomock"
// )

// func TestService_FilterJob(t *testing.T) {
// 	type args struct {
// 		ctx             context.Context
// 		jobApplications []models.RespondJobApplicant
// 	}
// 	tests := []struct {
// 		name                 string
// 		args                 args
// 		want                 []models.RespondJobApplicant
// 		wantErr              bool
// 		mockCacheResponse    func() (string, error)
// 		mockAddCacheResponse func() error
// 		mockRepoResponse     func() (models.Jobs, error)
// 	}{
// 		{
// 			// TODO: Add test cases.
// 			name: "fail to marshal the data from the redis",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "20",
// 							Budget:        "10000",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `Model: gorm.Model{
// 					ID: 1,
// 				},
// 				Company: models.Company{
// 					Model: gorm.Model{
// 						ID: uint(1),
// 					},
// 					Name: "jeevan",
// 				},
// 				Cid:   1,
// 				Title: "go",
// 				MaxNoticePeriod: "20",
// 				Budget:          "1000",
// 				JobLocation: []models.JobLocation{
// 					{
// 						Model: gorm.Model{ID: uint(1)},
// 					},
// 					{
// 						Model: gorm.Model{ID: uint(2)},
// 					},
// 				},
// 				Technology: []models.Technology{
// 					{
// 						Model: gorm.Model{ID: uint(1)},
// 					},
// 					{
// 						Model: gorm.Model{ID: uint(2)},
// 					},
// 				},
// 				WorkMode: []models.WorkMode{
// 					{Model: gorm.Model{ID: uint(1)}},
// 				},
// 				Shift: []models.Shift{
// 					{Model: gorm.Model{ID: uint(1)}},

// 				},
// 				Qualification: []models.Qualification{
// 					{Model: gorm.Model{ID: uint(1)}},
// 				},
// 				JobType: []models.JobType{
// 					{Model: gorm.Model{ID: uint(1)}},
// 				},
// 				MinNoticePeriod: "20",
// 				Jobdescription:  "go",`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while converting the budget",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "20",
// 							Budget:        "jeevan",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while converting the notice period",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "jeevan",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare notice period",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2000",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare job location",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{100},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare job type",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{100},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare job qualification",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{100},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare job shift",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{100},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare job shift",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{100},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "error while compare work mode",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "2",
// 							Budget:        "100",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{100},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "success case from redis",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "20",
// 							Budget:        "10000",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want: []models.RespondJobApplicant{
// 				{
// 					Name: "jeevan",
// 					Jid:  1,
// 					Jobs: models.JobApplicant{
// 						Jid:           1,
// 						Title:         "go",
// 						Np:            "20",
// 						Budget:        "10000",
// 						Salary:        "10000",
// 						JobLocation:   []uint{1},
// 						Technology:    []uint{1},
// 						WorkMode:      []uint{1},
// 						Qualification: []uint{1},
// 						Shift:         []uint{1},
// 						JobType:       []uint{1},
// 					},
// 				},
// 			},
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":6,"CreatedAt":"2023-11-08T13:02:36.768974+05:30","UpdatedAt":"2023-11-08T13:02:36.768974+05:30","DeletedAt":null,"cid":1,"title":"Java Developer","minnp":"0","maxnp":"60","budget":"7000000","JobLocation":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Banglore"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Hyderbad"}],"Technology":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"java"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"technologyName":"Python"}],"WorkMode":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Online"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"workMode":"Offline"}],"job_description":"backend developer","Qualification":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BE"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification":"BCA"}],"Shift":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Night"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift":"Day"}],"JobType":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"FullTime"},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"jobType":"PartTime"}]}`, nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "failure case from redis",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "20",
// 							Budget:        "10000",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			// TODO: Add test cases.
// 			name: "failure case to data to database",
// 			args: args{
// 				ctx: context.Background(),
// 				jobApplications: []models.RespondJobApplicant{
// 					{
// 						Name: "jeevan",
// 						Jid:  1,
// 						Jobs: models.JobApplicant{
// 							Jid:           1,
// 							Title:         "go",
// 							Np:            "20",
// 							Budget:        "10000",
// 							Salary:        "10000",
// 							JobLocation:   []uint{1},
// 							Technology:    []uint{1},
// 							WorkMode:      []uint{1},
// 							Qualification: []uint{1},
// 							Shift:         []uint{1},
// 							JobType:       []uint{1},
// 						},
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return errors.New("test error")
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCache := cache.NewMockCaching(mc)
// 			mockCache.EXPECT().GetCahceData(gomock.Any(), gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
// 			mockCache.EXPECT().AddToCache(tt.args.ctx, gomock.Any(), gomock.Any()).Return(tt.mockAddCacheResponse()).AnyTimes()

// 			mockRepo := repository.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().ViewJobDetailsBy(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

// 			s, _ := NewService(mockRepo, &auth.Auth{}, &cache.Rdb{})
// 			got, err := s.FilterJob(tt.args.ctx, tt.args.jobApplications)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.FilterJob() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.FilterJob() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
