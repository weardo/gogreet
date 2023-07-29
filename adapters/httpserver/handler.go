package httpserver

import (
	"fmt"
	"go_specs_greet/domain/interactions"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}