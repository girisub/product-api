package handlers

import (
	"net/http"

	"github.com/girisub/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productsResponse
// GetProducts returns the product from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling GET req")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal object", http.StatusInternalServerError)
	}
}
