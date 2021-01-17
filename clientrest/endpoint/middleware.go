package endpoint

import "net/http"

type (
	IMiddleware interface {
		HeaderContentTypeJson(http.Handler) http.Handler
	}

	middleware struct{}
)

func NewMiddleware() IMiddleware {
	return &middleware{}
}

func (m *middleware) HeaderContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}