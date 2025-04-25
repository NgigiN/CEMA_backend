package programs

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

// MockProgramsStore is a mock implementation of the ProgramsStore interface.
type MockProgramsStore struct {
	mock.Mock
}

func (m *MockProgramsStore) RegisterPrograms(program types.Programs) error {
	args := m.Called(program)
	return args.Error(0)
}

func (m *MockProgramsStore) GetPrograms() ([]types.Programs, error) {
	args := m.Called()
	return args.Get(0).([]types.Programs), args.Error(1)
}

func TestRegisterPrograms(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockProgramsStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.POST("/register", handler.RegisterPrograms)

	// Test case: Successful program registration
	mockStore.On("RegisterPrograms", mock.Anything).Return(nil)

	payload := types.Programs{
		Name:     "Program A",
		Symptoms: "Symptom A, Symptom B",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	mockStore.AssertCalled(t, "RegisterPrograms", payload)
}

func TestGetPrograms(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockProgramsStore)
	handler := NewHandler(mockStore)

	router := gin.Default()
	router.GET("/get", handler.GetPrograms)

	// Test case: Successful retrieval of programs
	mockPrograms := []types.Programs{
		{Name: "Program A", Symptoms: "Symptom A"},
		{Name: "Program B", Symptoms: "Symptom B"},
	}
	mockStore.On("GetPrograms").Return(mockPrograms, nil)

	req, _ := http.NewRequest(http.MethodGet, "/get", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	var response []types.Programs
	json.Unmarshal(resp.Body.Bytes(), &response)
	require.Equal(t, mockPrograms, response)
	mockStore.AssertCalled(t, "GetPrograms")
}
