// This file maps the HTTP endpoints to handler functions in the doctors service.
package doctors

import "github.com/gin-gonic/gin"

// RegisterRoutes registers the routes for doctor-related operations.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.RegisterDoctors)
	router.POST("/login", h.LoginDoctor)
}
