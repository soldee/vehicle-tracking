package main

import (
	"errors"
	"net/http"
	"vehicle-maps/api"
	"vehicle-maps/db"
	"vehicle-maps/response"
	"vehicle-maps/services"

	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// env variables
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT env variable not found")
	}

	// DI
	statusRepo := db.NewMongoStatusRepo(db.DBInstance())
	statusService := services.NewStatusService(statusRepo)
	subscribeService := services.NewSubscribeService(statusRepo)
	statusHandler := api.StatusHandler{
		StatusService:    statusService,
		SubscribeService: subscribeService,
	}

	router := chi.NewRouter()

	// middlewares
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	// route definition
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	router.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		response.HandleErrorResponse(w, 400, errors.New("something went wrong"))
	})
	router.Route("/vehicle", func(r chi.Router) {
		r.Get("/status", statusHandler.HandleGetStatus)
		r.Get("/status/subscribe", statusHandler.HandleSubscribe)
	})

	// server start
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
