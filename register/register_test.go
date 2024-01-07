package register

import (
	"github.com/aakash-rajur/http/params"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegister_Add(t *testing.T) {
	t.Parallel()

	type args struct {
		pattern string
		handler http.Handler
	}

	tests := []struct {
		name string
		r    Register
		args args
		want Register
	}{
		{
			name: "Add to empty register",
			r:    NewRegister(),
			args: args{
				pattern: "/",
				handler: nil,
			},
			want: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
			},
		},
		{
			name: "Add to non-empty register",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api",
				handler: nil,
			},
			want: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
			},
		},
		{
			name: "Add to non-empty register with same pattern",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api",
				handler: nil,
			},
			want: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
			},
		},
		{
			name: "Add to non-empty register with params",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api/{id}",
				handler: nil,
			},
			want: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
				{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
			},
		},
		{
			name: "Add to non-empty register with params and same pattern",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
				{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api/{id}",
				handler: nil,
			},
			want: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
				{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
				{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.Add(tt.args.pattern, tt.args.handler)

			assert.Equalf(t, tt.want, got, "Add() = %v, want %v", got, tt.want)
		})
	}
}

func TestRegister_Find(t *testing.T) {
	t.Parallel()

	type args struct {
		pattern string
	}

	type want struct {
		entry  Entry
		params params.Params
		err    error
	}

	tests := []struct {
		name string
		r    Register
		args args
		want want
	}{
		{
			name: "Find in empty register",
			r:    NewRegister(),
			args: args{
				pattern: "/",
			},
			want: want{
				entry:  Entry{},
				params: params.Params{},
				err:    ErrNotFound,
			},
		},
		{
			name: "Find in non-empty register",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api",
			},
			want: want{
				entry: Entry{
					segments: segments{"api"},
					Handler:  nil,
				},
				params: params.Params{},
				err:    nil,
			},
		},
		{
			name: "Find in non-empty register with params",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
				{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api/123",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
				params: params.Params{
					"id": "123",
				},
				err: nil,
			},
		},
		{
			name: "Find in non-empty register with params and same pattern",
			r: Register{
				{
					segments: segments{""},
					Handler:  nil,
				},
				{
					segments: segments{"api"},
					Handler:  nil,
				},
				{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
			},
			args: args{
				pattern: "/api/123",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "{id}"},
					Handler:  nil,
				},
				params: params.Params{
					"id": "123",
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry, params, err := tt.r.Find(tt.args.pattern)

			assert.Equalf(t, tt.want.entry, entry, "Find() entry = %v, want %v", entry, tt.want.entry)

			assert.Equalf(t, tt.want.params, params, "Find() params = %v, want %v", params, tt.want.params)

			assert.Equalf(t, tt.want.err, err, "Find() err = %v, want %v", err, tt.want.err)
		})
	}
}
