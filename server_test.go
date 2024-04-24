package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSession(t *testing.T) {
	t.Run("returns store item by key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/sessions/session_1", nil)
		response := httptest.NewRecorder()

		Session(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
