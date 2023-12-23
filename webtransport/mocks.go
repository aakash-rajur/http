package webtransport

import (
	"context"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/stretchr/testify/mock"
	"net/http"
	"time"
)

type mockH3Body struct{}

func (m *mockH3Body) Read(p []byte) (n int, err error) {
	panic("mock")
}

func (m *mockH3Body) Close() error {
	panic("mock")
}

func (m *mockH3Body) HTTPStream() http3.Stream {
	stream := new(mockStream)

	return stream
}

type mockStream struct{}

func (m *mockStream) StreamID() quic.StreamID {
	return 1
}

func (m *mockStream) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m *mockStream) CancelRead(code quic.StreamErrorCode) {
	panic("mock")
}

func (m *mockStream) SetReadDeadline(t time.Time) error {
	panic("mock")
}

func (m *mockStream) Write(p []byte) (n int, err error) {
	panic("mock")
}

func (m *mockStream) Close() error {
	panic("mock")
}

func (m *mockStream) CancelWrite(code quic.StreamErrorCode) {
	panic("mock")
}

func (m *mockStream) Context() context.Context {
	panic("mock")
}

func (m *mockStream) SetWriteDeadline(t time.Time) error {
	panic("mock")
}

func (m *mockStream) SetDeadline(t time.Time) error {
	panic("mock")
}

type mockHandler struct {
	mock.Mock
	wts any
}

func (mh *mockHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	mh.Called(rw, r)

	mh.wts = r.Context().Value(wtKey)

	rw.WriteHeader(http.StatusOK)
}
