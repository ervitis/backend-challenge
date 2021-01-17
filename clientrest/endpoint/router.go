package endpoint

import (
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	router struct {
		r *mux.Router

		service domain.IBasket
	}

	IRouter interface {
		GetRouter() *mux.Router
		LoadApi()
	}
)

func (rt *router) ApiV1() {
	middlewares := NewMiddleware()

	rt.r.Use(middlewares.HeaderContentTypeJson)

	v1 := rt.r.PathPrefix("/v1").Subrouter()

	v1.Use([]mux.MiddlewareFunc{}...)
	v1.HandleFunc("/basket", WithError(rt.Create)).Methods(http.MethodPost)
	v1.HandleFunc("/basket/{id}/checkout", WithError(rt.Checkout)).Methods(http.MethodPost)
	v1.HandleFunc("/basket/{id}", WithError(rt.AddToBasket)).Methods(http.MethodPut)
	v1.HandleFunc("/basket/{id}/amount", WithError(rt.GetTotalAmount)).Methods(http.MethodGet)
	v1.HandleFunc("/basket/{id}", WithError(rt.RemoveAll)).Methods(http.MethodDelete)
}

func (rt *router) GetRouter() *mux.Router {
	return rt.r
}

func (rt *router) LoadApi() {
	rt.ApiV1()
}

func NewRouter(service domain.IBasket) IRouter {
	return &router{r: mux.NewRouter(), service: service}
}
