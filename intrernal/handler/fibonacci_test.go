package handler

import (
	"errors"
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/intrernal/service"
	mock_service "github.com/alex-dev-master/fibonacci.git/intrernal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_getFibonacci(t *testing.T) {
	type mockBehavior func(s *mock_service.MockFibonacci, input model.Fibonacci)

	testTable := []struct {
		name                 string
		inputBody            string
		paramsForRequest     string
		inputUser            model.Fibonacci
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:             "OK",
			paramsForRequest: "x=0&y=2",
			inputUser: model.Fibonacci{
				X: 0,
				Y: 2,
			},
			mockBehavior: func(s *mock_service.MockFibonacci, input model.Fibonacci) {
				s.EXPECT().GetSlice(input).Return([]uint64{0, 1, 1}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[0,1,1]`,
		},
		{
			name:             "Empty fields",
			paramsForRequest: "",
			inputUser: model.Fibonacci{

			},
			mockBehavior:         func(s *mock_service.MockFibonacci, input model.Fibonacci) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:             "Fields values must be less than 92",
			paramsForRequest: "x=0&y=93",
			inputUser: model.Fibonacci{
				X: 0,
				Y: 93,
			},
			mockBehavior:         func(s *mock_service.MockFibonacci, input model.Fibonacci) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Y and X must be less than 92"}`,
		},
		{
			name:             "Y should have more than X",
			paramsForRequest: "x=5&y=3",
			inputUser: model.Fibonacci{
				X: 5,
				Y: 3,
			},
			mockBehavior:         func(s *mock_service.MockFibonacci, input model.Fibonacci) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Y should have more than X"}`,
		},
		{
			name:             "Service Failure",
			paramsForRequest: "x=0&y=2",
			inputUser: model.Fibonacci{
				X: 0,
				Y: 2,
			},
			mockBehavior: func(s *mock_service.MockFibonacci, input model.Fibonacci) {
				s.EXPECT().GetSlice(input).Return([]uint64{0, 1, 1}, errors.New("service failure"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			fibonacci := mock_service.NewMockFibonacci(c)
			testCase.mockBehavior(fibonacci, testCase.inputUser)

			services := &service.Service{Fibonacci: fibonacci}
			handler := NewHandler(services)

			//Test Server
			r := gin.New()
			r.GET("/api/get-fibonacci", handler.getFibonacci)

			// Create Request
			w := httptest.NewRecorder()

			//fmt.Println(values.Encode())
			req := httptest.NewRequest("GET", "/api/get-fibonacci?"+testCase.paramsForRequest, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)

		})
	}

}
