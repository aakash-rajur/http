package http

import (
	"errors"
	"github.com/aakash-rajur/http/register"
	"net/http"
	"sync"
)

func NewRouter() *Router {
	mux := &Router{
		middlewares: make(Middlewares, 0),
		register:    register.NewRegister(),
		notFound:    http.NotFoundHandler(),
	}

	return mux
}

type Router struct {
	mu          sync.RWMutex
	middlewares Middlewares
	next        http.HandlerFunc
	register    register.Register
	notFound    http.Handler
}

func (router *Router) HandleMethod(method, pattern string, handler http.Handler) {
	router.mu.Lock()

	defer router.mu.Unlock()

	path := pathWithMethod(method, pattern)

	router.register = router.register.Add(path, handler)
}

func (router *Router) HandleMethodFunc(method, pattern string, handlerFunc http.HandlerFunc) {
	router.HandleMethod(method, pattern, handlerFunc)
}

func (router *Router) Handle(pattern string, handler http.Handler) {
	router.HandleMethod(":http_method", pattern, handler)
}

func (router *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Handle(pattern, handlerFunc)
}

func (router *Router) NotFound(handler http.Handler) {
	router.mu.Lock()

	defer router.mu.Unlock()

	router.notFound = handler
}

func (router *Router) Use(middleware Middleware) {
	router.middlewares = router.middlewares.Append(middleware)

	router.next = router.middlewares.Chain(router.serve)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hw := &ResponseWriter{ResponseWriter: w}

	if router.next == nil {
		router.next = router.serve
	}

	router.next(hw, r)
}

func (router *Router) serve(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	router.mu.RLock()

	defer router.mu.RUnlock()

	path := pathWithMethod(r.Method, r.URL.Path)

	entry, params, err := router.register.Find(path)

	if err == nil {
		pr := r.WithContext(params.WithinContext(r.Context()))

		entry.Handler.ServeHTTP(w, pr)

		return
	}

	if !errors.Is(err, register.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	router.notFound.ServeHTTP(w, r)
}
