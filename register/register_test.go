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

	type testCase struct {
		name string
		r    Register
		args args
		want want
	}

	complexRegistry := NewRegister().
		Add("/", nil).
		Add("/health", nil).
		Add("/api/v2/books", nil).
		Add("/api/v2/books/{bookId}", nil).
		Add("/api/v2/users", nil).
		Add("/api/v2/users/{userId}", nil).
		Add("/api/v2/users/{userId}/books", nil).
		Add("/api/v2/rpc/{service}/{method}", nil).
		Add("/identity", nil).
		Add("/identity/{id}", nil).
		Add("/public", nil).
		Add("/private", nil)

	tests := []testCase{
		{
			name: "Find in empty register",
			r:    registerFromPatterns([]string{}, nil),
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
			r:    registerFromPatterns([]string{"/", "/api"}, nil),
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
			r:    registerFromPatterns([]string{"/", "/api", "/api/{id}"}, nil),
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
			r:    registerFromPatterns([]string{"/", "/api", "/api/{id}", "/api/{id}"}, nil),
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
			name: "Find in complex register: 1",
			r:    complexRegistry,
			args: args{
				pattern: "/",
			},
			want: want{
				entry: Entry{
					segments: segments{""},
					Handler:  nil,
				},
				params: params.Params{},
			},
		},
		{
			name: "Find in complex register: 2",
			r:    complexRegistry,
			args: args{
				pattern: "/health",
			},
			want: want{
				entry: Entry{
					segments: segments{"health"},
					Handler:  nil,
				},
				params: params.Params{},
			},
		},
		{
			name: "Find in complex register: 3",
			r:    complexRegistry,
			args: args{
				pattern: "/api/v2/books",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "v2", "books"},
					Handler:  nil,
				},
				params: params.Params{},
			},
		},
		{
			name: "Find in complex register: 4",
			r:    complexRegistry,
			args: args{
				pattern: "/api/v2/books/10",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "v2", "books", "{bookId}"},
					Handler:  nil,
				},
				params: params.Params{
					"bookId": "10",
				},
			},
		},
		{
			name: "Find in complex register: 5",
			r:    complexRegistry,
			args: args{
				pattern: "/api/v2/users",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "v2", "users"},
					Handler:  nil,
				},
				params: params.Params{},
			},
		},
		{
			name: "Find in complex register: 6",
			r:    complexRegistry,
			args: args{
				pattern: "/api/v2/users/10",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "v2", "users", "{userId}"},
					Handler:  nil,
				},
				params: params.Params{
					"userId": "10",
				},
			},
		},
		{
			name: "Find in complex register: 7",
			r:    complexRegistry,
			args: args{
				pattern: "/api/v2/users/10/books",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "v2", "users", "{userId}", "books"},
					Handler:  nil,
				},
				params: params.Params{
					"userId": "10",
				},
			},
		},
		{
			name: "Find in complex register: 8",
			r:    complexRegistry,
			args: args{
				pattern: "/api/v2/rpc/service/method",
			},
			want: want{
				entry: Entry{
					segments: segments{"api", "v2", "rpc", "{service}", "{method}"},
					Handler:  nil,
				},
				params: params.Params{
					"service": "service",
					"method":  "method",
				},
			},
		},
		{
			name: "Find in complex register: 9",
			r:    complexRegistry,
			args: args{
				pattern: "/identity",
			},
			want: want{
				entry: Entry{
					segments: segments{"identity"},
					Handler:  nil,
				},
				params: params.Params{},
			},
		},
		{
			name: "Find in complex register: 10",
			r:    complexRegistry,
			args: args{
				pattern: "/identity/10",
			},
			want: want{
				entry: Entry{
					segments: segments{"identity", "{id}"},
					Handler:  nil,
				},
				params: params.Params{
					"id": "10",
				},
			},
		},
		{
			name: "Find in complex register: 11",
			r:    complexRegistry,
			args: args{
				pattern: "/public",
			},
			want: want{
				entry: Entry{
					segments: segments{"public"},
					Handler:  nil,
				},
				params: params.Params{},
			},
		},
		{
			name: "Find in complex register: 12",
			r:    complexRegistry,
			args: args{
				pattern: "/private",
			},
			want: want{
				entry: Entry{
					segments: segments{"private"},
					Handler:  nil,
				},
				params: params.Params{},
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

func BenchmarkRegister_Find(b *testing.B) {
	r := NewRegister().
		Add("/health", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/books", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/books/{bookId}", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/users", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/users/{userId}", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/users/{userId}/books", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/rpc/{service}/{method}", http.HandlerFunc(http.NotFound))

	type args struct {
		pattern string
		params  params.Params
		err     error
	}

	testCases := []args{
		{
			pattern: "/health",
			params:  params.Params{},
			err:     nil,
		},
		{
			pattern: "/api/v2/books",
			params:  params.Params{},
			err:     nil,
		},
		{
			pattern: "/api/v2/books/10",
			params: params.Params{
				"bookId": "10",
			},
			err: nil,
		},
		{
			pattern: "/api/v2/users",
			params:  params.Params{},
			err:     nil,
		},
		{
			pattern: "/api/v2/users/10",
			params: params.Params{
				"userId": "10",
			},
		},
		{
			pattern: "/api/v2/users/10/books",
			params: params.Params{
				"userId": "10",
			},
			err: nil,
		},
		{
			pattern: "/api/v2/rpc/service/method",
			params: params.Params{
				"service": "service",
				"method":  "method",
			},
			err: nil,
		},
		{
			pattern: "/not-found",
			params:  params.Params(nil),
			err:     ErrNotFound,
		},
	}

	var entry Entry

	var p params.Params

	var err error

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			entry, p, err = r.Find(testCase.pattern)
		}
	}

	_, _, _ = entry, p, err
}

func registerFromPatterns(patterns []string, handler http.Handler) Register {
	r := NewRegister()

	for _, pattern := range patterns {
		r = r.Add(pattern, handler)
	}

	return r
}
