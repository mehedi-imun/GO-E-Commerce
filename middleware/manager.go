package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleWares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleWares: make([]Middleware, 0),
	}

}

func (mangr *Manager) Use(middlewares ...Middleware) {
	mangr.globalMiddleWares = append(mangr.globalMiddleWares, middlewares...)
}

func (mangr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleWare := range mangr.globalMiddleWares {
		n = globalMiddleWare(n)
	}
	return n
}
