package integrationtests

import (
	"banney/app/core"
	"banney/app/db"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var databaseUrl string

var testDBClient *db.Client

func TestMain(m *testing.M) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	resource, pool := PreparDockerPostgres()

	testDBClient = db.NewClient(databaseUrl, logger.Named("db"))

	if err != nil {
		panic(err)
	}

	host := "localhost:8090"

	server := core.NewServer(host, testDBClient, logger.Named("server"))

	server.Start()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	code := m.Run()

	server.Stop(ctx)
	server.Wait()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func PreparDockerPostgres() (*dockertest.Resource, *dockertest.Pool) {
	var db *sql.DB

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_PASSWORD=password",
			"POSTGRES_USER=username",
			"POSTGRES_DB=avito",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl = fmt.Sprintf("postgres://username:password@%s/avito?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(600) // Tell docker to hard kill the container in 10 minutes

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return resource, pool
}
