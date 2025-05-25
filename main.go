package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"ApiGym/handlers"
	"ApiGym/middleware"
	"ApiGym/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	models.InitDB()

	r := mux.NewRouter()

	// Rutas públicas
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Rutas protegidas
	r.Handle("/ejercicios", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetEjercicios))).Methods("GET")
	r.Handle("/ejercicios", middleware.AuthMiddleware(http.HandlerFunc(handlers.SetEjercicios))).Methods("POST")
	r.Handle("/ejercicios/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.ActEjercicio))).Methods("PUT")
	r.Handle("/ejercicios/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.DelEjercicio))).Methods("DELETE")
	r.HandleFunc("/verificar-email", handlers.VerificarEmailHandler).Methods("POST")
	r.HandleFunc("/api/resetear-password", handlers.ResetearPasswordHandler)

	fs := http.FileServer(http.Dir("./frontend"))
	r.PathPrefix("/").Handler(fs)

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500"}, // Cambia esto si tu frontend corre en otra URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Envolver el router con el middleware de CORS
	handler := c.Handler(r)

	log.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
