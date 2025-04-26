// This file contains the endpoints for the clients service.
package clients

import (
	"cema_backend/auth"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	// Public routes
	router.POST("/register", h.RegisterClients)
	router.POST("/search", h.SearchClient)

	// Protected routes
	protected := router.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.POST("/program-enroll", h.EnrollClient)
		protected.GET("/clients", h.GetAllClients)
		protected.POST("/prescription", h.CreatePrescription)
		protected.PUT("/prescription", h.UpdatePrescription)
		protected.DELETE("/delete", h.DeleteClient)
	}
}
