package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/pjvalent/BookLens/handlers"
	"github.com/pjvalent/BookLens/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")

	//get the port string from .env
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in the environment")
	}

	//get the db string from .env
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB String not found in the environment")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Can't connect to database.")
	}

	apiCfg := handlers.ApiConfig{
		DB: database.New(conn),
	}

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
	v1Router.Get("/healthz", handlers.HandleReadiness)
	v1Router.Get("/err", handlers.HandlerErr)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUserByApiKey))

	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Post("/createBook", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateBook))

	//mount the v1 router to the /v1 path, which itself is mapped to the /healthz path (/v1/ready)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Starting server on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
