package main

import (
	"go_specs_greet"
	"net/http"
)


func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	
	http.ListenAndServe(":8030", handler)
}