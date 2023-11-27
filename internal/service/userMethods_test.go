package service

import (
	"context"
	"errors"
	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/cache"
	"jeevan/jobportal/internal/models"
	"jeevan/jobportal/internal/repository"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_ResetPassword(t *testing.T) {
	type args struct {
		ctx       context.Context
		resetData models.ResetPassword
	}
	tests := []struct {
		name              string
		args              args
		wantErr           bool
		want              string
		mockCacheResponse func() (string, error)
		mockRepoResponse  func() error
	}{
		{
			name: "password comparision error",
			args: args{
				ctx: context.Background(),
				resetData: models.ResetPassword{
					Email:           "jeevangb@gmail.com",
					NewPassword:     "123",
					ConfirmPassword: "13",
					Otp:             "543215",
				},
			},
			wantErr:           true,
			want:              "password does not match",
			mockCacheResponse: nil,
			mockRepoResponse:  nil,
		},
		{
			name: "password comparision success",
			args: args{
				ctx: context.Background(),
				resetData: models.ResetPassword{
					Email:           "jeevangb@gmail.com",
					NewPassword:     "123",
					ConfirmPassword: "123",
					Otp:             "543215",
				},
			},
			wantErr: true,
			want:    "",
			mockCacheResponse: func() (string, error) {
				return "", errors.New("test error")
			},
			mockRepoResponse: nil,
		},
		{
			name: "otp comparision failure case",
			args: args{
				ctx: context.Background(),
				resetData: models.ResetPassword{
					Email:           "jeevangb@gmail.com",
					NewPassword:     "123",
					ConfirmPassword: "123",
					Otp:             "543215",
				},
			},
			wantErr: true,
			want:    "otp mismatch",
			mockCacheResponse: func() (string, error) {
				return "3424", nil
			},
			mockRepoResponse: nil,
		},
		{
			name: "password hashing failure",
			args: args{
				ctx: context.Background(),
				resetData: models.ResetPassword{
					Email:           "jeevangb@gmail.com",
					NewPassword:     "",
					ConfirmPassword: "",
					Otp:             "5432",
				},
			},
			wantErr: true,
			want:    "",
			mockCacheResponse: func() (string, error) {
				return "5432", nil
			},
			mockRepoResponse: nil,
		},
		{
			name: "password update failure",
			args: args{
				ctx: context.Background(),
				resetData: models.ResetPassword{
					Email:           "jeevangb@gmail.com",
					NewPassword:     "123",
					ConfirmPassword: "123",
					Otp:             "5432",
				},
			},
			wantErr: true,
			want:    "password reset failed",
			mockCacheResponse: func() (string, error) {
				return "5432", nil
			},
			mockRepoResponse: func() error {
				return errors.New("test error")
			},
		},
		{
			name: "password update failure",
			args: args{
				ctx: context.Background(),
				resetData: models.ResetPassword{
					Email:           "jeevangb@gmail.com",
					NewPassword:     "123",
					ConfirmPassword: "123",
					Otp:             "5432",
				},
			},
			wantErr: false,
			want:    "password reset failed",
			mockCacheResponse: func() (string, error) {
				return "5432", nil
			},
			mockRepoResponse: func() error {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCache := cache.NewMockCaching(mc)
			mockrepo := repository.NewMockUserRepo(mc)

			s, _ := NewService(mockrepo, &auth.Auth{}, mockCache)
			if tt.mockCacheResponse != nil {
				mockCache.EXPECT().CheckCacheOtp(gomock.Any(), tt.args.resetData.Email).Return(tt.mockCacheResponse()).AnyTimes()
			}
			if tt.mockRepoResponse != nil {
				mockrepo.EXPECT().UpdatePassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			if err := s.ResetPassword(tt.args.ctx, tt.args.resetData); (err != nil) != tt.wantErr {
				t.Errorf("Service.ResetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_CheckUserDataAndSendOtp(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.ForgotPasswod
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockREpo := repository.NewMockUserRepo(mc)
			mockCache := cache.NewMockCaching(mc)
			s, _ := NewService(mockREpo, &auth.Auth{}, mockCache)
			if err := s.CheckUserDataAndSendOtp(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckUserDataAndSendOtp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
