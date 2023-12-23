package params

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParams_Get(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		params   Params
		key      string
		fallback string
		expected string
	}{
		{
			name:     "key exists",
			params:   Params{"foo": "bar"},
			key:      "foo",
			fallback: "baz",
			expected: "bar",
		},
		{
			name:     "key does not exist",
			params:   Params{"foo": "bar"},
			key:      "baz",
			fallback: "qux",
			expected: "qux",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.params.Get(tc.key, tc.fallback)

			assert.Equalf(t, tc.expected, actual, "want %v, got %v", tc.expected, actual)
		})
	}
}

func TestParams_WithinContext(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		params   Params
		expected Params
	}{
		{
			name:     "success",
			params:   Params{"foo": "bar"},
			expected: Params{"foo": "bar"},
		},
		{
			name:     "nil params",
			params:   nil,
			expected: nil,
		},
		{
			name:     "empty params",
			params:   Params{},
			expected: Params{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tc.params.WithinContext(context.Background())

			actual, ok := FromContext(ctx)

			assert.Truef(t, ok, "want ok == true, got %v", ok)

			assert.Equalf(t, tc.expected, actual, "want %v, got %v", tc.expected, actual)
		})
	}
}

func TestFromContext(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		ctx      context.Context
		expected Params
		ok       bool
	}{
		{
			name:     "success",
			ctx:      Params{"foo": "bar"}.WithinContext(context.Background()),
			expected: Params{"foo": "bar"},
			ok:       true,
		},
		{
			name:     "nil context",
			ctx:      nil,
			expected: nil,
			ok:       false,
		},
		{
			name:     "empty context",
			ctx:      context.Background(),
			expected: nil,
			ok:       false,
		},
		{
			name:     "wrong type",
			ctx:      context.WithValue(context.Background(), paramsKey, "foo"),
			expected: nil,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := FromContext(tc.ctx)

			assert.Equalf(t, tc.expected, actual, "want %v, got %v", tc.expected, actual)

			assert.Equalf(t, tc.ok, ok, "want %v, got %v", tc.ok, ok)
		})
	}
}
