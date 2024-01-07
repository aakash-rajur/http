package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRouter_Get(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Get("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_GetFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.GetFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Post(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Post("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_PostFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.PostFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Put(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Put("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_PutFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.PutFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Patch(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Patch("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_PatchFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.PatchFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Delete(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Delete("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_DeleteFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.DeleteFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Head(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Head("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_HeadFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.HeadFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Options(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Options("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_OptionsFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.OptionsFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Trace(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Trace("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_TraceFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.TraceFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_Connect(t *testing.T) {
	router := NewRouter()

	handler := NewMockHandler(nil)

	router.Connect("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}

func TestRouter_ConnectFunc(t *testing.T) {
	router := NewRouter()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	router.ConnectFunc("/test", handler)

	assert.Equalf(
		t,
		fmt.Sprintf("%v", handler),
		fmt.Sprintf("%v", router.register[0].Handler),
		"expected handler to be registered",
	)
}
