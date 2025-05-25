package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3", "data.db")
	if err != nil{
		log.Fatal(err)
	}

	createUserTable := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	createEjercicioTable := `CREATE TABLE IF NOT EXISTS ejercicios(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT,
		descripcion TEXT,
		repeticiones INTEGER,
		usuario_id INTEGER,
		FOREIGN KEY(usuario_id) REFERENCES users(id)
	);`

	createResetTable := `CREATE TABLE IF NOT EXISTS password_resets (
		email TEXT NOT NULL,
		token TEXT NOT NULL,
		expires_at DATETIME NOT NULL
		);`

	_, err = DB.Exec(createUserTable)
	if err != nil{
		log.Fatal(err)
	}

	_, err = DB.Exec(createEjercicioTable)
	if err != nil{
		log.Fatal(err)
	}

	_, err = DB.Exec(createResetTable)
	if err != nil{
		log.Fatal(err)
	}
}