package main

import (
//	"api/router"
	"fmt"
	"log"
	"net/http"
	"api/middleware"
	"api/router"
	"github.com/rs/cors"
)

func main() {
	port := "8080"

	db := middleware.PostgresConnection()

	defer db.Close()

	mux := http.NewServeMux()


	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://www.swellstatus.com", "http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPut},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Connection", "Host", "Origin", "User-Agent", "Referer", "Cache-Control", "X-header"},
	})

	mux.Handle("/", router.Handlers())

	// Insert the middleware
	handler := c.Handler(mux)
	

	fmt.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, handler))	
}
