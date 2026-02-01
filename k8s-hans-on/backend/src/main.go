package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", serverHTTP)

	log.Println("Server starting on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func serverHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := fmt.Sprintf("La hora es %s y el hostname es %s", time.Now(), os.Getenv("HOSTNAME"))
	json.NewEncoder(w).Encode(map[string]string{"message": resp})
}
