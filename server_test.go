package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	store := StubStore{
		mu: sync.RWMutex{},
		data: map[string]string{
			"session_1": "11111",
			"session_2": "22222",
		},
	}
	server := &Server{store}

	t.Run("returns store item by key", func(t *testing.T) {
		request := newGetSessionByIdRequest("session_1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "11111"

		assertResponseBody(t, got, want)
	})
	t.Run("returns store item by different key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/sessions/session_2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "22222"

		assertResponseBody(t, got, want)
	})
}

func TestSet(t *testing.T) {
	store := StubStore{
		mu:   sync.RWMutex{},
		data: make(map[string]string)}
	server := &Server{store}

	t.Run("inserts key value pair into store", func(t *testing.T) {
		request := newSetSession("session_3", "33333")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "11111"

		assertResponseBody(t, got, want)
	})
}

func newSetSession(key string, value string) *http.Request {
	payload, _ := json.Marshal(value)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/sessions/%s", key), bytes.NewBuffer(payload))
	return req
}

func newGetSessionByIdRequest(sessionId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/sessions/%s", sessionId), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// server_test.go
type StubStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func (s StubStore) Get(id string) string {
	score := s.data[id]
	return score
}

func (s StubStore) Set(id string, value string) string {
	s.data[id] = value
	return id
}
