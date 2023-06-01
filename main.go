package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pedy4000/logstronaut/receiver"
	"github.com/pedy4000/logstronaut/util"
	"github.com/rs/zerolog/log"
)

func main() {
	// startup message	
	fmt.Println("Starting Logstronaut")

	// set log settings
	util.SetGlobalLogSettings("console", "info")

	// load config file
	config, err := util.LoadConfig(".")
	if err != nil {
		// in case of error, log the error and exit
		log.Fatal().Err(err).Msg("failed to load config file")
	}

	// create receiver service server
	receiverServer, err := receiver.NewServer(config)
	if err != nil {
		// in case of error, log the error and exit
		log.Fatal().Err(err).Msg("failed to create receiver service server")
	}

	go func() {
		// start the server or catch the error if it fails
		err := receiverServer.Start()
		if err != nil {
			// in case of error, log the error and exit
			log.Fatal().Err(err).Msg("failed to start receiver server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	log.Info().Msg("Shutting Down Monshi server gracefully")
}