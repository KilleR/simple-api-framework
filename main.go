package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := NewSimpleApiFramework()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello world")
		if err != nil {
			log.Println("Client went away:", err)
		}
	})

	r.HandleFunc("/ping", pingHandler)

	log.Fatalln(r.Start())
}
