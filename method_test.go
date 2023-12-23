package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathWithMethod(t *testing.T) {
	type args struct {
		method string
		path   string
	}

	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				method: "GET",
				path:   "/",
			},
			want: "GET",
		},
		{
			args: args{
				method: "POST",
				path:   "/test",
			},
			want: "POST/test",
		},
		{
			args: args{
				method: "PATCH",
				path:   "/test/:id",
			},
			want: "PATCH/test/:id",
		},
		{
			args: args{
				method: "DELETE",
				path:   "/test/123",
			},
			want: "DELETE/test/123",
		},
		{
			args: args{
				method: "PUT",
				path:   "/test/:id/:id",
			},
			want: "PUT/test/:id/:id",
		},
	}

	for i, tt := range tests {
		ttName := fmt.Sprintf("Test %d", i+1)

		t.Run(ttName, func(t *testing.T) {
			got := pathWithMethod(tt.args.method, tt.args.path)

			assert.Equalf(t, tt.want, got, "pathWithMethod() = %v, want %v", got, tt.want)
		})
	}
}
