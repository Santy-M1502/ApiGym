package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "ApiGym/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    if user.Email == "" || user.Password == "" {
	http.Error(w, "Email y contraseña requeridos", http.StatusBadRequest)
	return
    }


    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error al encriptar", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashed)

    if err := models.CreateUser(&user); err != nil {
        http.Error(w, "Error al guardar usuario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado"})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds models.User
    json.NewDecoder(r.Body).Decode(&creds)

    user, err := models.GetUserByEmail(creds.Email)
    if err != nil {
        http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
        return
    }

    var jwtKey = []byte(os.Getenv("JWT_SECRET"))

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
	http.Error(w, "Error al generar token", http.StatusInternalServerError)
	return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}