package handlers

import (
	"ApiGym/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetEjercicios(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)
    ejercicios, err := models.GetEjerciciosPorUsuario(userID)
    if err != nil {
        http.Error(w, "Error al obtener ejercicios", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(ejercicios)
}

func SetEjercicios(w http.ResponseWriter, r *http.Request){
    userID := r.Context().Value("userID").(int)

    var e models.Ejercicio
    if err := json.NewDecoder(r.Body).Decode(&e); err != nil{
        http.Error(w, "JSON invalido", http.StatusBadRequest)
        return
    }

    e.UsuarioID = userID

    if err := models.CrearEjercicio(&e); err != nil{
        fmt.Println("Error al guardar ejercicio: ", err)
        http.Error(w, "Error al guardar ejercicio", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(e)
}

func DelEjercicio(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    idStr := vars["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil{
        http.Error(w, "ID invalido", http.StatusBadRequest)
        return
    }

    userIDRaw := r.Context().Value("userID")
    var userID int
    switch v := userIDRaw.(type){
    case int:
        userID = v
    case float64:
        userID = int(v)
    default:
        http.Error(w, "Tipo invalido de user_id", http.StatusUnauthorized)
        return
    }

    err = models.EliminarEjercicio(id, userID)
    if err != nil{
        http.Error(w, "No se pudo eliminar ejercicio", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
