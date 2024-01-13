package params

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
			mr := httptest.NewRequest(http.MethodGet, "/", nil)

			ctx := tc.params.WithinContext(mr.Context())

			mr = mr.WithContext(ctx)

			actual, ok := FromRequest(mr)

			assert.Truef(t, ok, "want ok == true, got %v", ok)

			assert.Equalf(t, tc.expected, actual, "want %v, got %v", tc.expected, actual)
		})
	}
}

func TestFromRequest(t *testing.T) {
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
			mr := httptest.NewRequest(http.MethodGet, "/", nil)

			ctx := tc.params.WithinContext(mr.Context())

			mr = mr.WithContext(ctx)

			actual, ok := FromRequest(mr)

			assert.Truef(t, ok, "want ok == true, got %v", ok)

			assert.Equalf(t, tc.expected, actual, "want %v, got %v", tc.expected, actual)
		})
	}
}
