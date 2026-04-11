package products

import (
	"net/http"

	"github.com/hadygust/ecom/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// Call the service
	// Return JSON

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)

}
