package endpoint

import (
	"encoding/json"
	"github.com/ervitis/backend-challenge/clientrest/exceptions"
	"github.com/ervitis/backend-challenge/clientrest/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func (rt *router) Create(_ http.ResponseWriter, r *http.Request) (Responder, error) {
	var createOrder CreateOrder

	if err := json.NewDecoder(r.Body).Decode(&createOrder); err != nil {
		return nil, BadRequest("invalid body")
	}

	orderID := rt.service.CreateOrder(createOrder.UserID)
	return Created(WithBody(map[string]interface{}{"orderId": orderID})), nil
}

func (rt *router) Checkout(_ http.ResponseWriter, r *http.Request) (Responder, error) {
	orderID, err := strconv.Atoi(strings.TrimSpace(mux.Vars(r)["id"]))
	if err != nil || orderID <= 0 {
		return nil, BadRequest("orderId not valid")
	}

	rt.service.Checkout(orderID)
	return Ok(), nil
}

func (rt *router) AddToBasket(_ http.ResponseWriter, r *http.Request) (Responder, error) {
	var input []AddProductToOrder

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, BadRequest("invalid body")
	}

	orderID, err := strconv.Atoi(strings.TrimSpace(mux.Vars(r)["id"]))
	if err != nil || orderID <= 0 {
		return nil, BadRequest("missing or wrong parameter")
	}

	var items []model.Product
	for _, v := range input {
		item := model.Product{
			ProductID: v.ProductID,
			Quantity: v.Quantity,
		}
		items = append(items, item)
	}

	if err = rt.service.AddItemToBasket(orderID, items); err != nil {
		switch err.(type) {
		case exceptions.OrderNotFoundErr:
			return nil, NotFound(err.Error())
		}
	}
	return Ok(), nil
}

func (rt *router) GetTotalAmount(_ http.ResponseWriter, r *http.Request) (Responder, error) {
	orderID, err := strconv.Atoi(strings.TrimSpace(mux.Vars(r)["id"]))
	if err != nil || orderID <= 0 {
		return nil, BadRequest("missing or wrong parameter")
	}

	amount := rt.service.GetTotalAmount(orderID)
	return Ok(WithBody(map[string]interface{}{"amount": amount})), nil
}

func (rt *router) RemoveAll(_ http.ResponseWriter, r *http.Request) (Responder, error) {
	orderID, err := strconv.Atoi(strings.TrimSpace(mux.Vars(r)["id"]))
	if err != nil || orderID <= 0 {
		return nil, BadRequest("missing or wrong parameter")
	}

	if err := rt.service.DeleteOrder(orderID); err != nil {
		return nil, ServiceUnavailable("connection refused")
	}

	return NoContent(), nil
}
