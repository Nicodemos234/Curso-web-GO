package routes

import (
	"net/http"

	"github.com/nicodemos234/cursoweb/controllers"
)

func LoadRoots() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
