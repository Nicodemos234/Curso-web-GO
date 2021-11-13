package main

import (
	"net/http"

	"github.com/nicodemos234/cursoweb/routes"
)

func main() {
	routes.LoadRoots()
	http.ListenAndServe(":8000", nil)
}
