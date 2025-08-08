package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Start a goroutine that will crash the server after 10 seconds
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("💥 CRASH TIME! Server has been running for 10 seconds")
		panic("Intentional crash after 10 seconds of uptime")
	}()

	// Simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(startTime)
		response := fmt.Sprintf(`
🚀 Go Crasher Server

Uptime: %v
Time until crash: %v

This server will crash in %d seconds!
`, uptime, 10*time.Second-uptime, int((10*time.Second - uptime).Seconds()))

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, response)
	})

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(startTime)
		timeLeft := 10*time.Second - uptime

		if timeLeft <= 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "💀 About to crash!")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "✅ Healthy - %v until crash", timeLeft)
	})

	// Start server
	port := ":8080"
	fmt.Printf("🎯 Starting crasher server on http://localhost%s\n", port)
	fmt.Println("⏰ Server will crash after exactly 10 seconds")
	fmt.Println("📊 Visit http://localhost:8080 to see countdown")
	fmt.Println("🏥 Visit http://localhost:8080/health for health check")

	log.Fatal(http.ListenAndServe(port, nil))
}

var startTime = time.Now()
