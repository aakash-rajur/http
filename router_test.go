package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aakash-rajur/http/params"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
				body := bytes.NewReader([]byte("some important Body"))

				req := httptest.NewRequest(arg.method, arg.pattern, body)

				rr := httptest.NewRecorder()

				router.ServeHTTP(rr, req)

				assert.Equalf(t, http.StatusOK, rr.Code, "code should be %d", http.StatusOK)

				assert.Equalf(t, arg.expectation, rr.Body.String(), "Body should be %s", arg.expectation)

				arg.mockHandler.AssertNumberOfCalls(t, "ServeHTTP", 1)
			}

			for _, mockMiddleware := range test.mockMiddlewares {
				mockMiddleware.AssertNumberOfCalls(t, "Run", len(test.args))
			}
		})
	}
}

func TestRouter_ServeHTTP_ParseApi(t *testing.T) {
	tests := []TestRoute{
		// Objects
		{
			Name:    "POST /1/classes/{className}",
			Method:  "POST",
			Pattern: "/1/classes/{className}",
		},
		{
			Name:    "GET /1/classes/{className}/{objectId}",
			Method:  "GET",
			Pattern: "/1/classes/{className}/{objectId}",
		},
		{
			Name:    "PUT /1/classes/{className}/{objectId}",
			Method:  "PUT",
			Pattern: "/1/classes/{className}/{objectId}",
		},
		{
			Name:    "GET /1/classes/{className}",
			Method:  "GET",
			Pattern: "/1/classes/{className}",
		},
		{
			Name:    "DELETE /1/classes/{className}/{objectId}",
			Method:  "DELETE",
			Pattern: "/1/classes/{className}/{objectId}",
		},

		// Users
		{
			Name:    "POST /1/users",
			Method:  "POST",
			Pattern: "/1/users",
		},
		{
			Name:    "GET /1/login",
			Method:  "GET",
			Pattern: "/1/login",
		},
		{
			Name:    "GET /1/users/{objectId}",
			Method:  "GET",
			Pattern: "/1/users/{objectId}",
		},
		{
			Name:    "PUT /1/users/{objectId}",
			Method:  "PUT",
			Pattern: "/1/users/{objectId}",
		},
		{
			Name:    "GET /1/users",
			Method:  "GET",
			Pattern: "/1/users",
		},
		{
			Name:    "DELETE /1/users/{objectId}",
			Method:  "DELETE",
			Pattern: "/1/users/{objectId}",
		},
		{
			Name:    "POST /1/requestPasswordReset",
			Method:  "POST",
			Pattern: "/1/requestPasswordReset",
		},

		// Roles
		{
			Name:    "POST /1/roles",
			Method:  "POST",
			Pattern: "/1/roles",
		},
		{
			Name:    "GET /1/roles/{objectId}",
			Method:  "GET",
			Pattern: "/1/roles/{objectId}",
		},
		{
			Name:    "PUT /1/roles/{objectId}",
			Method:  "PUT",
			Pattern: "/1/roles/{objectId}",
		},
		{
			Name:    "GET /1/roles",
			Method:  "GET",
			Pattern: "/1/roles",
		},
		{
			Name:    "DELETE /1/roles/{objectId}",
			Method:  "DELETE",
			Pattern: "/1/roles/{objectId}",
		},

		// Files
		{
			Name:    "POST /1/files/{fileName}",
			Method:  "POST",
			Pattern: "/1/files/{fileName}",
		},

		// Analytics
		{
			Name:    "POST /1/events/{eventName}",
			Method:  "POST",
			Pattern: "/1/events/{eventName}",
		},

		// Push Notifications
		{
			Name:    "POST /1/push",
			Method:  "POST",
			Pattern: "/1/push",
		},

		// Installations
		{
			Name:    "POST /1/installations",
			Method:  "POST",
			Pattern: "/1/installations",
		},
		{
			Name:    "GET /1/installations/{objectId}",
			Method:  "GET",
			Pattern: "/1/installations/{objectId}",
		},
		{
			Name:    "PUT /1/installations/{objectId}",
			Method:  "PUT",
			Pattern: "/1/installations/{objectId}",
		},
		{
			Name:    "GET /1/installations",
			Method:  "GET",
			Pattern: "/1/installations",
		},
		{
			Name:    "DELETE /1/installations/{objectId}",
			Method:  "DELETE",
			Pattern: "/1/installations/{objectId}",
		},

		// Cloud Functions
		{
			Name:    "POST /1/functions",
			Method:  "POST",
			Pattern: "/1/functions",
		},
	}

	for _, tc := range tests {
		tc := tc

		router := setupRouter(tests)

		seed := fmt.Sprintf("parse:%d", time.Now().UnixNano())

		t.Run(tc.Name, validate(tc, seed, router))
	}
}

func setupRouter(testRoutes []TestRoute) http.Handler {
	router := NewRouter()

	m1 := func(w http.ResponseWriter, r *http.Request, next Next) {
		ctx := context.WithValue(r.Context(), "m1", "m1")

		rc := r.WithContext(ctx)

		next(rc)
	}

	m2 := func(w http.ResponseWriter, r *http.Request, next Next) {
		ctx := context.WithValue(r.Context(), "m2", "m2")

		rc := r.WithContext(ctx)

		next(rc)
	}

	router.Use(m1)

	router.Use(m2)

	handler := func(w http.ResponseWriter, r *http.Request) {
		pathParams, ok := params.FromRequest(r)

		if !ok {
			pathParams = make(map[string]string)
		}

		body := make(map[string]any)

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		m1 := r.Context().Value("m1")

		m2 := r.Context().Value("m2")

		payload := map[string]interface{}{
			"Params": pathParams,
			"Path":   r.URL.Path,
			"query":  r.URL.Query(),
			"Method": r.Method,
			"body":   body,
			"m1":     m1,
			"m2":     m2,
		}

		jsonPayload, err := json.Marshal(payload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		_, _ = w.Write(jsonPayload)
	}

	for _, tc := range testRoutes {
		router.HandleMethodFunc(tc.Method, tc.Pattern, handler)
	}

	return router
}

func validate(tc TestRoute, seed string, router http.Handler) func(t *testing.T) {
	return func(t *testing.T) {
		tcp := tc.generateParams(seed)

		want := map[string]any{
			"m1":     "m1",
			"m2":     "m2",
			"Method": tcp.Method,
			"Params": tcp.Params,
			"Path":   tcp.Path,
			"query":  map[string]any{},
			"body":   tcp.Body.Map(),
		}

		r := httptest.NewRequest(tcp.Method, tcp.Path, tcp.Body.Reader())

		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		assert.Equalf(t, http.StatusOK, w.Code, "%s:%s code should be %d", tcp.Method, tcp.Pattern, http.StatusOK)

		if w.Code != http.StatusOK {
			return
		}

		got := make(map[string]any)

		err := json.Unmarshal(w.Body.Bytes(), &got)

		assert.NoErrorf(t, err, "should not error while unmarshalling Body")

		assert.Equalf(t, want, got, "%s:%s Body should be %v", tcp.Method, tcp.Pattern, want)
	}
}
