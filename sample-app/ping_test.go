package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("check status code 200", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		requestRecorder := httptest.NewRecorder()
		pingHandler := http.HandlerFunc(Handler)

		pingHandler.ServeHTTP(requestRecorder, req)

		assert.Equal(t, requestRecorder.Code, http.StatusOK)
	})

	t.Run("check response body", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		requestRecorder := httptest.NewRecorder()
		pingHandler := http.HandlerFunc(Handler)

		pingHandler.ServeHTTP(requestRecorder, req)

		assert.Equal(t, requestRecorder.Body.String(), fmt.Sprintf("My response: pong"))
	})
}
