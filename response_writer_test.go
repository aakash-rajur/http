package http

import (
	"bufio"
	"fmt"
	"net"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResponseWriter_WriteHeader(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		statusCode int
	}{
		{statusCode: 200},
		{statusCode: 404},
		{statusCode: 500},
		{statusCode: 503},
		{statusCode: 301},
		{statusCode: 302},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("status code %d", tc.statusCode)

		t.Run(testName, func(t *testing.T) {
			rr := httptest.NewRecorder()

			rw := &ResponseWriter{
				ResponseWriter: rr,
			}

			rw.WriteHeader(tc.statusCode)

			assert.Equal(t, rw.StatusCode, tc.statusCode, "want same status code")
		})
	}
}

func TestResponseWriter_Write(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		body string
	}{
		{body: "Hello World 1"},
		{body: "Hello World 2"},
		{body: "Hello World 3"},
		{body: "Hello World 4"},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("body %s", tc.body)

		t.Run(testName, func(t *testing.T) {
			rr := httptest.NewRecorder()

			rw := &ResponseWriter{
				ResponseWriter: rr,
			}

			rw.Write([]byte(tc.body))

			assert.Equal(t, rr.Body.String(), tc.body, "want same body")
		})
	}
}

func TestResponseWriter_Hijacker(t *testing.T) {
	t.Parallel()

	t.Run("no hijacker", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rw := &ResponseWriter{
			ResponseWriter: rr,
		}

		hijacker, ok := rw.Hijacker()

		assert.False(t, ok, "want hijacker ok false")

		assert.Nil(t, hijacker, "want hijacker nil")
	})

	t.Run("hijacker", func(t *testing.T) {
		rr := &hijackerRw{ResponseRecorder: *httptest.NewRecorder()}

		rw := &ResponseWriter{
			ResponseWriter: rr,
		}

		hijacker, ok := rw.Hijacker()

		assert.True(t, ok, "want hijacker ok true")

		assert.NotNil(t, hijacker, "want hijacker not nil")
	})
}

func TestResponseWriter_Hijack(t *testing.T) {
	t.Parallel()

	rr := &hijackerRw{ResponseRecorder: *httptest.NewRecorder()}

	rw := &ResponseWriter{ResponseWriter: rr}

	conn, brw, err := rw.Hijack()

	assert.Nil(t, err, "want err nil")

	assert.NotNil(t, conn, "want conn not nil")

	assert.NotNil(t, brw, "want rw not nil")
}

func TestResponseWriter_Flusher(t *testing.T) {
	rw := &ResponseWriter{ResponseWriter: httptest.NewRecorder()}

	flusher, ok := rw.Flusher()

	assert.True(t, ok, "want flusher ok true")

	assert.NotNil(t, flusher, "want flusher not nil")
}

func TestResponseWriter_Flush(t *testing.T) {
	rr := httptest.NewRecorder()

	rw := &ResponseWriter{ResponseWriter: rr}

	rw.Flush()

	assert.True(t, rr.Flushed, "want flushed true")
}

type hijackerRw struct {
	httptest.ResponseRecorder
}

func (rw *hijackerRw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	conn := &hijackerConn{}

	brw := bufio.NewReadWriter(nil, nil)

	return conn, brw, nil
}

type hijackerConn struct {
	net.Conn
}

func (c *hijackerConn) Close() error {
	return nil
}

func (c *hijackerConn) LocalAddr() net.Addr {
	return nil
}

func (c *hijackerConn) RemoteAddr() net.Addr {
	return nil
}

func (c *hijackerConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *hijackerConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *hijackerConn) SetWriteDeadline(t time.Time) error {
	return nil
}
