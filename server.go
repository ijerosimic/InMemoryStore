package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Store interface {
	Get(id string) (session string)
	Set(id string, value string) (session string)
}
type Server struct {
	store Store
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/sessions/")
	value := ""
	switch r.Method {
	case http.MethodGet:
		value = s.handleGet(id)
	case http.MethodPost:
		value = s.handlePost(id, r)
	}
	fmt.Fprint(w, value)
}

func (s *Server) handleGet(id string) string {
	return s.store.Get(id)
}

type payload struct {
	Value string
}

func (s *Server) handlePost(id string, r *http.Request) string {
	decoder := json.NewDecoder(r.Body)
	var payload Payload
	err := decoder.Decode(&payload)
	if err != nil {
		fmt.Printf("POST - Error decoding json")
		return ""
	}

	//scanner := bufio.NewScanner(r.Body)
	//for i := 0; scanner.Scan() && i < 5; i++ {
	//	fmt.Println(scanner.Text())
	//}
	//if err := scanner.Err(); err != nil {
	//	panic(err)
	//}

	s.store.Set(id, payload.Val)
	return id
}

type Payload struct {
	Val string `json:"val"`
}
