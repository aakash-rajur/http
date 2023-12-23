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

	m1 := mockMiddleware{}

	m1.On(
		"Run",
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("*http.Request"),
		mock.AnythingOfType("Next"),
	)

	m2 := mockMiddleware{}

	m2.On(
		"Run",
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("*http.Request"),
		mock.AnythingOfType("Next"),
	)

	middlewares = middlewares.Append(m1.Run)

	middlewares = middlewares.Append(m2.Run)

	final := &mockHandler{}

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

	m1.AssertCalled(t, "Run", mrw, mr, mock.AnythingOfType("Next"))

	m1.AssertNumberOfCalls(t, "Run", 1)

	m2.AssertCalled(t, "Run", mrw, mr, mock.AnythingOfType("Next"))

	m2.AssertNumberOfCalls(t, "Run", 1)

	final.AssertCalled(t, "ServeHTTP", mrw, mr)

	final.AssertNumberOfCalls(t, "ServeHTTP", 1)
}
