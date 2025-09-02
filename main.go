package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Hello World endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Hello World")
	})

	// Crash endpoint
	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		panic("Intentional crash triggered by /crash endpoint")
	})

	// Health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Hello World")
	})

	// Start server
	port := ":8080"
	fmt.Printf("🎯 Starting server on http://localhost%s\n", port)
	fmt.Println("📊 Visit http://localhost:8080 for Hello World")
	fmt.Println("💥 Visit http://localhost:8080/crash to crash the server")

	log.Fatal(http.ListenAndServe(port, nil))
}

var startTime = time.Now()
