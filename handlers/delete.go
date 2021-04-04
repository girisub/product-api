package handlers

import (
	"net/http"
	"strconv"

	"github.com/girisub/product-api/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
// 201: noContent

// DeleteProduct deletes a product from the data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handling DELETE req")

	err := data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		http.Error(rw, data.ErrorProductNotFound.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
