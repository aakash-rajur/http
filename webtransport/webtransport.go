package webtransport

import (
	"context"
	"errors"
	h "github.com/aakash-rajur/http"
	"github.com/quic-go/webtransport-go"
	"net/http"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (*webtransport.Session, error) {
	value := r.Context().Value(wtKey)

	wt, ok := value.(*webtransport.Server)

	if !ok {
		return nil, errors.New("unable to upgrade to web transport")
	}

	return wt.Upgrade(w, r)
}

func Middleware(wt *webtransport.Server) h.Middleware {
	return func(w http.ResponseWriter, r *http.Request, next h.Next) {
		next(withWebTransport(r, wt))
	}
}

func withWebTransport(r *http.Request, wt *webtransport.Server) *http.Request {
	ctx := context.WithValue(r.Context(), wtKey, wt)

	return r.WithContext(ctx)
}

const wtKey = "web_transport"
