package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (s *SessionServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sessionId := strings.TrimPrefix(r.URL.Path, "/session/")
	value := s.store.GetSession(sessionId)
	fmt.Fprint(w, value)
}

type SessionStore interface {
	GetSession(sessionId string) (session string)
}
type SessionServer struct {
	store SessionStore
}
