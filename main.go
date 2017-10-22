package main

import (
	"io"
	"log"
	"fmt"
	"net/http"
)

const version string = "2.0.1"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	log.Printf("Listening on port 8000...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
