package clients

import (
	"cema_backend/logging"
	"cema_backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct contains the store for client operations
type Handler struct {
	store types.ClientStore
}

// NewHandler initializes a new Handler for the clients service
func NewHandler(store types.ClientStore) *Handler {
	return &Handler{store: store}
}

// RegisterClients handles the registration of a new client
func (h *Handler) RegisterClients(c *gin.Context) {
	var request types.Client
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Validate the request
	if request.FirstName == "" || request.LastName == "" || request.PhoneNumber == "" || request.Height == 0 || request.Weight == 0 || request.Age == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	if request.EmergencyContact == "" || request.EmergencyNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Emergency contact and number are required"})
		return
	}

	// Check if the client already exists
	_, err := h.store.SearchClient(request.PhoneNumber)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client already exists"})
		return
	} else if err.Error() != "client does not exist" {
		logging.Error("Error searching for client: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for client"})
		return
	}

	// Register the client
	err = h.store.RegisterClients(types.Client{
		FirstName:        request.FirstName,
		LastName:         request.LastName,
		PhoneNumber:      request.PhoneNumber,
		Height:           request.Height,
		Weight:           request.Weight,
		Age:              request.Age,
		EmergencyContact: request.EmergencyContact,
		EmergencyNumber:  request.EmergencyNumber,
	})

	if err != nil {
		logging.Error("Failed to Register Client: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error registering client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client registered successfully"})
}

// enrollClient handles the enrollment of a client in a program
func (h *Handler) EnrollClient(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		ProgramName string `json:"programName" binding:"required"`
	}

	// Bind the request payload
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the request
	if request.PhoneNumber == "" || request.ProgramName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhoneNumber and ProgramName are required"})
		return
	}

	// Enroll the client
	err := h.store.EnrollClient(request.PhoneNumber, request.ProgramName)
	if err != nil {
		logging.Error("Failed to Enroll Client: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error enrolling client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client enrolled successfully"})
}

// SearchClient handles the search for a client by email
func (h *Handler) SearchClient(c *gin.Context) {
	var request struct {
		Phonenumber string `json:"phonenumber" binding:"required"`
	}

	// Bind the request payload
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the request
	if request.Phonenumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	// Search for the client
	client, err := h.store.SearchClient(request.Phonenumber)
	if err != nil {
		logging.Error("Failed to Search Client: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client not Found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// GetAllClients handles the retrieval of all clients
func (h *Handler) GetAllClients(c *gin.Context) {
	clients, err := h.store.GetAllClients()
	if err != nil {
		logging.Error("Failed to Get All Clients: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving clients"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// DeleteClient handles the deletion of a client by phone number
func (h *Handler) DeleteClient(c *gin.Context) {
	var request struct {
		Phonenumber string `json:"phonenumber" binding:"required"`
	}

	// Validate the request
	if request.Phonenumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Delete the client
	// Assuming the phone number is unique for each client
	err := h.store.DeleteClient(request.Phonenumber)
	if err != nil {
		logging.Error("Failed to Delete Client: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

// CreatePrescription handles the creation of a new prescription
func (h *Handler) CreatePrescription(c *gin.Context) {
	var request types.Prescription
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.store.CreatePrescription(request)
	if err != nil {
		logging.Error("Failed to create prescription: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating prescription"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prescription created successfully"})
}

// UpdatePrescription handles the updating of an existing prescription
func (h *Handler) UpdatePrescription(c *gin.Context) {
	var request types.Prescription
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.store.UpdatePrescription(request)
	if err != nil {
		logging.Error("Failed to update prescription: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating prescription"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prescription updated successfully"})
}
