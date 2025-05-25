package models

import (
	"errors"
	"time"
)


type User struct{
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(u *User) error{
	stmt, err := DB.Prepare("INSERT INTO users(email, password) VALUES (?,?)")
	if err!= nil{
		return err
	}
	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil{
		return err
	}
	id, _ := res.LastInsertId()
	u.ID = int(id)
	return nil
}

func GetUserByEmail(email string) (*User, error) {
    var u User
    err := DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).
        Scan(&u.ID, &u.Email, &u.Password)

    if err != nil {
        return nil, errors.New("usuario no encontrado")
    }
    return &u, nil
}

func SavePasswordResetToken(email, token string, duration time.Duration) error {
    expiresAt := time.Now().Add(duration)
    _, err := DB.Exec("INSERT INTO password_resets (email, token, expires_at) VALUES (?, ?, ?)",
        email, token, expiresAt)
    return err
}

func GetEmailByValidToken(token string) (string, error) {
    var email string
    var expiresAt time.Time

    err := DB.QueryRow(`SELECT email, expires_at FROM password_resets WHERE token = ?`, token).
        Scan(&email, &expiresAt)
    if err != nil || time.Now().After(expiresAt) {
        return "", errors.New("token inválido o expirado")
    }
    return email, nil
}




