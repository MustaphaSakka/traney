package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//define routes
	http.HandleFunc("/start", start)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Let's start and enjoy this GO project!")
}
