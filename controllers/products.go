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
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		total := r.FormValue("total")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int.", err)
		}

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço para Float64.", err)
		}

		totalConverted, err := strconv.Atoi(total)
		if err != nil {
			log.Println("Erro na conversão do Quantidade para int.", err)
		}

		models.UpdateProduct(idConverted, name, description, priceConverted, totalConverted)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
