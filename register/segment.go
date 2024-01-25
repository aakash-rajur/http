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

func (s segments) cmp(other segments) int {
	sl, ol := len(s), len(other)

	minLength := min(sl, ol)

	for i := 0; i < minLength; i += 1 {
		a, b := s[i], other[i]

		comparison := a.cmp(b)

		if comparison == 0 {
			continue
		}

		return comparison
	}

	return sl - ol
}

type segment string

func (s segment) isParam() bool {
	if len(s) == 0 {
		return false
	}

	isParam := s[0] == '{' && s[len(s)-1] == '}'

	return isParam
}

func (s segment) name() string {
	if !s.isParam() {
		return ""
	}

	return string(s[1 : len(s)-1])
}

func (s segment) cmp(other segment) int {
	sp, op := s.isParam(), other.isParam()

	if sp && !op {
		return 0
	}

	if !sp && op {
		return 1
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
