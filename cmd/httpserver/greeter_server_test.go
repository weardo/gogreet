package main_test

import (
	"context"
	"go_specs_greet"
	"go_specs_greet/specifications"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context: "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8030:8030"},
		WaitingFor: wait.ForHTTP("/").WithPort("8030"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started: true,
	})

	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client {
		Timeout: 1 * time.Second,
	}

	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8030", Client: &client}
	specifications.GreetSpecification(t, driver)
}