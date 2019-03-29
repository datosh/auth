package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! v1 %s", time.Now())
}

func main() {
	log.Printf("Serving on :8888")
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8888", nil)
}
