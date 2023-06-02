package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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

	dbConn, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration("file://db/migration", config.DBSource)

	// create receiver service server
	receiverServer, err := receiver.NewServer(config, dbConn)
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

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}
