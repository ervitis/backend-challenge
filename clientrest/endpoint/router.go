package endpoint

import (
	"encoding/json"
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type (
	router struct {
		r *mux.Router

		service domain.IBasket
	}

	IRouter interface {
		GetRouter() *mux.Router
	}

	errorResponse struct {
		Code    int    `json:"code"`
		Date    string `json:"date"`
		Message string `json:"message"`
	}
)

func (r *router) ResponseWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("content-type", "application/json")

	response, _ := json.Marshal(&errorResponse{Message: message, Code: code, Date: time.Now().UTC().Format(time.RFC3339)})

	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func (r *router) Response(w http.ResponseWriter, code int, data ...interface{}) {
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(code)
	if len(data) == 0 {
		return
	}

	response, _ := json.Marshal(data[0])
	_, _ = w.Write(response)
}

func (r *router) ApiV1() {
	v1 := r.r.PathPrefix("/v1").Subrouter()

	v1.Use([]mux.MiddlewareFunc{}...)
	v1.HandleFunc("/basket", r.Create).Methods(http.MethodPost)
	v1.HandleFunc("/basket/{id}/checkout", r.Checkout).Methods(http.MethodPost)
	v1.HandleFunc("/basket/{id}", r.AddToBasket).Methods(http.MethodPut)
	v1.HandleFunc("/basket/{id}/amount", r.GetTotalAmount).Methods(http.MethodGet)
	v1.HandleFunc("/basket/{id}", r.RemoveAll).Methods(http.MethodDelete)
}

func (r *router) GetRouter() *mux.Router {
	return r.r
}

func NewRouter() IRouter {
	return &router{r: mux.NewRouter()}
}
