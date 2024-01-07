package http

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args LoggerConfig
		want []string
	}{
		{
			name: "default",
			args: LoggerConfig{
				Output: NewMemoryWriter(),
			},
			want: []string{
				time.Now().Format(time.DateOnly),
				"HTTP/1",
				"200",
				"192.",
				"GET",
				"/",
				"application/json",
				"text/plain",
			},
		},
		{
			name: "custom log formatter",
			args: LoggerConfig{
				Output: NewMemoryWriter(),
				LogFormatter: func(l LogFormatterParams) string {
					return "test"
				},
			},
			want: []string{
				"test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := Logger(tt.args)

			handler := func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}

			middlewares := Middlewares{middleware}

			handler = middlewares.Chain(handler)

			mrw := &ResponseWriter{ResponseWriter: httptest.NewRecorder()}

			mr := httptest.NewRequest(http.MethodGet, "/", nil)

			mr.Header.Set("Content-Type", "application/json")

			handler(mrw, mr)

			mw := tt.args.Output.(*MemoryWriter)

			for _, each := range tt.want {
				assert.Containsf(t, mw.Content, each, "Logger() = %v, want %v", mw.Content, each)
			}
		})
	}
}

func Test_defaultLogFormatter(t *testing.T) {

	tests := []struct {
		name string
		args LogFormatterParams
		want []string
	}{
		{
			name: "default",
			args: LogFormatterParams{
				Timestamp:           time.Now(),
				ProtocolVersion:     1,
				StatusCode:          200,
				ClientIP:            "192.168.0.10",
				Method:              http.MethodGet,
				Path:                "/",
				RequestContentType:  "application/json",
				ResponseContentType: "text/plain",
				Latency:             time.Duration(0),
			},
			want: []string{
				time.Now().Format(time.DateOnly),
				"HTTP/1",
				"200",
				"192.168.0.10",
				"GET",
				"/",
				"application/json",
				"text/plain",
			},
		},
		{
			name: "custom",
			args: LogFormatterParams{
				Timestamp:           time.Now(),
				ProtocolVersion:     1,
				StatusCode:          200,
				ClientIP:            "192.168.0.10",
				Method:              http.MethodPost,
				Path:                "/api/v1/books",
				Query:               map[string][]string{"test": {"test"}},
				RequestContentType:  "application/json",
				ResponseContentType: "text/plain",
				Latency:             time.Duration(0),
			},
			want: []string{
				time.Now().Format(time.DateOnly),
				"HTTP/1",
				"200",
				"192.168.0.10",
				"POST",
				"/api/v1/books",
				"application/json",
				"text/plain",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultLogFormatter(tt.args)

			for _, each := range tt.want {
				assert.Containsf(t, got, each, "defaultLogFormatter() = %v, want %v", got, each)
			}
		})
	}
}

func Test_saneConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args LoggerConfig
		want LoggerConfig
	}{
		{
			name: "default",
			args: LoggerConfig{},
			want: LoggerConfig{
				Output:       os.Stdout,
				LogFormatter: defaultLogFormatter,
			},
		},
		{
			name: "custom output",
			args: LoggerConfig{
				Output: NewMemoryWriter(),
			},
			want: LoggerConfig{
				Output:       NewMemoryWriter(),
				LogFormatter: defaultLogFormatter,
			},
		},
		{
			name: "custom log formatter",
			args: LoggerConfig{
				LogFormatter: func(l LogFormatterParams) string {
					return "test"
				},
			},
			want: LoggerConfig{
				Output:       os.Stdout,
				LogFormatter: func(l LogFormatterParams) string { return "test" },
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := saneConfig(tt.args)

			assert.NotNilf(t, got.Output, "saneConfig() = %v, want %v", got.Output, tt.want.Output)

			assert.NotNilf(t, got.LogFormatter, "saneConfig() = %v, want %v", got.LogFormatter, tt.want.LogFormatter)
		})
	}
}

func Test_log(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args LoggerConfig
		want []string
	}{
		{
			name: "default",
			args: saneConfig(LoggerConfig{Output: NewMemoryWriter()}),
			want: []string{
				time.Now().Format(time.DateOnly),
				"HTTP/1",
				"200",
				"192.",
				"GET",
				"/",
				"application/json",
				"text/plain",
			},
		},
		{
			name: "json log format",
			args: saneConfig(
				LoggerConfig{
					Output: NewMemoryWriter(),
					LogFormatter: func(l LogFormatterParams) string {
						jsonOutput, _ := json.Marshal(l)

						return string(jsonOutput)
					},
				},
			),
			want: []string{
				fmt.Sprintf(`"timestamp":"%sT`, time.Now().Format(time.DateOnly)),
				`"protocol_version":1`,
				`"status_code":200`,
				`"client_ip":"192.`,
				`"method":"GET"`,
				`"path":"/"`,
				`"request_content_type":"application/json"`,
				`"response_content_type":"text/plain"`,
				`"latency":`,
			},
		},
		{
			name: "stringer log format",
			args: saneConfig(
				LoggerConfig{
					Output: NewMemoryWriter(),
					LogFormatter: func(l LogFormatterParams) string {
						return fmt.Sprintf("%s", l)
					},
				},
			),
			want: []string{
				fmt.Sprintf(`Timestamp: %s`, time.Now().Format(time.DateOnly)),
				`ProtocolVersion: 1`,
				`StatusCode: 200`,
				`ClientIP: 192.`,
				`Method: GET`,
				`Path: /`,
				`RequestContentType: application/json`,
				`ResponseContentType: text/plain`,
				`Latency: `,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := func(w http.ResponseWriter, r *http.Request, next Next) {
				log(tt.args, w, r, next)
			}

			handler := func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}

			middlewares := Middlewares{middleware}

			handler = middlewares.Chain(handler)

			mrw := &ResponseWriter{ResponseWriter: httptest.NewRecorder()}

			mr := httptest.NewRequest(http.MethodGet, "/", nil)

			mr.Header.Set("Content-Type", "application/json")

			handler(mrw, mr)

			mw := tt.args.Output.(*MemoryWriter)

			for _, each := range tt.want {
				assert.Containsf(t, mw.Content, each, "log() = %v, want %v", mw.Content, each)
			}
		})
	}
}

func NewMemoryWriter() *MemoryWriter {
	return new(MemoryWriter)
}

type MemoryWriter struct {
	Content string
}

func (m *MemoryWriter) Write(p []byte) (n int, err error) {
	m.Content = string(p)

	return len(p), nil
}
