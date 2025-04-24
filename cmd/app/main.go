// Entry point for the application
// This file contains the main function that initializes and runs the API server.
package app

import (
	"cema_backend/logging"
	"cema_backend/service/clients"
	"cema_backend/service/doctors"
	"cema_backend/service/programs"
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// This struct represents the API server with its address and database connection.
type APIServer struct {
	addr string
	db   *sql.DB
}

// Initializes a new API server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server and sets up the routes
func (s *APIServer) Run() error {
	router := gin.Default()

	// base case for the server to check if the server is reachable
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// CORS configuration, allows all for now
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Register Doctor routes
	// Each service has its own store and handler but they all use the same database connection
	doctorStore := doctors.NewStore(s.db)
	doctorHandler := doctors.NewHandler(doctorStore)
	doctorRoutes := router.Group("/doctors")
	doctorHandler.RegisterRoutes(doctorRoutes)

	// Register Programs routes
	programStore := programs.NewStore(s.db)
	programHandler := programs.NewHandler(programStore)
	programRoutes := router.Group("/programs")
	programHandler.RegisterRoutes(programRoutes)

	//Register Client routes
	clientStore := clients.NewStore(s.db)
	clientHandler := clients.NewHandler(clientStore)
	clientRoutes := router.Group("/clients")
	clientHandler.RegisterRoutes(clientRoutes)

	logging.Info("Listening on port: " + s.addr)
	return router.Run(s.addr)
}
