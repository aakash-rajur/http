package http

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type mockMiddleware struct {
	mock.Mock
}

func (mm *mockMiddleware) Run(rw http.ResponseWriter, r *http.Request, next Next) {
	mm.Called(rw, r, next)

	next(r)
}

type mockHandler struct {
	mock.Mock
}

func (mh *mockHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	mh.Called(rw, r)

	rw.WriteHeader(http.StatusOK)
}
