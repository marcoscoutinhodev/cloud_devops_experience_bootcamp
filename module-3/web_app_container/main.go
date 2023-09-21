package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World"))
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

