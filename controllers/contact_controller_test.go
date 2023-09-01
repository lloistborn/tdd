package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"tdd-go-fiber/domain"
	"testing"
)

type MockContactUseCase struct {
	mock.Mock
}

func (m *MockContactUseCase) GetContacts() ([]domain.Contact, error) {
	args := m.Called()
	return args.Get(0).([]domain.Contact), args.Error(1)
}

func TestGetContacts(t *testing.T) {
	// define the input and expected output
	var tests = []struct {
		description        string
		route              string
		expectedStatusCode int
		mockData           []domain.Contact
		mockError          error
	}{
		{
			description:        "should return 200 OK",
			route:              "/api/v1/contacts",
			expectedStatusCode: 200,
			mockData: []domain.Contact{
				{ID: 1, FirstName: "John", LastName: "Doe", Phone: "08123456789"},
			},
			mockError: nil,
		},
		{
			description:        "should return 500 Internal Server Error",
			route:              "/api/v1/contacts",
			expectedStatusCode: 500,
			mockData:           []domain.Contact{},
			mockError:          fiber.ErrInternalServerError,
		},
	}

	for _, test := range tests {
		// define mock
		mockContactUseCase := new(MockContactUseCase)
		mockContactUseCase.On("GetContacts").Return(test.mockData, test.mockError)

		// define routes for tests
		app := fiber.New()
		app.Get("/api/v1/contacts", func(c *fiber.Ctx) error {
			return GetContacts(c, mockContactUseCase)
		})

		// define request
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		// define response
		resp, _ := app.Test(req)

		// assert response
		assert.Equal(t, test.expectedStatusCode, resp.StatusCode, test.description)
	}
}
