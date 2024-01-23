package handlers

import (
	"net/http"
	"strconv"

	"github.com/Wise-Wizard/GoLang-Projects/product-api/data"
	"github.com/gorilla/mux"
)

// getProducts returns the products from the data store
func (p *Products) ListAllProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	prods := data.GetProducts()

	// serialize the list to JSON
	err := data.ToJSON(prods, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) ListSingleProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("[ERROR] converting string to int", err)
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("id", id)
	prod, err := data.GetProduct(id)
	if err != nil {
		p.l.Println("[ERROR] fetching product", err)
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	p.l.Println("prod", prod)
	// serialize the list to JSON
	err = data.ToJSON(prod, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
