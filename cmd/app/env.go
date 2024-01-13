package main

import (
	"os"
	"strings"
)

func loadFromOS() Env {
	env := make(Env)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		env[pair[0]] = pair[1]
	}

	return env
}

type Env map[string]string

func (e Env) Get(key string, defaultValue string) string {
	value, ok := e[key]

	if !ok {
		return defaultValue
	}

	return value
}
