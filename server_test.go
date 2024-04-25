package main

import (
	"bytes"
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
		request := newGetRequest("session_1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "11111"

		assertResponseBody(t, got, want)
	})
	t.Run("returns store item by different key", func(t *testing.T) {
		request := newGetRequest("session_2")
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
		request := newSetRequest("session_3")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "session_3"

		assertResponseBody(t, got, want)
	})
}

func newSetRequest(key string) *http.Request {
	payload := []byte(`{
    	"id": "33333"
	}`)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/sessions/%s", key), bytes.NewBuffer(payload))
	return req
}

func newGetRequest(sessionId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/sessions/%s", sessionId), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

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
