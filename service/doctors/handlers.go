package doctors

import (
	"cema_backend/logging"
	"cema_backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct acts as a bridge between the HTTP layer and the store layer.
type Handler struct {
	store types.DoctorStore
}

// NewHandler initializes a new Handler instance with the given DoctorStore.
func NewHandler(store types.DoctorStore) *Handler {
	return &Handler{store: store}
}

// RegisterDoctors handles the request to register a new doctor.
func (h *Handler) RegisterDoctors(c *gin.Context) {
	var request types.DoctorRegistration
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Validate the request
	if request.FirstName == "" || request.LastName == "" || request.PhoneNumber == "" || request.Department == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Register the doctor
	err := h.store.RegisterDoctors(types.DoctorRegistration{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Department:  request.Department,
		Password:    request.Password,
	})
	if err != nil {
		logging.Error("Failed to Register Doctor: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error registering doctor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor registered successfully"})
}

// LoginDoctor handles the request to log in a doctor.
func (h *Handler) LoginDoctor(c *gin.Context) {
	var request types.DocLogInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Validate the request
	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	// Log in the doctor
	err := h.store.LoginDoctor(request.Email, request.Password)
	if err != nil {
		logging.Error("Failed to Login Doctor: " + err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Doctor logged in successfully"})
}
