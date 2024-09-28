package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port not found in the environment")
	}

	fmt.Println(portString)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//define a new router for the handler readiness to map to the /healthz path (to check if server is live and running)
	v1Router := chi.NewRouter()
	//scope the endpoint to only be get
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	//mount the v1 router to the /v1 path, which itself is mapped to the /healthz path (/v1/ready)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Starting server on port %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
