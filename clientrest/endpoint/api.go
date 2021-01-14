package endpoint

import (
	"encoding/json"
	"fmt"
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"net/http"
)

func (r *router) Create(w http.ResponseWriter, rq *http.Request) {
	var createOrder domain.CreateOrder

	if err := json.NewDecoder(rq.Body).Decode(&createOrder); err != nil {
		r.ResponseWithError(w, http.StatusBadRequest, "error body")
		return
	}

	orderID := r.service.CreateOrder()
	r.Response(w, http.StatusCreated, []byte(fmt.Sprintf(`{"order": "%s"}`, orderID)))
}

func (r *router) Checkout(w http.ResponseWriter, rq *http.Request) {

}

func (r *router) AddToBasket(w http.ResponseWriter, rq *http.Request) {

}

func (r *router) GetTotalAmount(w http.ResponseWriter, rq *http.Request) {

}

func (r *router) RemoveAll(w http.ResponseWriter, rq *http.Request) {

}