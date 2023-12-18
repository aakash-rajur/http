package http

import (
	"bufio"
	"net"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter

	StatusCode int
}

func (rw *ResponseWriter) Hijacker() (http.Hijacker, bool) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)

	return hijacker, ok
}

func (rw *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return rw.ResponseWriter.(http.Hijacker).Hijack()
}

func (rw *ResponseWriter) Flusher() (http.Flusher, bool) {
	flusher, ok := rw.ResponseWriter.(http.Flusher)

	return flusher, ok
}

func (rw *ResponseWriter) Flush() {
	rw.ResponseWriter.(http.Flusher).Flush()
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode

	rw.ResponseWriter.WriteHeader(statusCode)
}
