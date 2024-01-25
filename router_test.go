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

func TestRouter_ServeHTTP_GithubApi(t *testing.T) {
	testCases := []TestRoute{
		// OAuth Authorizations
		{
			Name:    "GET /authorizations",
			Method:  "GET",
			Pattern: "/authorizations",
		},
		{
			Name:    "GET /authorizations/{id}",
			Method:  "GET",
			Pattern: "/authorizations/{id}",
		},
		{
			Name:    "POST /authorizations",
			Method:  "POST",
			Pattern: "/authorizations",
		},
		{
			Name:    "PUT /authorizations/clients/{client_id}",
			Method:  "PUT",
			Pattern: "/authorizations/clients/{client_id}",
		},
		{
			Name:    "PATCH /authorizations/{id}",
			Method:  "PATCH",
			Pattern: "/authorizations/{id}",
		},
		{
			Name:    "DELETE /authorizations/{id}",
			Method:  "DELETE",
			Pattern: "/authorizations/{id}",
		},
		{
			Name:    "GET /applications/{client_id}/tokens/{access_token}",
			Method:  "GET",
			Pattern: "/applications/{client_id}/tokens/{access_token}",
		},
		{
			Name:    "DELETE /applications/{client_id}/tokens",
			Method:  "DELETE",
			Pattern: "/applications/{client_id}/tokens",
		},
		{
			Name:    "DELETE /applications/{client_id}/tokens/{access_token}",
			Method:  "DELETE",
			Pattern: "/applications/{client_id}/tokens/{access_token}",
		},

		// Activity
		{
			Name:    "GET /events",
			Method:  "GET",
			Pattern: "/events",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/events",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/events",
		},
		{
			Name:    "GET /networks/{owner}/{repo}/events",
			Method:  "GET",
			Pattern: "/networks/{owner}/{repo}/events",
		},
		{
			Name:    "GET /orgs/{org}/events",
			Method:  "GET",
			Pattern: "/orgs/{org}/events",
		},
		{
			Name:    "GET /users/{user}/received_events",
			Method:  "GET",
			Pattern: "/users/{user}/received_events",
		},
		{
			Name:    "GET /users/{user}/received_events/public",
			Method:  "GET",
			Pattern: "/users/{user}/received_events/public",
		},
		{
			Name:    "GET /users/{user}/events",
			Method:  "GET",
			Pattern: "/users/{user}/events",
		},
		{
			Name:    "GET /users/{user}/events/public",
			Method:  "GET",
			Pattern: "/users/{user}/events/public",
		},
		{
			Name:    "GET /users/{user}/events/orgs/{org}",
			Method:  "GET",
			Pattern: "/users/{user}/events/orgs/{org}",
		},
		{
			Name:    "GET /feeds",
			Method:  "GET",
			Pattern: "/feeds",
		},
		{
			Name:    "GET /notifications",
			Method:  "GET",
			Pattern: "/notifications",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/notifications",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/notifications",
		},
		{
			Name:    "PUT /notifications",
			Method:  "PUT",
			Pattern: "/notifications",
		},
		{
			Name:    "PUT /repos/{owner}/{repo}/notifications",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/notifications",
		},
		{
			Name:    "GET /notifications/threads/{id}",
			Method:  "GET",
			Pattern: "/notifications/threads/{id}",
		},
		{
			Name:    "PATCH /notifications/threads/{id}",
			Method:  "PATCH",
			Pattern: "/notifications/threads/{id}",
		},
		{
			Name:    "GET /notifications/threads/{id}/subscription",
			Method:  "GET",
			Pattern: "/notifications/threads/{id}/subscription",
		},
		{
			Name:    "PUT /notifications/threads/{id}/subscription",
			Method:  "PUT",
			Pattern: "/notifications/threads/{id}/subscription",
		},
		{
			Name:    "DELETE /notifications/threads/{id}/subscription",
			Method:  "DELETE",
			Pattern: "/notifications/threads/{id}/subscription",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/stargazers",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/stargazers",
		},
		{
			Name:    "GET /users/{user}/starred",
			Method:  "GET",
			Pattern: "/users/{user}/starred",
		},
		{
			Name:    "GET /user/starred",
			Method:  "GET",
			Pattern: "/user/starred",
		},
		{
			Name:    "GET /user/starred/{owner}/{repo}",
			Method:  "GET",
			Pattern: "/user/starred/{owner}/{repo}",
		},
		{
			Name:    "PUT /user/starred/{owner}/{repo}",
			Method:  "PUT",
			Pattern: "/user/starred/{owner}/{repo}",
		},
		{
			Name:    "DELETE /user/starred/{owner}/{repo}",
			Method:  "DELETE",
			Pattern: "/user/starred/{owner}/{repo}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/subscribers",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/subscribers",
		},
		{
			Name:    "GET /users/{user}/subscriptions",
			Method:  "GET",
			Pattern: "/users/{user}/subscriptions",
		},
		{
			Name:    "GET /user/subscriptions",
			Method:  "GET",
			Pattern: "/user/subscriptions",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/subscription",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/subscription",
		},
		{
			Name:    "PUT /repos/{owner}/{repo}/subscription",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/subscription",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/subscription",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/subscription",
		},
		{
			Name:    "GET /user/subscriptions/{owner}/{repo}",
			Method:  "GET",
			Pattern: "/user/subscriptions/{owner}/{repo}",
		},
		{
			Name:    "PUT /user/subscriptions/{owner}/{repo}",
			Method:  "PUT",
			Pattern: "/user/subscriptions/{owner}/{repo}",
		},
		{
			Name:    "DELETE /user/subscriptions/{owner}/{repo}",
			Method:  "DELETE",
			Pattern: "/user/subscriptions/{owner}/{repo}",
		},

		// Gists
		{
			Name:    "GET /users/{user}/gists",
			Method:  "GET",
			Pattern: "/users/{user}/gists",
		},
		{
			Name:    "GET /gists",
			Method:  "GET",
			Pattern: "/gists",
		},
		/*
			// matches with 'GET /gists/{id}'
			{
				Name:    "GET /gists/public",
				Method:  "GET",
				Pattern: "/gists/public",
			},
			// matches with 'GET /gists/{id}'
			{
				Name:    "GET /gists/starred",
				Method:  "GET",
				Pattern: "/gists/starred",
			},
		*/
		{
			Name:    "GET /gists/{id}",
			Method:  "GET",
			Pattern: "/gists/{id}",
		},
		{
			Name:    "POST /gists",
			Method:  "POST",
			Pattern: "/gists",
		},
		{
			Name:    "PATCH /gists/{id}",
			Method:  "PATCH",
			Pattern: "/gists/{id}",
		},
		{
			Name:    "PUT /gists/{id}/star",
			Method:  "PUT",
			Pattern: "/gists/{id}/star",
		},
		{
			Name:    "DELETE /gists/{id}/star",
			Method:  "DELETE",
			Pattern: "/gists/{id}/star",
		},
		{
			Name:    "GET /gists/{id}/star",
			Method:  "GET",
			Pattern: "/gists/{id}/star",
		},
		{
			Name:    "POST /gists/{id}/forks",
			Method:  "POST",
			Pattern: "/gists/{id}/forks",
		},
		{
			Name:    "DELETE /gists/{id}",
			Method:  "DELETE",
			Pattern: "/gists/{id}",
		},

		// Git Data
		{
			Name:    "GET /repos/{owner}/{repo}/git/blobs/{sha}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/git/blobs/{sha}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/git/blobs",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/git/blobs",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/git/commits/{sha}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/git/commits/{sha}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/git/commits",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/git/commits",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/git/refs/*ref",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/git/refs/*ref",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/git/refs",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/git/refs",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/git/refs",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/git/refs",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/git/refs/*ref",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/git/refs/*ref",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/git/refs/*ref",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/git/refs/*ref",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/git/tags/{sha}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/git/tags/{sha}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/git/tags",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/git/tags",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/git/trees/{sha}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/git/trees/{sha}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/git/trees",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/git/trees",
		},

		// Issues
		{
			Name:    "GET /issues",
			Method:  "GET",
			Pattern: "/issues",
		},
		{
			Name:    "GET /user/issues",
			Method:  "GET",
			Pattern: "/user/issues",
		},
		{
			Name:    "GET /orgs/{org}/issues",
			Method:  "GET",
			Pattern: "/orgs/{org}/issues",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/issues",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/issues",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/issues/{number}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/issues/{number}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/issues",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/issues",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/issues/{number}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/issues/{number}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/assignees",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/assignees",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/assignees/{assignee}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/assignees/{assignee}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/issues/{number}/comments",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/comments",
		},
		/*
			// matches with 'GET /repos/{owner}/{repo}/issues/{number}'
			{
				Name:    "GET /repos/{owner}/{repo}/issues/comments",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/issues/comments",
			},
			// matches with 'GET /repos/{owner}/{repo}/issues/{number}/comments'
			{
				Name:    "GET /repos/{owner}/{repo}/issues/comments/{id}",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/issues/comments/{id}",
			},
		*/
		{
			Name:    "POST /repos/{owner}/{repo}/issues/{number}/comments",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/comments",
		},
		/*
				// matches with 'PATCH /repos/{owner}/{repo}/issues/{number}/'
			{
				Name:    "PATCH /repos/{owner}/{repo}/issues/comments/{id}",
				Method:  "PATCH",
				Pattern: "/repos/{owner}/{repo}/issues/comments/{id}",
			},
		*/
		{
			Name:    "DELETE /repos/{owner}/{repo}/issues/comments/{id}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/issues/comments/{id}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/issues/{number}/events",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/events",
		},
		/*
			// matches with 'GET /repos/{owner}/{repo}/issues/{number}'
			{
				Name:    "GET /repos/{owner}/{repo}/issues/events",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/issues/events",
			},
			// matches with 'GET /repos/{owner}/{repo}/issues/{number}/events'
			{
				Name:    "GET /repos/{owner}/{repo}/issues/events/{id}",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/issues/events/{id}",
			},
		*/
		{
			Name:    "GET /repos/{owner}/{repo}/labels",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/labels",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/labels/{Name}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/labels/{Name}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/labels",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/labels",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/labels/{Name}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/labels/{Name}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/labels/{Name}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/labels/{Name}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/issues/{number}/labels",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/labels",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/issues/{number}/labels",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/labels",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/issues/{number}/labels/{Name}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/labels/{Name}",
		},
		{
			Name:    "PUT /repos/{owner}/{repo}/issues/{number}/labels",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/labels",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/issues/{number}/labels",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/issues/{number}/labels",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/milestones/{number}/labels",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/milestones/{number}/labels",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/milestones",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/milestones",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/milestones/{number}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/milestones/{number}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/milestones",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/milestones",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/milestones/{number}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/milestones/{number}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/milestones/{number}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/milestones/{number}",
		},

		// Miscellaneous
		{
			Name:    "GET /emojis",
			Method:  "GET",
			Pattern: "/emojis",
		},
		{
			Name:    "GET /gitignore/templates",
			Method:  "GET",
			Pattern: "/gitignore/templates",
		},
		{
			Name:    "GET /gitignore/templates/{Name}",
			Method:  "GET",
			Pattern: "/gitignore/templates/{Name}",
		},
		{
			Name:    "POST /markdown",
			Method:  "POST",
			Pattern: "/markdown",
		},
		{
			Name:    "POST /markdown/raw",
			Method:  "POST",
			Pattern: "/markdown/raw",
		},
		{
			Name:    "GET /meta",
			Method:  "GET",
			Pattern: "/meta",
		},
		{
			Name:    "GET /rate_limit",
			Method:  "GET",
			Pattern: "/rate_limit",
		},

		// Organizations
		{
			Name:    "GET /users/{user}/orgs",
			Method:  "GET",
			Pattern: "/users/{user}/orgs",
		},
		{
			Name:    "GET /user/orgs",
			Method:  "GET",
			Pattern: "/user/orgs",
		},
		{
			Name:    "GET /orgs/{org}",
			Method:  "GET",
			Pattern: "/orgs/{org}",
		},
		{
			Name:    "PATCH /orgs/{org}",
			Method:  "PATCH",
			Pattern: "/orgs/{org}",
		},
		{
			Name:    "GET /orgs/{org}/members",
			Method:  "GET",
			Pattern: "/orgs/{org}/members",
		},
		{
			Name:    "GET /orgs/{org}/members/{user}",
			Method:  "GET",
			Pattern: "/orgs/{org}/members/{user}",
		},
		{
			Name:    "DELETE /orgs/{org}/members/{user}",
			Method:  "DELETE",
			Pattern: "/orgs/{org}/members/{user}",
		},
		{
			Name:    "GET /orgs/{org}/public_members",
			Method:  "GET",
			Pattern: "/orgs/{org}/public_members",
		},
		{
			Name:    "GET /orgs/{org}/public_members/{user}",
			Method:  "GET",
			Pattern: "/orgs/{org}/public_members/{user}",
		},
		{
			Name:    "PUT /orgs/{org}/public_members/{user}",
			Method:  "PUT",
			Pattern: "/orgs/{org}/public_members/{user}",
		},
		{
			Name:    "DELETE /orgs/{org}/public_members/{user}",
			Method:  "DELETE",
			Pattern: "/orgs/{org}/public_members/{user}",
		},
		{
			Name:    "GET /orgs/{org}/teams",
			Method:  "GET",
			Pattern: "/orgs/{org}/teams",
		},
		{
			Name:    "GET /teams/{id}",
			Method:  "GET",
			Pattern: "/teams/{id}",
		},
		{
			Name:    "POST /orgs/{org}/teams",
			Method:  "POST",
			Pattern: "/orgs/{org}/teams",
		},
		{
			Name:    "PATCH /teams/{id}",
			Method:  "PATCH",
			Pattern: "/teams/{id}",
		},
		{
			Name:    "DELETE /teams/{id}",
			Method:  "DELETE",
			Pattern: "/teams/{id}",
		},
		{
			Name:    "GET /teams/{id}/members",
			Method:  "GET",
			Pattern: "/teams/{id}/members",
		},
		{
			Name:    "GET /teams/{id}/members/{user}",
			Method:  "GET",
			Pattern: "/teams/{id}/members/{user}",
		},
		{
			Name:    "PUT /teams/{id}/members/{user}",
			Method:  "PUT",
			Pattern: "/teams/{id}/members/{user}",
		},
		{
			Name:    "DELETE /teams/{id}/members/{user}",
			Method:  "DELETE",
			Pattern: "/teams/{id}/members/{user}",
		},
		{
			Name:    "GET /teams/{id}/repos",
			Method:  "GET",
			Pattern: "/teams/{id}/repos",
		},
		{
			Name:    "GET /teams/{id}/repos/{owner}/{repo}",
			Method:  "GET",
			Pattern: "/teams/{id}/repos/{owner}/{repo}",
		},
		{
			Name:    "PUT /teams/{id}/repos/{owner}/{repo}",
			Method:  "PUT",
			Pattern: "/teams/{id}/repos/{owner}/{repo}",
		},
		{
			Name:    "DELETE /teams/{id}/repos/{owner}/{repo}",
			Method:  "DELETE",
			Pattern: "/teams/{id}/repos/{owner}/{repo}",
		},
		{
			Name:    "GET /user/teams",
			Method:  "GET",
			Pattern: "/user/teams",
		},

		// Pull Requests
		{
			Name:    "GET /repos/{owner}/{repo}/pulls",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/pulls",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/pulls/{number}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/pulls",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/pulls",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/pulls/{number}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/pulls/{number}/commits",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}/commits",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/pulls/{number}/files",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}/files",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/pulls/{number}/merge",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}/merge",
		},
		{
			Name:    "PUT /repos/{owner}/{repo}/pulls/{number}/merge",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}/merge",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/pulls/{number}/comments",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}/comments",
		},
		/*
			// matches with 'GET /repos/{owner}/{repo}/pulls/{number}'
			{
				Name:    "GET /repos/{owner}/{repo}/pulls/comments",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/pulls/comments",
			},
			// matches with 'GET /repos/{owner}/{repo}/pulls/{number}/comments'
			{
				Name:    "GET /repos/{owner}/{repo}/pulls/comments/{number}",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/pulls/comments/{number}",
			},
		*/
		{
			Name:    "PUT /repos/{owner}/{repo}/pulls/{number}/comments",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/pulls/{number}/comments",
		},
		/*
			{
				Name:    "PATCH /repos/{owner}/{repo}/pulls/comments/{number}",
				Method:  "PATCH",
				Pattern: "/repos/{owner}/{repo}/pulls/comments/{number}",
			},
		*/
		{
			Name:    "DELETE /repos/{owner}/{repo}/pulls/comments/{number}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/pulls/comments/{number}",
		},

		// Repositories
		{
			Name:    "GET /user/repos",
			Method:  "GET",
			Pattern: "/user/repos",
		},
		{
			Name:    "GET /users/{user}/repos",
			Method:  "GET",
			Pattern: "/users/{user}/repos",
		},
		{
			Name:    "GET /orgs/{org}/repos",
			Method:  "GET",
			Pattern: "/orgs/{org}/repos",
		},
		{
			Name:    "GET /repositories",
			Method:  "GET",
			Pattern: "/repositories",
		},
		{
			Name:    "POST /user/repos",
			Method:  "POST",
			Pattern: "/user/repos",
		},
		{
			Name:    "POST /orgs/{org}/repos",
			Method:  "POST",
			Pattern: "/orgs/{org}/repos",
		},
		{
			Name:    "GET /repos/{owner}/{repo}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/contributors",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/contributors",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/languages",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/languages",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/teams",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/teams",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/tags",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/tags",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/branches",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/branches",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/branches/{branch}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/branches/{branch}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/collaborators",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/collaborators",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/collaborators/{user}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/collaborators/{user}",
		},
		{
			Name:    "PUT /repos/{owner}/{repo}/collaborators/{user}",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/collaborators/{user}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/collaborators/{user}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/collaborators/{user}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/comments",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/comments",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/commits/{sha}/comments",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/commits/{sha}/comments",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/commits/{sha}/comments",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/commits/{sha}/comments",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/comments/{id}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/comments/{id}",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/comments/{id}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/comments/{id}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/comments/{id}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/comments/{id}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/commits",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/commits",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/commits/{sha}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/commits/{sha}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/readme",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/readme",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/contents/*Path",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/contents/*Path",
		},
		{
			Name:    "PUT /repos/{owner}/{repo}/contents/*Path",
			Method:  "PUT",
			Pattern: "/repos/{owner}/{repo}/contents/*Path",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/contents/*Path",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/contents/*Path",
		},
		/*
			{
				Name:    "GET /repos/{owner}/{repo}/{archive_format}/{ref}",
				Method:  "GET",
				Pattern: "/repos/{owner}/{repo}/{archive_format}/{ref}",
			},
		*/
		{
			Name:    "GET /repos/{owner}/{repo}/keys",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/keys",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/keys/{id}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/keys/{id}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/keys",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/keys",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/keys/{id}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/keys/{id}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/keys/{id}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/keys/{id}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/downloads",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/downloads",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/downloads/{id}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/downloads/{id}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/downloads/{id}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/downloads/{id}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/forks",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/forks",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/forks",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/forks",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/hooks",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/hooks",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/hooks/{id}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/hooks/{id}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/hooks",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/hooks",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/hooks/{id}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/hooks/{id}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/hooks/{id}/tests",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/hooks/{id}/tests",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/hooks/{id}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/hooks/{id}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/merges",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/merges",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/releases",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/releases",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/releases/{id}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/releases/{id}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/releases",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/releases",
		},
		{
			Name:    "PATCH /repos/{owner}/{repo}/releases/{id}",
			Method:  "PATCH",
			Pattern: "/repos/{owner}/{repo}/releases/{id}",
		},
		{
			Name:    "DELETE /repos/{owner}/{repo}/releases/{id}",
			Method:  "DELETE",
			Pattern: "/repos/{owner}/{repo}/releases/{id}",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/releases/{id}/assets",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/releases/{id}/assets",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/stats/contributors",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/stats/contributors",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/stats/commit_activity",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/stats/commit_activity",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/stats/code_frequency",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/stats/code_frequency",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/stats/participation",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/stats/participation",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/stats/punch_card",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/stats/punch_card",
		},
		{
			Name:    "GET /repos/{owner}/{repo}/statuses/{ref}",
			Method:  "GET",
			Pattern: "/repos/{owner}/{repo}/statuses/{ref}",
		},
		{
			Name:    "POST /repos/{owner}/{repo}/statuses/{ref}",
			Method:  "POST",
			Pattern: "/repos/{owner}/{repo}/statuses/{ref}",
		},

		// Search
		{
			Name:    "GET /search/repositories",
			Method:  "GET",
			Pattern: "/search/repositories",
		},
		{
			Name:    "GET /search/code",
			Method:  "GET",
			Pattern: "/search/code",
		},
		{
			Name:    "GET /search/issues",
			Method:  "GET",
			Pattern: "/search/issues",
		},
		{
			Name:    "GET /search/users",
			Method:  "GET",
			Pattern: "/search/users",
		},
		{
			Name:    "GET /legacy/issues/search/{owner}/{repository}/{state}/{keyword}",
			Method:  "GET",
			Pattern: "/legacy/issues/search/{owner}/{repository}/{state}/{keyword}",
		},
		{
			Name:    "GET /legacy/repos/search/{keyword}",
			Method:  "GET",
			Pattern: "/legacy/repos/search/{keyword}",
		},
		{
			Name:    "GET /legacy/user/search/{keyword}",
			Method:  "GET",
			Pattern: "/legacy/user/search/{keyword}",
		},
		{
			Name:    "GET /legacy/user/email/{email}",
			Method:  "GET",
			Pattern: "/legacy/user/email/{email}",
		},

		// Users
		{
			Name:    "GET /users/{user}",
			Method:  "GET",
			Pattern: "/users/{user}",
		},
		{
			Name:    "GET /user",
			Method:  "GET",
			Pattern: "/user",
		},
		{
			Name:    "PATCH /user",
			Method:  "PATCH",
			Pattern: "/user",
		},
		{
			Name:    "GET /users",
			Method:  "GET",
			Pattern: "/users",
		},
		{
			Name:    "GET /user/emails",
			Method:  "GET",
			Pattern: "/user/emails",
		},
		{
			Name:    "POST /user/emails",
			Method:  "POST",
			Pattern: "/user/emails",
		},
		{
			Name:    "DELETE /user/emails",
			Method:  "DELETE",
			Pattern: "/user/emails",
		},
		{
			Name:    "GET /users/{user}/followers",
			Method:  "GET",
			Pattern: "/users/{user}/followers",
		},
		{
			Name:    "GET /user/followers",
			Method:  "GET",
			Pattern: "/user/followers",
		},
		{
			Name:    "GET /users/{user}/following",
			Method:  "GET",
			Pattern: "/users/{user}/following",
		},
		{
			Name:    "GET /user/following",
			Method:  "GET",
			Pattern: "/user/following",
		},
		{
			Name:    "GET /user/following/{user}",
			Method:  "GET",
			Pattern: "/user/following/{user}",
		},
		{
			Name:    "GET /users/{user}/following/{target_user}",
			Method:  "GET",
			Pattern: "/users/{user}/following/{target_user}",
		},
		{
			Name:    "PUT /user/following/{user}",
			Method:  "PUT",
			Pattern: "/user/following/{user}",
		},
		{
			Name:    "DELETE /user/following/{user}",
			Method:  "DELETE",
			Pattern: "/user/following/{user}",
		},
		{
			Name:    "GET /users/{user}/keys",
			Method:  "GET",
			Pattern: "/users/{user}/keys",
		},
		{
			Name:    "GET /user/keys",
			Method:  "GET",
			Pattern: "/user/keys",
		},
		{
			Name:    "GET /user/keys/{id}",
			Method:  "GET",
			Pattern: "/user/keys/{id}",
		},
		{
			Name:    "POST /user/keys",
			Method:  "POST",
			Pattern: "/user/keys",
		},
		{
			Name:    "PATCH /user/keys/{id}",
			Method:  "PATCH",
			Pattern: "/user/keys/{id}",
		},
		{
			Name:    "DELETE /user/keys/{id}",
			Method:  "DELETE",
			Pattern: "/user/keys/{id}",
		},
	}

	for _, tc := range testCases {
		tc := tc

		router := setupRouter(testCases)

		seed := fmt.Sprintf("github:%d", time.Now().UnixNano())

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
