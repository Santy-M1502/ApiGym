package models

import (
    "errors"
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



