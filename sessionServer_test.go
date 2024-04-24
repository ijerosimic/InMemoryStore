package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSession(t *testing.T) {
	store := StubSessionStore{
		map[string]string{
			"session_1": "session_12345",
			"session_2": "session_23456",
		},
	}
	server := &SessionServer{store}

	t.Run("returns store item by key", func(t *testing.T) {
		request := newGetSessionByIdRequest("session_1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "session"

		assertResponseBody(t, got, want)
	})
	t.Run("returns store item by different key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/session/session_2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "session"

		assertResponseBody(t, got, want)
	})
}

func newGetSessionByIdRequest(sessionId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/session/%s", sessionId), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// server_test.go
type StubSessionStore struct {
	scores map[string]string
}

func (s StubSessionStore) GetSession(sessionId string) string {
	score := s.scores[sessionId]
	return score
}
