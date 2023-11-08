package handler

// import (
// 	"context"
// 	"errors"
// 	"jeevan/jobportal/internal/auth"
// 	"jeevan/jobportal/internal/middleware"
// 	"jeevan/jobportal/internal/models"
// 	"jeevan/jobportal/internal/service"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// 	"github.com/golang-jwt/jwt/v5"
// 	"go.uber.org/mock/gomock"
// )

// func TestHandler_Login(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "Success while logginmg user",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`"email":"jee@gmail.com"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)
// 				ms.EXPECT().UserLogin(gomock.Any(), gomock.Any()).Return("", errors.New("test error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide valid email and password"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.Login(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }

// func TestHandler_SignUp(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "Success while creating user",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"username":"jeevan","email":"jee@gmail.com","password":"1234"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)
// 				ms.EXPECT().UserSignup(gomock.Any(), gomock.Any()).Return(models.User{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"username":"","email":""}`,
// 		},
// 		{
// 			name: "failure while creating user",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"username":"jeevan","password":"1234"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)
// 				ms.EXPECT().UserSignup(gomock.Any(), gomock.Any()).Return(models.User{}, errors.New("test error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide valid username, email and password"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.SignUp(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
