package service

import (
	"context"
	"errors"
	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.Hr
		cid     uint64
	}
	tests := []struct {
		name         string
		args         args
		want         models.ResponseJobId
		wantErr      bool
		mockResponse func() (models.ResponseJobId, error)
	}{
		// TODO: Add test cases.
		{
			name: "error case to get all data ",
			args: args{
				ctx: context.Background(),
			},
			want:    models.ResponseJobId{},
			wantErr: true,
			mockResponse: func() (models.ResponseJobId, error) {
				return models.ResponseJobId{}, errors.New("test error")

			},
		},
		{
			name: "success case",
			args: args{
				ctx: context.Background(),
				jobData: models.Hr{
					Title:          "merndeceloper",
					Minnp:          "0",
					Maxnp:          "9",
					Budget:         "1000",
					JobLocation:    []uint{1, 2},
					Technology:     []uint{1, 2},
					WorkMode:       []uint{1, 2},
					JobDescription: "go",
					Qualification:  []uint{1, 2},
					Shift:          []uint{1, 2},
					JobType:        []uint{1, 2},
				},
				cid: 1,
			},

			want: models.ResponseJobId{ID: 1},

			wantErr: false,

			mockResponse: func() (models.ResponseJobId, error) {
				return models.ResponseJobId{ID: 1}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockResponse != nil {
				mockRepo.EXPECT().CreateJob(tt.args.ctx, gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			}

			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FilterJob(t *testing.T) {
	type args struct {
		ctx             context.Context
		jobApplications []models.RespondJobApplicant
	}
	tests := []struct {
		name    string
		args    args
		want    []models.RespondJobApplicant
		wantErr bool
		setup   func(mockRepo repository.MockUserRepo)
	}{
		/////////////////////////////////////////////////////////////////////////////////////////////////////////
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				jobApplications: []models.RespondJobApplicant{
					{
						Name: "jeevan",
						Jid:  uint(1),
						Jobs: models.JobApplicant{
							Jid:            uint(1),
							Title:          "developer",
							Salary:         "100",
							Np:             "7",
							Budget:         "1000",
							JobLocation:    []uint{1, 2},
							Technology:     []uint{1},
							WorkMode:       []uint{1},
							JobDescription: "gooo",
							Qualification:  []uint{1, 2},
							Shift:          []uint{1, 2},
							JobType:        []uint{1, 2},
						},
					},
					// {
					// 	Name: "afthab",
					// 	Jid:  uint(2),
					// 	Jobs: models.JobApplicant{
					// 		Jid:            uint(2),
					// 		Title:          "developer",
					// 		Salary:         "100",
					// 		Np:             "7",
					// 		Budget:         "1000",
					// 		JobLocation:    []uint{1, 2},
					// 		Technology:     []uint{1},
					// 		WorkMode:       []uint{1},
					// 		JobDescription: "gooo",
					// 		Qualification:  []uint{1, 2},
					// 		Shift:          []uint{1, 2},
					// 		JobType:        []uint{1, 2},
					// 	},
					// },
					// {
					// 	Name: "ravan",
					// 	Jid:  uint(1),
					// 	Jobs: models.JobApplicant{
					// 		Jid:            uint(1),
					// 		Title:          "developer",
					// 		Salary:         "100",
					// 		Np:             "7",
					// 		Budget:         "1000",
					// 		JobLocation:    []uint{1, 2},
					// 		Technology:     []uint{1},
					// 		WorkMode:       []uint{1},
					// 		JobDescription: "gooo",
					// 		Qualification:  []uint{1, 2},
					// 		Shift:          []uint{1, 2},
					// 		JobType:        []uint{1, 2},
					// 	},
					// },
				},
			},

			want: []models.RespondJobApplicant{
				// {
				// 	Name: "afthab",
				// 	Jid:  uint(1),
				// 	Jobs: models.JobApplicant{
				// 		Jid:            uint(1),
				// 		Title:          "developer",
				// 		Salary:         "100",
				// 		Np:             "7",
				// 		Budget:         "1000",
				// 		JobLocation:    []uint{1, 2},
				// 		Technology:     []uint{1},
				// 		WorkMode:       []uint{1},
				// 		JobDescription: "gooo",
				// 		Qualification:  []uint{1, 2},
				// 		Shift:          []uint{1, 2},
				// 		JobType:        []uint{1, 2},
				// 	},
				// },
				{
					Name: "jeevan",
					Jid:  uint(1),
					Jobs: models.JobApplicant{
						Jid:            uint(1),
						Title:          "developer",
						Salary:         "100",
						Np:             "7",
						Budget:         "1000",
						JobLocation:    []uint{1, 2},
						Technology:     []uint{1},
						WorkMode:       []uint{1},
						JobDescription: "gooo",
						Qualification:  []uint{1, 2},
						Shift:          []uint{1, 2},
						JobType:        []uint{1, 2},
					},
				},
				// {
				// 	Name: "ravan",
				// 	Jid:  uint(1),
				// 	Jobs: models.JobApplicant{
				// 		Jid:            uint(0),
				// 		Title:          "developer",
				// 		Salary:         "100",
				// 		Np:             "7",
				// 		Budget:         "1000",
				// 		JobLocation:    []uint{1, 2},
				// 		Technology:     []uint{1},
				// 		WorkMode:       []uint{1},
				// 		JobDescription: "gooo",
				// 		Qualification:  []uint{1, 2},
				// 		Shift:          []uint{1, 2},
				// 		JobType:        []uint{1, 2},
				// 	},
				// },
			},
			wantErr: false,
			setup: func(mockRepo repository.MockUserRepo) {
				mockRepo.EXPECT().ViewJobDetailsBy(gomock.Any(), uint64(1)).Return(models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:   1,
					Title: "developer",
					// MinNoticePeriod: "0",
					MaxNoticePeriod: "40",
					Budget:          "2000",
					JobLocation: []models.JobLocation{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Technology: []models.Technology{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					WorkMode: []models.WorkMode{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					Qualification: []models.Qualification{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Shift: []models.Shift{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
					JobType: []models.JobType{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
				}, nil).Times(1)

				// mockRepo.EXPECT().ViewJobDetailsBy(gomock.Any(), uint64(2)).Return(models.Jobs{
				// 	Model: gorm.Model{
				// 		ID: 2,
				// 	},
				// 	Company: models.Company{
				// 		Model: gorm.Model{
				// 			ID: 1,
				// 		},
				// 	},
				// 	Cid:   1,
				// 	Title: "developer",
				// 	// MinNoticePeriod: "0",
				// 	MaxNoticePeriod: "40",
				// 	Budget:          "2000",
				// 	JobLocation: []models.JobLocation{
				// 		{
				// 			Model: gorm.Model{
				// 				ID: uint(1),
				// 			},
				// 		},
				// 		{
				// 			Model: gorm.Model{
				// 				ID: uint(2),
				// 			},
				// 		},
				// 	},
				// 	Technology: []models.Technology{
				// 		{
				// 			Model: gorm.Model{ID: uint(1)},
				// 		},
				// 		{
				// 			Model: gorm.Model{ID: uint(2)},
				// 		},
				// 	},
				// 	WorkMode: []models.WorkMode{
				// 		{
				// 			Model: gorm.Model{ID: uint(1)},
				// 		},
				// 		{
				// 			Model: gorm.Model{ID: uint(2)},
				// 		},
				// 	},
				// 	Qualification: []models.Qualification{
				// 		{
				// 			Model: gorm.Model{
				// 				ID: uint(1),
				// 			},
				// 		},
				// 		{
				// 			Model: gorm.Model{
				// 				ID: uint(2),
				// 			},
				// 		},
				// 	},
				// 	Shift: []models.Shift{
				// 		{Model: gorm.Model{ID: uint(1)}},
				// 		{Model: gorm.Model{ID: uint(2)}},
				// 	},
				// 	JobType: []models.JobType{
				// 		{Model: gorm.Model{ID: uint(1)}},
				// 		{Model: gorm.Model{ID: uint(2)}},
				// 	},
				// }, nil).Times(1)
			},
		},
		{
			name: "failure case budget conversion",
			args: args{
				ctx: context.Background(),
				jobApplications: []models.RespondJobApplicant{
					{
						Name: "jeevan",
						Jid:  uint(1),
						Jobs: models.JobApplicant{
							Jid:            uint(1),
							Title:          "developer",
							Salary:         "100",
							Np:             "7",
							Budget:         "Ten thousand",
							JobLocation:    []uint{1, 2},
							Technology:     []uint{1},
							WorkMode:       []uint{1},
							JobDescription: "gooo",
							Qualification:  []uint{1, 2},
							Shift:          []uint{1, 2},
							JobType:        []uint{1, 2},
						},
					},
				},
			},

			want:    nil,
			wantErr: false,
			setup: func(mockRepo repository.MockUserRepo) {
				mockRepo.EXPECT().ViewJobDetailsBy(gomock.Any(), uint64(1)).Return(models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:   1,
					Title: "developer",
					// MinNoticePeriod: "0",
					MaxNoticePeriod: "40",
					Budget:          "2000",
					JobLocation: []models.JobLocation{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Technology: []models.Technology{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					WorkMode: []models.WorkMode{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					Qualification: []models.Qualification{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Shift: []models.Shift{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
					JobType: []models.JobType{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
				}, nil).Times(0)

			},
		},
		{
			name: "failure budget comparison",
			args: args{
				ctx: context.Background(),
				jobApplications: []models.RespondJobApplicant{
					{
						Name: "jeevan",
						Jid:  uint(1),
						Jobs: models.JobApplicant{
							Jid:            uint(1),
							Title:          "developer",
							Salary:         "100",
							Np:             "7",
							Budget:         "10000",
							JobLocation:    []uint{1, 2},
							Technology:     []uint{1},
							WorkMode:       []uint{1},
							JobDescription: "gooo",
							Qualification:  []uint{1, 2},
							Shift:          []uint{1, 2},
							JobType:        []uint{1, 2},
						},
					},
				},
			},

			want:    nil,
			wantErr: false,
			setup: func(mockRepo repository.MockUserRepo) {
				mockRepo.EXPECT().ViewJobDetailsBy(gomock.Any(), uint64(1)).Return(models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:   1,
					Title: "developer",
					// MinNoticePeriod: "0",
					MaxNoticePeriod: "40",
					Budget:          "2000",
					JobLocation: []models.JobLocation{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Technology: []models.Technology{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					WorkMode: []models.WorkMode{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					Qualification: []models.Qualification{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Shift: []models.Shift{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
					JobType: []models.JobType{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
				}, nil).Times(0)

			},
		},
		{
			name: "failure notice period comparison",
			args: args{
				ctx: context.Background(),
				jobApplications: []models.RespondJobApplicant{
					{
						Name: "jeevan",
						Jid:  uint(1),
						Jobs: models.JobApplicant{
							Jid:            uint(1),
							Title:          "developer",
							Salary:         "100",
							Np:             "20",
							Budget:         "10000",
							JobLocation:    []uint{1, 2},
							Technology:     []uint{1},
							WorkMode:       []uint{1},
							JobDescription: "gooo",
							Qualification:  []uint{1, 2},
							Shift:          []uint{1, 2},
							JobType:        []uint{1, 2},
						},
					},
				},
			},

			want:    nil,
			wantErr: false,
			setup: func(mockRepo repository.MockUserRepo) {
				mockRepo.EXPECT().ViewJobDetailsBy(gomock.Any(), uint64(1)).Return(models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:   1,
					Title: "developer",
					// MinNoticePeriod: "0",
					MaxNoticePeriod: "40",
					Budget:          "two thousand",
					JobLocation: []models.JobLocation{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Technology: []models.Technology{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					WorkMode: []models.WorkMode{
						{
							Model: gorm.Model{ID: uint(1)},
						},
						{
							Model: gorm.Model{ID: uint(2)},
						},
					},
					Qualification: []models.Qualification{
						{
							Model: gorm.Model{
								ID: uint(1),
							},
						},
						{
							Model: gorm.Model{
								ID: uint(2),
							},
						},
					},
					Shift: []models.Shift{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
					JobType: []models.JobType{
						{Model: gorm.Model{ID: uint(1)}},
						{Model: gorm.Model{ID: uint(2)}},
					},
				}, nil).Times(0)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			tt.setup(*mockRepo)

			s := &Service{
				UserRepo: mockRepo,
			}
			// if tt.mockResponse != nil {
			// 	mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, uint(1)).Return(tt.mockResponse()).AnyTimes()
			// }

			// s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.FilterJob(tt.args.ctx, tt.args.jobApplications)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
