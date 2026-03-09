package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gmburov/service-autocompletion/internal/api"
	"github.com/gmburov/service-autocompletion/internal/config"
)

func main() {
	cfg, err := config.Load(os.Getenv("AC_CONFIG"))
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	handler := api.NewHandler(cfg)
	addr := cfg.Addr()

	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, handler.Router()); err != nil {
		log.Fatal(err)
	}
}
