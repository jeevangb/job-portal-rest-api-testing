package service

import (
	"context"
	"errors"
	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_UserLogin(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name         string
		s            *Service
		args         args
		want         string
		wantErr      bool
		mockResponse func() (models.User, error)
	}{
		{
			name: "error case at validation email to login",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "Jeevan",
					Email:    "",
					Password: "bhguk23",
				},
			},
			want:    "",
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{}, errors.New("error test")
			},
		},
		{
			name: "Sucess in hashing password",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "Jeevan",
					Email:    "jee@gmail.com",
					Password: "bhguk23",
				},
			},
			want:    "",
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{
					Username:     "Jeevan",
					Email:        "jee@gmail.com",
					PasswordHash: "passhash",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockResponse != nil {
				mockRepo.EXPECT().CheckEmail(tt.args.ctx, tt.args.userData.Email).Return(tt.mockResponse()).AnyTimes()
			}

			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.UserLogin(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.UserLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UserSignup(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name         string
		s            *Service
		args         args
		want         models.User
		wantErr      bool
		mockResponse func() (models.User, error)
	}{
		{
			name: "failure to create table",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "Jeevan",
					Password: "1234",
				},
			},
			want:    models.User{},
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{}, errors.New("error test")
			},
		},
		{
			name: "Success to create table",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "Jeevan",
					Email:    "jeevan@gmail.com",
					Password: "1234",
				},
			},
			want: models.User{
				Username:     "Jeevan",
				Email:        "jeevan@gmail",
				PasswordHash: "gukgjtdyfd",
			},
			wantErr: false,
			mockResponse: func() (models.User, error) {
				return models.User{
					Username:     "Jeevan",
					Email:        "jeevan@gmail",
					PasswordHash: "gukgjtdyfd",
				}, nil
			},
		},
		{
			name: "failure to hashPassword",
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Password: "1234",
				},
			},
			want:    models.User{},
			wantErr: true,
			mockResponse: func() (models.User, error) {
				return models.User{}, errors.New("error test")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockResponse != nil {
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			}

			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.UserSignup(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}
