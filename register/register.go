package register

import (
	"errors"
	p "github.com/aakash-rajur/http/params"
	"net/http"
	"slices"
)

func NewRegister() Register {
	return make(Register, 0)
}

type Register []Entry

func (r Register) Add(pattern string, handler http.Handler) Register {
	ss := segmentsFromPath(pattern)

	entry := Entry{
		segments: ss,
		Handler:  handler,
	}

	updated := append(r, entry)

	slices.SortStableFunc(
		updated,
		func(a, b Entry) int {
			value, _ := a.segments.cmp(b.segments, 0)

			return value
		},
	)

	return updated
}

func (r Register) Find(pattern string) (Entry, p.Params, error) {
	if len(r) == 0 {
		return Entry{}, p.Params{}, ErrNotFound
	}

	safePath := cleanPath(pattern)

	ss := segmentsFromPath(safePath)

	left, right := 0, len(r)-1

	cursor := 0

	for left <= right {
		mid := left + (right-left)/2

		entry := r[mid]

		comparison, last := entry.segments.cmp(ss, cursor)

		if comparison == 0 {
			params := entry.segments.params(ss)

			return entry, params, nil
		}

		cursor = last

		// mid > target, go left
		if comparison > 0 {
			right = mid - 1

			continue
		}

		// mid < target, go right
		if comparison < 0 {
			left = mid + 1

			continue
		}
	}

	return Entry{}, nil, ErrNotFound
}

type Entry struct {
	segments segments
	Handler  http.Handler
}

var ErrNotFound = errors.New("entry not found")
