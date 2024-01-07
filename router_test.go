package http

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	assert.NotNilf(t, router, "router should not be nil")

	assert.Equalf(t, 0, len(router.middlewares), "middlewares should be empty")

	assert.Equalf(t, 0, len(router.register), "register should be empty")

	assert.NotNilf(t, router.register, "register should not be nil")
}

func TestRouter_HandleMethod(t *testing.T) {
	t.Parallel()

	type args struct {
		method  string
		pattern string
		handler http.Handler
	}

	tests := []struct {
		name string
		args []args
	}{
		{
			name: "should add a mockHandler",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: new(MockHandler),
				},
			},
		},
		{
			name: "should add multiple handlers to same route",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					method:  "POST",
					pattern: "/",
					handler: new(MockHandler),
				},
			},
		},
		{
			name: "should add a mockHandler to different routes",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					method:  "GET",
					pattern: "/users",
					handler: new(MockHandler),
				},
			},
		},
		{
			name: "should add multiple handlers to different routes",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					method:  "POST",
					pattern: "/",
					handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}),
				},
				{
					method:  "GET",
					pattern: "/users",
					handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}),
				},
				{
					method:  "POST",
					pattern: "/users",
					handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewRouter()

			for _, arg := range test.args {
				router.HandleMethod(arg.method, arg.pattern, arg.handler)
			}

			assert.Equalf(t, len(test.args), len(router.register), "register should have %d entries", len(test.args))
		})
	}
}

func TestRouter_HandleMethodFunc(t *testing.T) {
	t.Parallel()

	type args struct {
		method  string
		pattern string
		handler http.HandlerFunc
	}

	tests := []struct {
		name string
		args []args
	}{
		{
			name: "should add a mockHandler",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
		{
			name: "should add multiple handlers to same route",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					method:  "POST",
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
		{
			name: "should add a mockHandler to different routes",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					method:  "GET",
					pattern: "/users",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
		{
			name: "should add multiple handlers to different routes",
			args: []args{
				{
					method:  "GET",
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					method:  "POST",
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					method:  "GET",
					pattern: "/users",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					method:  "POST",
					pattern: "/users",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewRouter()

			for _, arg := range test.args {
				router.HandleMethodFunc(arg.method, arg.pattern, arg.handler)
			}

			assert.Equalf(t, len(test.args), len(router.register), "register should have %d entries", len(test.args))
		})
	}
}

func TestRouter_Handle(t *testing.T) {
	t.Parallel()

	type args struct {
		pattern string
		handler http.Handler
	}

	tests := []struct {
		name string
		args []args
	}{
		{
			name: "should add a mockHandler",
			args: []args{
				{
					pattern: "/",
					handler: new(MockHandler),
				},
			},
		},
		{
			name: "should add multiple handlers to same route",
			args: []args{
				{
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					pattern: "/",
					handler: new(MockHandler),
				},
			},
		},
		{
			name: "should add a mockHandler to different routes",
			args: []args{
				{
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					pattern: "/users",
					handler: new(MockHandler),
				},
			},
		},
		{
			name: "should add multiple handlers to different routes",
			args: []args{
				{
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					pattern: "/",
					handler: new(MockHandler),
				},
				{
					pattern: "/users",
					handler: new(MockHandler),
				},
				{
					pattern: "/users",
					handler: new(MockHandler),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewRouter()

			for _, arg := range test.args {
				router.Handle(arg.pattern, arg.handler)
			}

			assert.Equalf(t, len(test.args), len(router.register), "register should have %d entries", len(test.args))
		})
	}
}

func TestRouter_HandleFunc(t *testing.T) {
	t.Parallel()

	type args struct {
		pattern string
		handler http.HandlerFunc
	}

	tests := []struct {
		name string
		args []args
	}{
		{
			name: "should add a mockHandler",
			args: []args{
				{
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
		{
			name: "should add multiple handlers to same route",
			args: []args{
				{
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
		{
			name: "should add a mockHandler to different routes",
			args: []args{
				{
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					pattern: "/users",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
		{
			name: "should add multiple handlers to different routes",
			args: []args{
				{
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					pattern: "/",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					pattern: "/users",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
				{
					pattern: "/users",
					handler: func(writer http.ResponseWriter, request *http.Request) {},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewRouter()

			for _, arg := range test.args {
				router.HandleFunc(arg.pattern, arg.handler)
			}

			assert.Equalf(t, len(test.args), len(router.register), "register should have %d entries", len(test.args))
		})
	}
}

func TestRouter_NotFound(t *testing.T) {
	router := NewRouter()

	handler := new(MockHandler)

	router.NotFound(handler)

	assert.Equalf(t, handler, router.notFound, "notFound should be %v", handler)
}

func TestRouter_Use(t *testing.T) {
	t.Parallel()

	type args struct {
		middlewares Middlewares
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "should add a middleware",
			args: args{
				middlewares: Middlewares{NewMockMiddleware().Middleware()},
			},
		},
		{
			name: "should add multiple middlewares",
			args: args{
				middlewares: Middlewares{
					NewMockMiddleware().Middleware(),
					NewMockMiddleware().Middleware(),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewRouter()

			for _, middleware := range test.args.middlewares {
				router.Use(middleware)
			}

			assert.Equalf(t, len(test.args.middlewares), len(router.middlewares), "middlewares should have %d entries", len(test.args.middlewares))

			length := len(test.args.middlewares)

			for i, middleware := range test.args.middlewares {
				assert.Equalf(
					t,
					fmt.Sprintf("%v", middleware),
					fmt.Sprintf("%v", router.middlewares[length-i-1]),
					"middlewares[%d] should be %v",
					i,
					middleware,
				)
			}
		})
	}
}

func TestRouter_ServeHTTP(t *testing.T) {
	t.Parallel()

	type args struct {
		method      string
		pattern     string
		mockHandler *MockHandler
		expectation string
	}

	okHandler := func(expectation string) *MockHandler {
		return NewMockHandler(
			http.HandlerFunc(
				func(writer http.ResponseWriter, request *http.Request) {
					writer.WriteHeader(http.StatusOK)

					_, _ = writer.Write([]byte(expectation))
				},
			),
		)
	}

	tests := []struct {
		name            string
		args            []args
		mockMiddlewares []*MockMiddleware
	}{
		{
			name:            "should serve a mockHandler",
			mockMiddlewares: []*MockMiddleware{},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
			},
		},
		{
			name:            "should serve multiple handlers",
			mockMiddlewares: []*MockMiddleware{},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "POST",
					pattern:     "/",
					mockHandler: okHandler("POST/"),
					expectation: "POST/",
				},
			},
		},
		{
			name:            "should serve a mockHandler to different routes",
			mockMiddlewares: []*MockMiddleware{},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "GET",
					pattern:     "/users",
					mockHandler: okHandler("GET/users"),
					expectation: "GET/users",
				},
			},
		},
		{
			name:            "should serve multiple handlers to different routes",
			mockMiddlewares: []*MockMiddleware{},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "POST",
					pattern:     "/",
					mockHandler: okHandler("POST/"),
					expectation: "POST/",
				},
				{
					method:      "GET",
					pattern:     "/users",
					mockHandler: okHandler("GET/users"),
					expectation: "GET/users",
				},
				{
					method:      "POST",
					pattern:     "/users",
					mockHandler: okHandler("POST/users"),
					expectation: "POST/users",
				},
			},
		},
		{
			name:            "should serve a mockHandler with a middleware",
			mockMiddlewares: []*MockMiddleware{NewMockMiddleware()},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
			},
		},
		{
			name:            "should serve multiple handlers with a middleware",
			mockMiddlewares: []*MockMiddleware{NewMockMiddleware()},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "POST",
					pattern:     "/",
					mockHandler: okHandler("POST/"),
					expectation: "POST/",
				},
			},
		},
		{
			name:            "should serve a mockHandler to different routes with a middleware",
			mockMiddlewares: []*MockMiddleware{NewMockMiddleware()},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "GET",
					pattern:     "/users",
					mockHandler: okHandler("GET/users"),
					expectation: "GET/users",
				},
			},
		},
		{
			name:            "should serve multiple handlers to different routes with a middleware",
			mockMiddlewares: []*MockMiddleware{NewMockMiddleware()},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "POST",
					pattern:     "/",
					mockHandler: okHandler("POST/"),
					expectation: "POST/",
				},
				{
					method:      "GET",
					pattern:     "/users",
					mockHandler: okHandler("GET/users"),
					expectation: "GET/users",
				},
				{
					method:      "POST",
					pattern:     "/users",
					mockHandler: okHandler("POST/users"),
					expectation: "POST/users",
				},
			},
		},
		{
			name:            "should serve a mockHandler with multiple middlewares",
			mockMiddlewares: []*MockMiddleware{NewMockMiddleware(), NewMockMiddleware()},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
			},
		},
		{
			name:            "should serve multiple handlers with multiple middlewares",
			mockMiddlewares: []*MockMiddleware{NewMockMiddleware(), NewMockMiddleware()},
			args: []args{
				{
					method:      "GET",
					pattern:     "/",
					mockHandler: okHandler("GET/"),
					expectation: "GET/",
				},
				{
					method:      "POST",
					pattern:     "/",
					mockHandler: okHandler("POST/"),
					expectation: "POST/",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := NewRouter()

			for _, arg := range test.args {
				router.HandleMethod(arg.method, arg.pattern, arg.mockHandler.Handler())
			}

			for _, mockMiddleware := range test.mockMiddlewares {
				router.Use(mockMiddleware.Middleware())
			}

			for _, arg := range test.args {
				body := bytes.NewReader([]byte("some important body"))

				req := httptest.NewRequest(arg.method, arg.pattern, body)

				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				assert.Equalf(t, http.StatusOK, rr.Code, "code should be %d", http.StatusOK)

				assert.Equalf(t, arg.expectation, rr.Body.String(), "body should be %s", arg.expectation)

				arg.mockHandler.AssertNumberOfCalls(t, "ServeHTTP", 1)
			}

			for _, mockMiddleware := range test.mockMiddlewares {
				mockMiddleware.AssertNumberOfCalls(t, "Run", len(test.args))
			}
		})
	}
}
