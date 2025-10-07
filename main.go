package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Hello World endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Hello World!")
	})

	// Crash endpoint
	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Goodbye!")
		go func() {
			time.Sleep(100 * time.Millisecond) // Give response time to send
			os.Exit(1)
		}()
	})

	// Health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Hello World")
	})

	// Start server
	port := ":3000"
	fmt.Printf("🎯 Starting server on http://localhost%s\n", port)
	fmt.Println("📊 Visit http://localhost:3000 for Hello World")
	fmt.Println("💥 Visit http://localhost:3000/crash to crash the server")

	log.Fatal(http.ListenAndServe(port, nil))
}

var startTime = time.Now()
