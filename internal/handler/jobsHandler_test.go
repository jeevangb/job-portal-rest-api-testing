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

// func TestHandler_AddJob(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		// ////////////////////////test cases
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
// 		///////////////////////
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
// 		//////////////////////////////////////////////////
// 		{
// 			name: "failure to  decode",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
// 					"title": "Senior Software Engineer",
// 					"minNoticePeriod": "0",
// 					"maxNoticePeriod": "60",
// 					"budget": "800000",
// 					"locations": [1, 2],
// 					"technologies": [1, 2],
// 					"workmodes": [1, 2],
// 					"description": "Senior Software Engineer position,
// 					"qualifications": [1, 2],
// 					"shifts": [1, 2],
// 					"jobTypes": [1, 2]
// 				  }`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `"Bad Request"{"error":"please provide valid name, location and field"}`,
// 		},
// 		///////////////////////////////////////////////////
// 		{
// 			name: "success to add data database",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
// 					"title": "Senior Software Engineer",
// 					"minNoticePeriod": "0",
// 					"maxNoticePeriod": "60",
// 					"budget": "800000",
// 					"locations": [1, 2],
// 					"technologies": [1, 2],
// 					"workmodes": [1, 2],
// 					"description": "Senior Software Engineer position.",
// 					"qualifications": [1, 2],
// 					"shifts": [1, 2],
// 					"jobTypes": [1, 2]
// 				  }`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().AddJobDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.ResponseJobId{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `"Bad Request"{"ID":0}`,
// 		},
// 		{
// 			name: "failure to add data database",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
// 					"title": "Senior Software Engineer",
// 					"minNoticePeriod": "0",
// 					"maxNoticePeriod": "60",
// 					"budget": "800000",
// 					"locations": [1, 2],
// 					"technologies": [1, 2],
// 					"workmodes": [1, 2],
// 					"description": "Senior Software Engineer position.",
// 					"qualifications": [1, 2],
// 					"shifts": [1, 2],
// 					"jobTypes": [1, 2]
// 				  }`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().AddJobDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.ResponseJobId{}, errors.New("test error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `"Bad Request"{"error":"test error"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.AddJob(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }

// func TestHandler_ProcessJobDetails(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{

// 		// ////////////////////////test cases
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
// 		/////////////////////////////////////////////////////////////////
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
// 		//////////////////////////////////////////////////
// 		{
// 			name: "failure to  decode",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com:8080", strings.NewReader(`[
// 					{
// 						"name": "1Jeevan",
// 						"jid": 1,
// 						"job_appication": {
// 							"cid": 1,
// 							"salary": "12000",
// 							"noticePeriod": "30",
// 							"budget": "800000",
// 							"jobLocations": [
// 								1,
// 								2
// 							],
// 							"technologies": [
// 								1,
// 								2
// 							],
// 							"workmodes": [
// 								1,
// 								2
// 							],
// 							"qualifications": [
// 								1,
// 								24
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobTypes": [
// 								1,
// 								2

// 						}
// 					},
// 					 {
// 						"name": "2afthab",
// 						"jid": 1,
// 						"job_appication": {
// 							"cid": 1,
// 							"salary": "12000",
// 							"noticePeriod": "30",
// 							"budget": "800000",
// 							"jobLocations": [
// 								1,
// 								2
// 							],
// 							"technologies": [
// 								1,
// 								2
// 							],
// 							"workmodes": [
// 								1,
// 								2
// 							],
// 							"qualifications": [
// 								1,
// 								2
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobTypes": [
// 								1,
// 								2
// 							]
// 						}
// 					}
// 				]`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide all fields"}`,
// 		},
// 		{
// 			name: "success to filter records from service layer",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
// 					{
// 						"name": "1Jeevan",
// 						"jid": 1,
// 						"job_appication": {
// 							"cid": 1,
// 							"salary": "12000",
// 							"noticePeriod": "30",
// 							"budget": "800000",
// 							"jobLocations": [
// 								1,
// 								2
// 							],
// 							"technologies": [
// 								1,
// 								2
// 							],
// 							"workmodes": [
// 								1,
// 								2
// 							],
// 							"qualifications": [
// 								1,
// 								24
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobTypes": [
// 								1,
// 								2
// 							]
// 						}
// 					}
// 				]`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().FilterJob(gomock.Any(), gomock.Any()).Return([]models.RespondJobApplicant{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `[]`,
// 		},
// 		{
// 			name: "failure to filter records from service layer",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
// 					{
// 						"name": "1Jeevan",
// 						"jid": 1,
// 						"job_appication": {
// 							"cid": 1,
// 							"salary": "12000",
// 							"noticePeriod": "30",
// 							"budget": "800000",
// 							"jobLocations": [
// 								1,
// 								2
// 							],
// 							"technologies": [
// 								1,
// 								2
// 							],
// 							"workmodes": [
// 								1,
// 								2
// 							],
// 							"qualifications": [
// 								1,
// 								24
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobTypes": [
// 								1,
// 								2
// 							]
// 						}
// 					}
// 				]`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := service.NewMockUserService(mc)

// 				ms.EXPECT().FilterJob(gomock.Any(), gomock.Any()).Return([]models.RespondJobApplicant{}, errors.New("test errror")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"unable to filter records"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := &Handler{
// 				Service: ms,
// 			}

// 			h.ProcessJobDetails(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
