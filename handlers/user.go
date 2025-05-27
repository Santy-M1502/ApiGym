package handlers

import (
	"ApiGym/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
    smtpHost = "smtp.gmail.com"
    smtpPort = "587"
    smtpUser = "serviceapigym@gmail.com"
    smtpPass = "ofgm eaft dsei czkq"
)


func VerificarEmailHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var datos struct {
        Email string `json:"email"`
    }

    bodyBytes, _ := io.ReadAll(r.Body)
    fmt.Println("Cuerpo recibido:", string(bodyBytes))
    r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

    if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
        fmt.Println("Error al decodificar JSON:", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]any{
            "ok":    false,
            "error": "Datos inválidos",
        })
        return
    }

    if datos.Email == "" {
        fmt.Println("Email vacío")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]any{
            "ok": false,
            "error": "El email está vacío",
        })
        return
    }

    _, err := models.GetUserByEmail(datos.Email)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]any{
            "ok":     true,
            "existe": false,
        })
        return
    }

    token := uuid.NewString()
    err = models.SavePasswordResetToken(datos.Email, token, time.Hour)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]any{
            "ok": false,
            "error": "No se pudo generar el token",
        })
        return
    }

    err = EnviarEmailRecuperacion(datos.Email, token)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]any{
            "ok": false,
            "error": "No se pudo enviar el correo",
        })
        return
    }

    json.NewEncoder(w).Encode(map[string]any{
        "ok":     true,
        "existe": true,
        "token":  token,
    })
}


func EnviarEmailRecuperacion(destinatario, token string) error {
    from := smtpUser
    to := []string{destinatario}
    msg := []byte(fmt.Sprintf(
        "Subject: Recuperación de contraseña\n\nHola,\n\nRecibimos una solicitud para restablecer tu contraseña.\n\nHaz clic en este enlace para continuar:\n\nhttp://localhost:5500/frontend/reset-password.html?token=%s\n\nSi no fuiste tú, ignora este mensaje.", token,
    ))
    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
    return err
}

func ResetearPasswordHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var datos struct {
        Token    string `json:"token"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]any{"ok": false, "error": "Datos inválidos"})
        return
    }

    email, err := models.GetEmailByValidToken(datos.Token)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]any{"ok": false, "error": "Token inválido o vencido"})
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(datos.Password), bcrypt.DefaultCost)
    _, err = models.DB.Exec("UPDATE users SET password = ? WHERE email = ?", hashed, email)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]any{"ok": false, "error": "No se pudo actualizar la contraseña"})
        return
    }

    models.DB.Exec("DELETE FROM password_resets WHERE token = ?", datos.Token)

    json.NewEncoder(w).Encode(map[string]any{"ok": true})
}

func VerificarEmailRegister(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var datos struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }

    bodyBytes, _ := io.ReadAll(r.Body)
    r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

    if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
        fmt.Println("Error al decodificar JSON:", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]any{
            "ok":    false,
            "error": "Datos inválidos",
        })
        return
    }

    if datos.Email == "" {
        fmt.Println("Email vacío")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]any{
            "ok": false,
            "error": "El email está vacío",
        })
        return
    }

    _, err := models.GetUserByEmail(datos.Email)
    if err != nil {
        json.NewEncoder(w).Encode(map[string]any{
            "ok":     true,
            "existe": false,
        })
        return
    }

    json.NewEncoder(w).Encode(map[string]any{
        "ok":     true,
        "existe": true,
    })

}