package http

import (
	"net/http"
)

func NewClient() (*http.Client, error) {
	return &http.Client{}, nil
}