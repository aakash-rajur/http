package http

import "net/http"

func (router *Router) Get(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodGet, pattern, handler)
}

func (router *Router) GetFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Get(pattern, handlerFunc)
}

func (router *Router) Post(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodPost, pattern, handler)
}

func (router *Router) PostFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Post(pattern, handlerFunc)
}

func (router *Router) Put(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodPut, pattern, handler)
}

func (router *Router) PutFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Put(pattern, handlerFunc)
}

func (router *Router) Patch(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodPatch, pattern, handler)
}

func (router *Router) PatchFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Patch(pattern, handlerFunc)
}

func (router *Router) Delete(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodDelete, pattern, handler)
}

func (router *Router) DeleteFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Delete(pattern, handlerFunc)
}

func (router *Router) Head(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodHead, pattern, handler)
}

func (router *Router) HeadFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Head(pattern, handlerFunc)
}

func (router *Router) Options(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodOptions, pattern, handler)
}

func (router *Router) OptionsFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Options(pattern, handlerFunc)
}

func (router *Router) Trace(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodTrace, pattern, handler)
}

func (router *Router) TraceFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Trace(pattern, handlerFunc)
}

func (router *Router) Connect(pattern string, handler http.Handler) {
	router.HandleMethod(http.MethodConnect, pattern, handler)
}

func (router *Router) ConnectFunc(pattern string, handlerFunc http.HandlerFunc) {
	router.Get(pattern, handlerFunc)
}
