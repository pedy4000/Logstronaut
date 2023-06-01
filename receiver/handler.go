package receiver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/pedy4000/logstronaut/db/sqlc"
	"github.com/rs/zerolog/log"
)

// saveMessageRequest is the request body for the "/save" route.
type saveMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

// saveMessage handles the "/save" route.
// it saves the message to the storage destination
func (server *Server) saveMessage(ctx *gin.Context) {
	// bind the request body to the request struct
	var request saveMessageRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// if the request body is invalid, return a 400 error
		ctx.JSON(400, errorResponse(err))
		return
	}

	// saves the message the remote procedure
	msg, err := db.New(server.dbConn).StoreMessage(ctx, request.Message)
	if err != nil {
		// if the RPC call returns an error, return a 500 error
		log.Error().Err(err).Msg("error while saving message content")
		ctx.JSON(500, errorResponse(fmt.Errorf("error while saving message")))
		return
	}

	// return a 200 OK response
	ctx.JSON(200, msg)
}

// health handles the "/health" route.
func (server *Server) health(ctx *gin.Context) {
	// return a 200 OK response
	if err := server.dbConn.Ping(); err != nil {
		ctx.JSON(503, errorResponse(err))
		return
	}
	ctx.JSON(200, gin.H{"status": "healthy"})
}
