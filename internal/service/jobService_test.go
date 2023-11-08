package service

// import (
// 	"context"
// 	"errors"
// 	"jeevan/jobportal/internal/auth"
// 	"jeevan/jobportal/internal/models"
// 	"jeevan/jobportal/internal/repository"
// 	"reflect"
// 	"testing"

// 	"go.uber.org/mock/gomock"
// )

// func TestService_ViewJobById(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		jid uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         models.Jobs
// 		wantErr      bool
// 		mockResponse func() (models.Jobs, error)
// 	}{
// 		{
// 			name: "error from db",
// 			args: args{
// 				ctx: context.Background(),
// 				jid: 1,
// 			},
// 			want:    models.Jobs{},
// 			wantErr: true,
// 			mockResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			name: "success from db",
// 			args: args{
// 				ctx: context.Background(),
// 				jid: 11,
// 			},
// 			want: models.Jobs{
// 				Company: models.Company{
// 					Name: "Tek System",
// 				}, Cid: 1,
// 				Title: "Software",
// 			},
// 			wantErr: false,
// 			mockResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Company: models.Company{
// 						Name: "Tek System",
// 					}, Cid: 1, Title: "Software",
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockResponse != nil {
// 				mockRepo.EXPECT().ViewJobDetailsBy(tt.args.ctx, tt.args.jid).Return(tt.mockResponse()).AnyTimes()
// 			}

// 			s, _ := NewService(mockRepo, &auth.Auth{})
// 			got, err := s.ViewJobById(tt.args.ctx, tt.args.jid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewJobById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewAllJobs(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         []models.Jobs
// 		wantErr      bool
// 		mockResponse func() ([]models.Jobs, error)
// 	}{
// 		{
// 			name: "error case to get all data ",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			want:    nil,
// 			wantErr: true,
// 			mockResponse: func() ([]models.Jobs, error) {
// 				return nil, errors.New("test error")

// 			},
// 		},
// 		{
// 			name: "success case to get all data ",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			want: []models.Jobs{
// 				{
// 					Cid:    1,
// 					Title:  "Tek system",
// 					Salary: "500000",
// 				},
// 				{
// 					Cid:    2,
// 					Title:  "Accenture",
// 					Salary: "6000",
// 				},
// 			},
// 			wantErr: false,
// 			mockResponse: func() ([]models.Jobs, error) {
// 				return []models.Jobs{
// 					{
// 						Cid:    1,
// 						Title:  "Tek system",
// 						Salary: "500000",
// 					},
// 					{
// 						Cid:    2,
// 						Title:  "Accenture",
// 						Salary: "6000",
// 					},
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockResponse != nil {
// 				mockRepo.EXPECT().FindAllJobs(tt.args.ctx).Return(tt.mockResponse()).AnyTimes()
// 			}

// 			s, _ := NewService(mockRepo, &auth.Auth{})
// 			got, err := s.ViewAllJobs(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewAllJobs() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewAllJobs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_AddJobDetails(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		jobData models.Jobs
// 		cid     uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         models.Jobs
// 		wantErr      bool
// 		mockResponse func() (models.Jobs, error)
// 	}{
// 		{
// 			name: "error case",
// 			args: args{
// 				ctx:     context.Background(),
// 				jobData: models.Jobs{},
// 			},
// 			want:    models.Jobs{},
// 			wantErr: true,
// 			mockResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			name: "success case",
// 			args: args{
// 				ctx: context.Background(),
// 				jobData: models.Jobs{
// 					Title:  "Developer",
// 					Salary: "45000",
// 					Cid:    1,
// 				},
// 				cid: 1,
// 			},
// 			want: models.Jobs{
// 				Title:  "Developer",
// 				Salary: "45000",
// 				Cid:    1,
// 			},

// 			wantErr: false,

// 			mockResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Title:  "Developer",
// 					Salary: "45000",
// 					Cid:    1,
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockResponse != nil {
// 				mockRepo.EXPECT().CreateJob(tt.args.ctx, tt.args.jobData).Return(tt.mockResponse()).AnyTimes()
// 			}

// 			s, _ := NewService(mockRepo, &auth.Auth{})
// 			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewJobByCid(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		cid uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         []models.Jobs
// 		wantErr      bool
// 		mockResponse func() ([]models.Jobs, error)
// 	}{
// 		{
// 			name: "failure to get details",
// 			args: args{
// 				ctx: context.Background(),
// 				cid: uint64(1),
// 			},
// 			want:    nil,
// 			wantErr: true,
// 			mockResponse: func() ([]models.Jobs, error) {
// 				return nil, errors.New("test error")
// 			},
// 		},
// 		{
// 			name: "success case to get details",
// 			args: args{
// 				ctx: context.Background(),
// 				cid: uint64(1),
// 			},
// 			want: []models.Jobs{
// 				{
// 					Cid:    1,
// 					Title:  "Developer",
// 					Salary: "500000",
// 				},
// 				{
// 					Cid:    2,
// 					Title:  "Tester",
// 					Salary: "600000",
// 				},
// 			},
// 			wantErr: false,
// 			mockResponse: func() ([]models.Jobs, error) {
// 				return []models.Jobs{
// 					{
// 						Cid:    1,
// 						Title:  "Developer",
// 						Salary: "500000",
// 					},
// 					{
// 						Cid:    2,
// 						Title:  "Tester",
// 						Salary: "600000",
// 					},
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockResponse != nil {
// 				mockRepo.EXPECT().FindJob(gomock.Any(), tt.args.cid).Return(tt.mockResponse()).AnyTimes()
// 			}

// 			s, _ := NewService(mockRepo, &auth.Auth{})

// 			got, err := s.ViewJobByCid(tt.args.ctx, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewJobByCid() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewJobByCid() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
