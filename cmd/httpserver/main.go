package main

import (
	"go_specs_greet/adapters/httpserver"
	"net/http"
)


func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	
	http.ListenAndServe(":8030", handler)
}