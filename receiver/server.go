// package receiver implements the HTTP handler for the logstronaut.
// it receives messages over HTTP and saves them to the filesystem.
// HTTP requests are served by gin.
package receiver

import (
	"database/sql"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/pedy4000/logstronaut/util"
	"github.com/rs/zerolog/log"

	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// Server serves HTTP requests for the receiver
type Server struct {
	config util.Config // configuration for the server
	dbConn *sql.DB
	router *gin.Engine // gin router
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, dbConn *sql.DB) (*Server, error) {
	// create a new server object
	server := &Server{
		config: config,
		dbConn: dbConn,
	}

	// set up gin router
	server.setupRouter()
	return server, nil
}

// setupRouter sets up the gin router and adds the routes.
// it also adds the logger and recovery middleware.
// path "/save" is used to save messages to the filesystem.
func (server *Server) setupRouter() {
	// create a new gin router in order to avoid default logging middleware
	router := gin.New()
	// add the zerolog compatible logger middleware
	// here we use ginzerolog package to create a middleware which uses zerolog as the logger
	router.Use(ginzerolog.Logger("gin"))
	// add the recovery middleware to recover from panics
	// and return a 500 error
	router.Use(gin.Recovery())

	// add metrics middleware
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	// add the route for saveing messages
	router.POST("/save", server.saveMessage)

	// add the routes for liveness and readiness
	router.GET("/health", server.health)
	router.GET("/readiness", server.health)

	// set the router to the server object
	server.router = router
}

// Start runs the Gin HTTP server
// it will return an error if the server fails to start, otherwise it will block the thread
// and wait for the server to stop. in case of peacefull stop, it will return nil.
func (server *Server) Start() error {
	log.Info().Str("ser_name", "receiver").Msg("Starting gin server...")

	recordMetrics()
	log.Info().Msg("recording prom metrics")

	return server.router.Run(server.config.ReceiverAddress)
}

// errorResponse is a helper function to create a gin.H map with the error key and the error message.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
