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

// func TestHandler_ViewCompany(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{

// 		/////////////////////////////////////////////////////////////////////////////////
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
// 		//////////////////////////////////////////////////////////////////////////
// 		{
// 			name: "missing jwt claims",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusUnauthorized,
// 			expectedResponse:   `{"error":"Unauthorized"}`,
// 		},
// 		{
// 			name: "failure while fetching jobs from service",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().ViewCompanyById(c.Request.Context(), gomock.Any()).Return(models.Company{}, errors.New("test error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `"Bad Request"{"error":"test error"}`,
// 		},
// 		///////////////////////////////////////////////////////////////////////////
// 		{
// 			name: "success while fetching jobs from service",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().ViewCompanyById(c.Request.Context(), gomock.Any()).Return(models.Company{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","location":""}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.ViewCompany(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }

// func TestHandler_ViewAllCompanies(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		////////////////////////////////////////////////////////////////////////////////////////////////////
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
// 		// ////////////////////////////////////////////////////////////////
// 		{
// 			name: "missing jwt claims",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusUnauthorized,
// 			expectedResponse:   `{"error":"Unauthorized"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().ViewAllCompanies(gomock.Any()).Return([]models.Company{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `[]`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.ViewAllCompanies(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }

// func TestHandler_AddCompany(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		// TODO: Add test cases.
// 		///////////////////////////////////////////////////////////////////////////////////////////
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
// 			name: "missing jwt claims",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusUnauthorized,
// 			expectedResponse:   `{"error":"Unauthorized"}`,
// 		},

// 		{
// 			name: "Success while creating a company",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"name":"Tek system","location":"Bnglr"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)
// 				ms.EXPECT().AddCompanyDetails(gomock.Any(), gomock.Any()).Return(models.Company{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","location":""}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.AddCompany(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
