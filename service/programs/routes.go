// It maps HTTP endpoints to their corresponding handler functions.
package programs

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for program-related operations.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.RegisterPrograms)
	router.GET("/all", h.GetPrograms)
}
