package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/nicodemos234/cursoweb/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		total := r.FormValue("total")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		totalConverted, err := strconv.Atoi(total)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CreateNewProduct(name, description, priceConverted, totalConverted)
	}
	http.Redirect(w, r, "/", 301)
}
