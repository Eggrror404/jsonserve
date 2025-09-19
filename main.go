package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Check if servers.json exists in current directory
	if _, err := os.Stat("servers.json"); os.IsNotExist(err) {
		log.Fatalf("servers.json not found in current directory")
	}

	// Serve servers.json on /servers.json
	http.HandleFunc("/servers.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "servers.json")
	})

	addr := ":8080"
	log.Printf("Serving servers.json at http://localhost%s/servers.json", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

