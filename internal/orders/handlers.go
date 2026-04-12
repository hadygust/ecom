package orders

import (
	"log"
	"net/http"

	"github.com/hadygust/ecom/json"
)

type handler struct {
	svc Service
}

func NewHandler(svc Service) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	var order createOrderParams

	if err := json.Read(r, &order); err != nil {
		// log.Println("called")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(order)

	in_order, err := h.svc.PlaceOrder(r.Context(), order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.Write(w, http.StatusCreated, in_order)
}
