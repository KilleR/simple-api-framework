package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	signingResponseChan chan string
)

func init() {
	signingResponseChan = make(chan string)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello world")
		if err != nil {
			log.Println("Client went away:", err)
		}
	})

	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/sign", signingHandler)
	r.HandleFunc("/getsigned", signingResponseHandler)

	// set up CORS
	headersOk := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "X-Auth-Token", "Authorization", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatalln(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
