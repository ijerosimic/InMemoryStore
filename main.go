package main

import (
	"log"
	"net/http"
)

type InMemorySessionStore struct {
	data map[string]string
}

func (i *InMemorySessionStore) Get(key string) (value string) {
	return i.data[key]
}

func (i *InMemorySessionStore) Set(key string, value string) {
	i.data[key] = value
}

func (i *InMemorySessionStore) GetSession(sessionId string) (session string) {
	return i.Get(sessionId)
}

func main() {
	store := &InMemorySessionStore{
		data: make(map[string]string)}
	store.Set("session_1", "session_1")
	handler := &SessionServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", handler))
}
