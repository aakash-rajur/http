package webtransport

import (
	h "github.com/aakash-rajur/http"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_withWebTransport(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	wt := &webtransport.Server{}

	got := withWebTransport(r, wt)

	cwt := got.Context().Value(wtKey)

	assert.Equalf(t, wt, cwt, "want wt == cwt, got %v", cwt)
}

func TestMiddleware(t *testing.T) {
	wt := &webtransport.Server{}

	m := Middleware(wt)

	assert.NotNilf(t, m, "want m != nil")

	middlewares := h.Middlewares{m}

	r := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	handler := &mockHandler{}

	handler.On(
		"ServeHTTP",
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("*http.Request"),
	)

	assert.Nilf(t, handler.wts, "want handler.wts == nil, got %v", handler.wts)

	middlewares.Chain(handler.ServeHTTP)(w, r)

	handler.AssertNumberOfCalls(t, "ServeHTTP", 1)

	assert.Equalf(t, wt, handler.wts, "want wt == handler.wts, got %v", handler.wts)
}

func TestUpgrade(t *testing.T) {
	wt := &webtransport.Server{
		H3: http3.Server{},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	r := httptest.NewRequest(http.MethodConnect, "/", &mockH3Body{})

	r.Proto = "webtransport"
	r.ProtoMajor = 3
	r.ProtoMinor = 0
	r.Header.Set("Sec-Webtransport-Http3-Draft02", "1")

	r = withWebTransport(r, wt)

	w := httptest.NewRecorder()

	s, err := Upgrade(w, r)

	// asserting non-nil error and nil session because the mockH3Body is not a valid body
	// we're only testing the call to Upgrade and the error returned is from underlying mechanism
	assert.Errorf(t, err, "want err != nil, got %v", err)

	assert.Nilf(t, s, "want s == nil, got %v", s)
}
