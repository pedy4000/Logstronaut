package receiver

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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
	myfile, err := os.OpenFile(server.config.StorageDestination, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error().Err(err).Msg("error while accessing target file")
		ctx.JSON(500, errorResponse(fmt.Errorf("error while accessing target file")))
		return
	}
	defer myfile.Close()

	// Write the string to the file
	_, err = myfile.WriteString(request.Message + "\n")
	if err != nil {
		// if the RPC call returns an error, return a 500 error
		log.Error().Err(err).Msg("error while writing string")
		ctx.JSON(500, errorResponse(fmt.Errorf("error in saving message")))
		return
	}

	// return a 200 OK response
	ctx.JSON(200, gin.H{"status": "ok"})
}
