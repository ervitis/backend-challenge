package endpoint

import (
	"github.com/gorilla/mux"
)

type (
	router struct {
		r *mux.Router
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
