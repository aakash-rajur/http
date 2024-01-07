package http

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewares_Append(t *testing.T) {
	middlewares := make(Middlewares, 0)

	m1 := func(w http.ResponseWriter, r *http.Request, next Next) {
		next(r)
	}

	m2 := func(w http.ResponseWriter, r *http.Request, next Next) {
		next(r)
	}

	middlewares = middlewares.Append(m1)

	assert.Equalf(t, 1, len(middlewares), "want len(middlewares) == 1, got %v", len(middlewares))

	middlewares = middlewares.Append(m2)

	assert.Equalf(t, 2, len(middlewares), "want len(middlewares) == 2, got %v", len(middlewares))
}

func TestMiddlewares_Chain(t *testing.T) {
	middlewares := make(Middlewares, 0)

	mm1, mm2 := NewMockMiddleware(), NewMockMiddleware()

	middlewares = middlewares.Append(mm1.Middleware())

	middlewares = middlewares.Append(mm2.Middleware())

	final := &MockHandler{}

	final.On(
		"ServeHTTP",
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("*http.Request"),
	)

	handler := middlewares.Chain(final.ServeHTTP)

	assert.NotNil(t, handler)

	mrw := httptest.NewRecorder()

	mr := httptest.NewRequest(http.MethodGet, "/", nil)

	handler(mrw, mr)

	mm1.AssertCalled(t, "Run", mrw, mr, mock.AnythingOfType("Next"))

	mm1.AssertNumberOfCalls(t, "Run", 1)

	mm2.AssertCalled(t, "Run", mrw, mr, mock.AnythingOfType("Next"))

	mm2.AssertNumberOfCalls(t, "Run", 1)

	final.AssertCalled(t, "ServeHTTP", mrw, mr)

	final.AssertNumberOfCalls(t, "ServeHTTP", 1)
}
