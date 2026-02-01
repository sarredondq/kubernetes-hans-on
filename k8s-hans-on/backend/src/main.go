package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	HostName string    `json:"hostname"`
}

func main() {
	http.HandleFunc("/", serverHTTP)

	log.Println("Server starting on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func serverHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn{
		Time:     time.Now(),
		HostName: os.Getenv("HOSTNAME"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
