package http

import (
	"net/http"
)

type Middleware func(http.ResponseWriter, *http.Request, Next)

type Next func(*http.Request)

type Middlewares []Middleware

func (m Middlewares) Append(middleware Middleware) Middlewares {
	return append(m, middleware)
}

func (m Middlewares) Chain(final http.HandlerFunc) http.HandlerFunc {
	result := final

	for i := len(m) - 1; i > -1; i -= 1 {
		handler := result

		middleware := m[i]

		result = func(w http.ResponseWriter, r *http.Request) {
			next := func(r *http.Request) {
				handler(w, r)
			}

			middleware(w, r, next)
		}
	}

	return result
}
