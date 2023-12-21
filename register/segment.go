package register

import (
	"cmp"
	"github.com/aakash-rajur/http/params"
	"path"
	"strings"
)

func segmentsFromPath(pattern string) segments {
	safePath := cleanPath(pattern)

	partials := strings.Split(safePath, "/")

	if partials[0] == "" {
		partials = partials[1:]
	}

	ss := make(segments, len(partials))

	for i, partial := range partials {
		ss[i] = segment(partial)
	}

	return ss
}

type segments []segment

func (s segments) params(other segments) params.Params {
	p := make(params.Params)

	for index, each := range s {
		if !each.isParam() {
			continue
		}

		if index >= len(other) {
			break
		}

		key, value := each.name(), other[index]

		p[key] = string(value)
	}

	return p
}

func (s segments) cmp(other segments, start int) (int, int) {
	sl, ol := len(s), len(other)

	length := min(sl, ol)

	for i := start; i < length; i += 1 {
		a, b := s[i], other[i]

		value := a.cmp(b)

		if value == 0 {
			continue
		}

		return value, i
	}

	normalized := max(-1, min(1, sl-ol))

	return normalized, length
}

type segment string

func (s segment) isParam() bool {
	return s[0] == ':'
}

func (s segment) name() string {
	if !s.isParam() {
		return ""
	}

	return string(s[1:])
}

func (s segment) cmp(other segment) int {
	if s.isParam() {
		return 0
	}

	return cmp.Compare(s, other)
}

func cleanPath(pattern string) string {
	safePath := path.Clean(pattern)

	if safePath[0] != '/' {
		safePath = "/" + safePath
	}

	return safePath
}
