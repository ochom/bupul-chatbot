package main

import (
	"fmt"
	"os"
)

// GetEnv ...
func GetEnv(name, val string) string {
	v := os.Getenv(name)
	if v == "" {
		return val
	}

	return v
}

// MustGetEnv ...
func MustGetEnv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", fmt.Errorf("env %s not defined", name)
	}

	return v, nil
}
