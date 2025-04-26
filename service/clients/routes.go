// This file contains the endpoints for the clients service.
package clients

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.RegisterClients)
	router.POST("/program-enroll", h.EnrollClient)
	router.POST("/search", h.SearchClient)
	router.GET("/clients", h.GetAllClients)
	router.POST("/prescription", h.CreatePrescription)
	router.PUT("/prescription", h.UpdatePrescription)
	router.DELETE("/delete", h.DeleteClient)
}
