package interactions_test

import (
	"go_specs_greet/domain/interactions"
	"go_specs_greet/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t, 
		specifications.GreetAdapter(interactions.Greet),
	)
}