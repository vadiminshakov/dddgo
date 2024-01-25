package testing

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

//nolint:exhaustruct
func StartTestPostgres(ctx context.Context) (
	*sql.DB, func(context.Context) error, error,
) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, dbContainer.Terminate, errors.Wrap(err, "failed to start Postgres test container")
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		return nil, dbContainer.Terminate, errors.Wrap(err, "failed to get host of the Postgres test container")
	}

	port, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, dbContainer.Terminate, errors.Wrap(err, "failed to get port of the Postgres test container")
	}

	dbURI := fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb", host, port.Port())

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, dbContainer.Terminate, errors.Wrap(err, "failed to open Postgres test container")
	}

	return db, dbContainer.Terminate, nil
}
