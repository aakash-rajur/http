package params

import "context"

type Params map[string]string

func (p Params) Get(key, fallback string) string {
	value, ok := p[key]

	if !ok {
		return fallback
	}

	return value
}

func (p Params) WithinContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, paramsKey, p)
}

func FromContext(ctx context.Context) (Params, bool) {
	if ctx == nil {
		return nil, false
	}

	p, ok := ctx.Value(paramsKey).(Params)

	return p, ok
}

const paramsKey = "http_params"
