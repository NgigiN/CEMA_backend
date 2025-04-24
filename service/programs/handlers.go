package programs

import (
	"cema_backend/logging"
	"cema_backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct acts as a bridge between the HTTP layer and the store layer.
type Handler struct {
	store types.ProgramsStore
}

// NewHandler initializes a new Handler instance with the given ProgramsStore.
func NewHandler(store types.ProgramsStore) *Handler {
	return &Handler{store: store}
}

// RegisterPrograms handles the request to register a new program.
func (h *Handler) RegisterPrograms(c *gin.Context) {
	var request types.Programs
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the request
	if request.Name == "" || request.Symptoms == "" || request.Severity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	// Registers the program
	err := h.store.RegisterPrograms(types.Programs{
		Name:     request.Name,
		Symptoms: request.Symptoms,
		Severity: request.Severity,
	})
	if err != nil {
		logging.Error("Failed to Register Program: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error registering program"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Program registered successfully"})
}

// GetPrograms handles the HTTP GET request to fetch all programs.
func (h *Handler) GetPrograms(c *gin.Context) {
	programs, err := h.store.GetPrograms()
	if err != nil {
		logging.Error("Failed to get programs: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching programs"})
		return
	}
	// returns the programs
	c.JSON(http.StatusOK, programs)
}
