package clients

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

// MockClientStore is a mock implementation of the ClientStore interface.
type MockClientStore struct {
	mock.Mock
}

// DeleteClient implements types.ClientStore.
func (m *MockClientStore) DeleteClient(phonenumber string) error {
	args := m.Called(phonenumber)
	return args.Error(0)
}

// GetAllClients implements types.ClientStore.
func (m *MockClientStore) GetAllClients() ([]types.Client, error) {
	args := m.Called()
	return args.Get(0).([]types.Client), args.Error(1)
}

// RegisterClients implements types.ClientStore.
func (m *MockClientStore) RegisterClients(client types.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

// SearchClient implements types.ClientStore.
func (m *MockClientStore) SearchClient(phonenumber string) (types.ClientResponse, error) {
	args := m.Called(phonenumber)
	return args.Get(0).(types.ClientResponse), args.Error(1)
}

// UpdateClient implements types.ClientStore.
func (m *MockClientStore) UpdateClient(client types.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientStore) EnrollClient(email, programID string) error {
	args := m.Called(email, programID)
	return args.Error(0)
}

func TestEnrollClient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockClientStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.POST("/enroll", handler.EnrollClient)

	// Test case: Successful enrollment
	mockStore.On("EnrollClient", "client@example.com", "program123").Return(nil)

	payload := map[string]string{
		"email":     "client@example.com",
		"programID": "program123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/enroll", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "EnrollClient", "client@example.com", "program123")
}

func TestSearchClient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockClientStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.POST("/search", handler.SearchClient)

	// Test case: Successful search
	mockClient := types.ClientResponse{
		ID:               1,
		FirstName:        "John",
		LastName:         "Doe",
		PhoneNumber:      "1234567890",
		Height:           180.5,
		Weight:           75.0,
		Age:              30,
		Programs:         []types.Programs{},
		EmergencyContact: "Jane Doe",
		EmergencyNumber:  "0987654321",
	}
	// Add mock behavior for SearchClient
	mockStore.On("SearchClient", "1234567890").Return(mockClient, nil)

	payload := map[string]string{
		"phonenumber": "1234567890",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/search", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "SearchClient", "1234567890")
}
func TestGetAllClients(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockClientStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.GET("/getall", handler.GetAllClients)

	// Test case: Successful retrieval of all clients
	mockClients := []types.Client{
		{
			ID:               1,
			FirstName:        "John",
			LastName:         "Doe",
			PhoneNumber:      "1234567890",
			EmergencyContact: "Jane Doe",
			EmergencyNumber:  "0987654321",
		},
	}
	mockStore.On("GetAllClients").Return(mockClients, nil)

	req, _ := http.NewRequest(http.MethodGet, "/getall", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "GetAllClients")
}
func TestRegisterClients(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockClientStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.POST("/register", handler.RegisterClients)

	// Test case: Successful registration
	mockStore.On("SearchClient", "0115491173").Return(types.ClientResponse{}, nil)
	mockStore.On("RegisterClients", mock.Anything).Return(nil)

	payload := map[string]interface{}{
		"firstname":         "John",
		"lastname":          "Doe",
		"phonenumber":       "0115491173",
		"age":               10,
		"height":            180,
		"weight":            80,
		"emergency_contact": "father",
		"emergency_number":  "0987654321",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "SearchClient", "0115491173")
	mockStore.AssertCalled(t, "RegisterClients", mock.Anything)
}
