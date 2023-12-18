package http

import (
	p "path"
)

func pathWithMethod(method, path string) string {
	return p.Join(method, path)
}
