package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/karthikrao/location-history/internal/server"
	"github.com/karthikrao/location-history/internal/store"
)

func main() {
	addr := os.Getenv("HISTORY_SERVER_LISTEN_ADDR")
	if addr == "" {
		addr = "8080"
	}
	log := log.New(os.Stdout, "", log.LstdFlags)
	store := store.New(log)
	server := server.New(store, log)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", addr), server.Router()))
}
