package main_test

import (
	"fmt"
	"go_specs_greet/adapters"
	"go_specs_greet/adapters/grpcserver"
	"go_specs_greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port = "50052"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, driver)
}