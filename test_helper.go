package http

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type TestRouteParams struct {
	Method  string
	Pattern string
	Path    string
	Body    Body
	Params  map[string]any
}

type TestRoute struct {
	Name    string `json:"Name" yaml:"Name"`
	Method  string `json:"method" yaml:"method"`
	Pattern string `json:"Pattern" yaml:"Pattern"`
}

func (tr *TestRoute) generatePath(seed string) (string, map[string]any) {
	re, err := regexp.Compile("{([^}]+)}")

	if err != nil {
		return "", nil
	}

	matches := re.FindAllStringSubmatch(tr.Pattern, -1)

	path := tr.Pattern

	params := make(map[string]any)

	for _, match := range matches {
		paramMatch, paramName := match[0], match[1]

		paramSeed := paramName

		if seed != "" {
			paramSeed = seed + paramName
		}

		paramValue := generateHash(paramSeed)

		path = strings.Replace(path, paramMatch, paramValue, -1)

		params[paramName] = paramValue
	}

	return path, params
}

func (tr *TestRoute) generateBody(seed string) Body {
	payload := Body{
		"message": "Hello, World!",
		"Pattern": tr.Pattern,
		"seed":    seed,
	}

	return payload
}

func (tr *TestRoute) generateParams(seed string) TestRouteParams {
	path, params := tr.generatePath(seed)

	return TestRouteParams{
		Method:  tr.Method,
		Pattern: tr.Pattern,
		Path:    path,
		Params:  params,
		Body:    tr.generateBody(seed),
	}
}

func (tr *TestRoute) alternatePattern() string {
	re, err := regexp.Compile("{([^}]+)}")

	if err != nil {
		return ""
	}

	return re.ReplaceAllString(tr.Pattern, ":$1")
}

func generateHash(seed string) string {
	h := sha256.New()

	h.Write([]byte(seed))

	hash := h.Sum(nil)

	hashB64 := fmt.Sprintf("%.4x", hash)

	return hashB64
}

type Body map[string]any

func (b Body) Reader() io.Reader {
	jsonPayload, err := json.Marshal(b)

	if err != nil {
		return nil
	}

	return strings.NewReader(string(jsonPayload))
}

func (b Body) Map() map[string]any {
	return b
}
