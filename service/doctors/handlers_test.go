package doctors

import (
	"bytes"
	"cema_backend/types"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockDoctorStore is a mock implementation of the DoctorStore interface.
type MockDoctorStore struct {
	mock.Mock
}

func (m *MockDoctorStore) RegisterDoctors(doctor types.DoctorRegistration) error {
	args := m.Called(doctor)
	return args.Error(0)
}

func (m *MockDoctorStore) LoginDoctor(email, password string) error {
	args := m.Called(email, password)
	return args.Error(0)
}

func TestRegisterDoctors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockDoctorStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.POST("/register", handler.RegisterDoctors)

	// Test case: Successful registration
	mockStore.On("RegisterDoctors", mock.Anything).Return(nil)

	payload := types.DoctorRegistration{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "1234567890",
		Department:  "Cardiology",
		Password:    "password123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "RegisterDoctors", payload)
}

func TestLoginDoctor(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockDoctorStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.POST("/login", handler.LoginDoctor)

	// Test case: Successful login
	mockStore.On("LoginDoctor", "john.doe@example.com", "password123").Return(nil)

	payload := types.DocLogInRequest{
		Email:    "john.doe@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "LoginDoctor", "john.doe@example.com", "password123")
}
