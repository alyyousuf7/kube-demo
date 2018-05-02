package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm being served by: %s\n", os.Getenv("POD_NAME"))
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", pageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
