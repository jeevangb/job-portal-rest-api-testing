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

// func TestService_AddCompanyDetails(t *testing.T) {
// 	type args struct {
// 		ctx         context.Context
// 		companyData models.Company
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         models.Company
// 		wantErr      bool
// 		mockResponse func() (models.Company, error)
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "failure add Company data",
// 			args: args{
// 				ctx:         context.Background(),
// 				companyData: models.Company{},
// 			},
// 			want:    models.Company{},
// 			wantErr: true,
// 			mockResponse: func() (models.Company, error) {
// 				return models.Company{}, errors.New("error test")
// 			},
// 		},
// 		{
// 			name: "success add Company data",
// 			args: args{
// 				ctx: context.Background(),
// 				companyData: models.Company{
// 					Name:     "teksystem",
// 					Location: "banglore",
// 				},
// 			},
// 			want: models.Company{
// 				Name:     "teksystem",
// 				Location: "banglore",
// 			},
// 			wantErr: false,
// 			mockResponse: func() (models.Company, error) {
// 				return models.Company{
// 					Name:     "teksystem",
// 					Location: "banglore",
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockResponse != nil {
// 				mockRepo.EXPECT().CreateCompany(tt.args.ctx, tt.args.companyData).Return(tt.mockResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockRepo, &auth.Auth{})
// 			got, err := s.AddCompanyDetails(tt.args.ctx, tt.args.companyData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.AddCompanyDetails() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.AddCompanyDetails() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewAllCompanies(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         []models.Company
// 		wantErr      bool
// 		mockResponse func() ([]models.Company, error)
// 	}{
// 		{
// 			name: "Failure to fetch all data from db",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			want:    nil,
// 			wantErr: true,
// 			mockResponse: func() ([]models.Company, error) {
// 				return nil, errors.New("error test")
// 			},
// 		},
// 		{
// 			name: "success to fetch all data from db",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			want: []models.Company{
// 				{
// 					Name:     "Tek system",
// 					Location: "Banglore",
// 				},
// 				{
// 					Name:     "Accenture",
// 					Location: "Banglore",
// 				},
// 			},
// 			wantErr: false,
// 			mockResponse: func() ([]models.Company, error) {
// 				return []models.Company{
// 					{
// 						Name:     "Tek system",
// 						Location: "Banglore",
// 					},
// 					{
// 						Name:     "Accenture",
// 						Location: "Banglore",
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
// 				mockRepo.EXPECT().ViewCompanies(tt.args.ctx).Return(tt.mockResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockRepo, &auth.Auth{})
// 			got, err := s.ViewAllCompanies(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewAllCompanies() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewAllCompanies() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestService_ViewCompanyById(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		cid uint64
// 	}
// 	tests := []struct {
// 		name         string
// 		s            *Service
// 		args         args
// 		want         models.Company
// 		wantErr      bool
// 		mockResponse func() (models.Company, error)
// 	}{
// 		{
// 			name: "failure case",
// 			args: args{
// 				ctx: context.Background(),
// 				cid: 1,
// 			},
// 			want:    models.Company{},
// 			wantErr: true,
// 			mockResponse: func() (models.Company, error) {
// 				return models.Company{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			name: "success case",
// 			args: args{
// 				ctx: context.Background(),
// 				cid: 1,
// 			},
// 			want: models.Company{
// 				Name:     "Tek system",
// 				Location: "Banglore",
// 			},
// 			wantErr: false,
// 			mockResponse: func() (models.Company, error) {
// 				return models.Company{
// 					Name:     "Tek system",
// 					Location: "Banglore",
// 				}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := repository.NewMockUserRepo(mc)
// 			if tt.mockResponse != nil {
// 				mockRepo.EXPECT().ViewCompanyById(tt.args.ctx, tt.args.cid).Return(tt.mockResponse()).AnyTimes()
// 			}
// 			s, _ := NewService(mockRepo, &auth.Auth{})
// 			got, err := s.ViewCompanyById(tt.args.ctx, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ViewCompanyById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ViewCompanyById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
