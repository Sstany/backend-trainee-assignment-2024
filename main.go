package main

import (
	"banney/app/core"
	"banney/app/db"
	"banney/sdk"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	dbURL := os.Getenv(sdk.EnvPostgres)

	dbClient := db.NewClient(dbURL, logger.Named("db"))

	if err := dbClient.Start(); err != nil {
		panic(err)
	}

	host := os.Getenv(sdk.EnvHost)

	server := core.NewServer(host, dbClient, logger.Named("server"))

	server.Start()

	server.Wait()
}
