package main

import (
	"fmt"
	"sync"
)

type InMemoryStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewInMemoryStore() (store *InMemoryStore) {
	return &InMemoryStore{
		mu:   sync.RWMutex{},
		data: make(map[string]string)}
}

func (i *InMemoryStore) Populate() {
	i.Set("session_1", "11111")
	i.Set("session_2", "22222")
	i.Set("session_3", "33333")
	i.Set("session_4", "44444")
	i.Set("session_5", "55555")
}

func (i *InMemoryStore) Set(id string, value string) string {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.data[id] = value
	return id
}

func (i *InMemoryStore) Get(id string) (session string) {
	i.mu.RLock()
	defer i.mu.Unlock()
	if _, exists := i.data[id]; exists {
		return i.data[id]
	} else {
		fmt.Printf("Get - Not found")
		return ""
	}
}

func (i *InMemoryStore) UpdateSession(sessionId string, newValue string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, exists := i.data[sessionId]; exists {
		i.data[sessionId] = newValue
	} else {
		fmt.Printf("Update - Not found")
	}
}

func (i *InMemoryStore) DeleteSession(sessionId string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, exists := i.data[sessionId]; exists {
		delete(i.data, sessionId)
	} else {
		fmt.Printf("Delete - Not found")
	}
}
