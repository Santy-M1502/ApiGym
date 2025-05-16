package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"

	"ApiGym/handlers"
	"ApiGym/middleware"
	"ApiGym/models"
)

func main()  {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error cargando .env")
	}

	models.InitDB()

	r:=mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/ejercicios", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetEjercicios))).Methods("GET")
	r.Handle("/ejercicios", middleware.AuthMiddleware(http.HandlerFunc(handlers.SetEjercicios))).Methods("POST")
	r.Handle("/ejercicios/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.DelEjercicio))).Methods("DELETE")

	log.Println("Servidor corriendo en :8080")
	http.ListenAndServe(":8080", r)
}