package http

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func Logger(config LoggerConfig) Middleware {
	cfg := saneConfig(config)

	return func(w http.ResponseWriter, r *http.Request, next Next) {
		log(cfg, w, r, next)
	}
}

func saneConfig(in LoggerConfig) LoggerConfig {
	out := LoggerConfig{
		Output:          os.Stdout,
		LogFormat:       "%v | HTTP/%d | %4d | %10v | %30s | %30s | %15s | %7s %-7s \n",
		TimestampFormat: time.DateTime,
		LogFormatter: func(config LoggerConfig, params LogFormatterParams) string {
			return fmt.Sprintf(
				config.LogFormat,
				params.Timestamp.Format(config.TimestampFormat),
				params.ProtocolVersion,
				params.StatusCode,
				params.Latency,
				params.RequestContentType,
				params.ResponseContentType,
				params.ClientIP,
				params.Method,
				params.Path,
			)
		},
	}

	if in.Output != nil {
		out.Output = in.Output
	}

	if in.LogFormat != "" {
		out.LogFormat = in.LogFormat
	}

	if in.TimestampFormat != "" {
		out.TimestampFormat = in.TimestampFormat
	}

	if in.LogFormatter != nil {
		out.LogFormatter = in.LogFormatter
	}

	return out
}

func log(
	cfg LoggerConfig,
	w http.ResponseWriter,
	r *http.Request,
	next Next,
) {
	start := time.Now()

	next(r)

	end := time.Now()

	statusCode := 0

	hw, ok := w.(*ResponseWriter)

	if ok {
		statusCode = hw.StatusCode
	}

	hw.Header().Get("Content-Type")

	clientIps := make([]string, 0)

	xff := r.Header.Get("X-Forwarded-For")

	if strings.TrimSpace(xff) != "" {
		clientIps = append(clientIps, xff)
	}

	xri := r.Header.Get("X-Real-IP")

	if strings.TrimSpace(xri) != "" {
		clientIps = append(clientIps, xri)
	}

	ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

	if strings.TrimSpace(ip) != "" {
		clientIps = append(clientIps, ip)
	}

	clientIp := strings.Join(clientIps, ", ")

	requestContentType := r.Header.Get("Content-Type")

	if requestContentType == "" {
		requestContentType = "text/plain"
	}

	requestContentEncoding := r.Header.Get("Content-Encoding")

	if requestContentEncoding == "" {
		requestContentEncoding = "identity"
	}

	responseContentType := w.Header().Get("Content-Type")

	if responseContentType == "" {
		responseContentType = "text/plain"
	}

	responseContentEncoding := w.Header().Get("Content-Encoding")

	if responseContentEncoding == "" {
		responseContentEncoding = "identity"
	}

	params := LogFormatterParams{
		Timestamp:               end,
		StatusCode:              statusCode,
		Latency:                 end.Sub(start),
		ClientIP:                clientIp,
		Method:                  r.Method,
		Path:                    r.URL.Path,
		Query:                   r.URL.Query(),
		RequestContentType:      requestContentType,
		RequestContentEncoding:  requestContentEncoding,
		ResponseContentType:     responseContentType,
		ResponseContentEncoding: responseContentEncoding,
		ProtocolVersion:         r.ProtoMajor,
	}

	_, _ = fmt.Fprint(cfg.Output, cfg.LogFormatter(cfg, params))
}

type LoggerConfig struct {
	Output          io.Writer
	LogFormat       string
	TimestampFormat string
	LogFormatter    func(LoggerConfig, LogFormatterParams) string
}

type LogFormatterParams struct {
	Timestamp               time.Time           `json:"timestamp"`
	StatusCode              int                 `json:"status_code,omitempty"`
	Latency                 time.Duration       `json:"latency,omitempty"`
	ClientIP                string              `json:"client_ip,omitempty"`
	Method                  string              `json:"method,omitempty"`
	Path                    string              `json:"path,omitempty"`
	Query                   map[string][]string `json:"query,omitempty"`
	RequestContentType      string              `json:"request_content_type,omitempty"`
	RequestContentEncoding  string              `json:"request_content_encoding,omitempty"`
	ResponseContentType     string              `json:"response_content_type,omitempty"`
	ResponseContentEncoding string              `json:"response_content_encoding,omitempty"`
	ProtocolVersion         int                 `json:"protocolVersion,omitempty"`
}
