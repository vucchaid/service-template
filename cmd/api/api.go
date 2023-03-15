package api

import (
	"log"
	"net/http"
	"service-template/cmd/api/routes"
)

func ServeHTTP(host string) error {
	router := routes.NewRouter()
	log.Println("[INFO]:Server started (", host, ")")
	return http.ListenAndServe(host, router)
}
