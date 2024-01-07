package http

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

func NewMockMiddleware() *MockMiddleware {
	mm := new(MockMiddleware)

	mm.On(
		"Run",
		mock.AnythingOfType("*http.ResponseWriter"),
		mock.AnythingOfType("*http.Request"),
		mock.AnythingOfType("Next"),
	)

	return mm
}

type MockMiddleware struct {
	mock.Mock
}

func (mm *MockMiddleware) Run(rw http.ResponseWriter, r *http.Request, next Next) {
	mm.Called(rw, r, next)

	next(r)
}

func (mm *MockMiddleware) Middleware() Middleware {
	return mm.Run
}

func NewMockHandler(handler http.Handler) *MockHandler {
	mh := new(MockHandler)

	mh.handler = handler

	mh.On(
		"ServeHTTP",
		mock.AnythingOfType("*http.ResponseWriter"),
		mock.AnythingOfType("*http.Request"),
	)

	return mh
}

type MockHandler struct {
	mock.Mock
	handler http.Handler
}

func (mh *MockHandler) Handler() http.Handler {
	return mh
}

func (mh *MockHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	mh.Called(rw, r)

	if mh.handler != nil {
		mh.handler.ServeHTTP(rw, r)

		return
	}

	rw.WriteHeader(http.StatusOK)
}
