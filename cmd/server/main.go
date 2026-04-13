package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sakib-maho/golang-beego-restapi-swagger-v1/internal/api"
	"github.com/sakib-maho/golang-beego-restapi-swagger-v1/internal/store"
)

func main() {
	addr := os.Getenv("APP_ADDRESS")
	if addr == "" {
		addr = ":8080"
	}

	taskStore := store.NewTaskStore()
	handler := api.NewHandler(taskStore)
	router := api.NewRouter(handler)

	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
