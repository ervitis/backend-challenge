package endpoint

import (
	"github.com/ervitis/backend-challenge/clientrest/domain"
	"github.com/gorilla/mux"
)

type (
	router struct {
		r *mux.Router

		service domain.IBasket
	}

	IRouter interface {
		GetRouter() *mux.Router
	}
)

func (r *router) GetRouter() *mux.Router {
	return r.r
}

func NewRouter() IRouter {
	return &router{r: mux.NewRouter()}
}
