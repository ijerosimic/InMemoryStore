package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	store := NewInMemoryStore()
	store.Populate()

	handler := &Server{store: store}
	err := http.ListenAndServe(":5000", handler)

	fmt.Println("Server listening...")

	if err != nil {
		fmt.Println("Server died")
		return
	} else {
		os.Exit(1)
	}
}
