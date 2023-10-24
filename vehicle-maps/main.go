package main

import (
	"net/http"
	//"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// env variables
	godotenv.Load(".env")

	portString := os.Getenv("PORT");
	if (portString == "") {
		log.Fatal("PORT env variable not found")
	}


	router := chi.NewRouter()


	// middlewares
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:		[]string{"https://*", "http://*"},
		AllowedMethods:		[]string{"GET","POST","PUT","DELETE"},
		AllowedHeaders:		[]string{"*"},
		ExposedHeaders:		[]string{"Link"},
		AllowCredentials:	false,
		MaxAge:				300,
	}))

	router.Use(middleware.Logger)


	// static files route definition
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)


	// route definition
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})


	// server start
	srv := &http.Server{
		Handler: 	router,
		Addr: 		":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}