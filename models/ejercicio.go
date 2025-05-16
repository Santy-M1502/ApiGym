package models

import(
	"fmt"
)

type Ejercicio struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Repeticiones int `json:"repeticiones"`
	UsuarioID   int    `json:"usuario_id"`
}

func GetEjerciciosPorUsuario(userID int) ([]Ejercicio, error) {
	rows, err := DB.Query("SELECT id, nombre, descripcion, usuario_id FROM ejercicios WHERE usuario_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ejercicios := []Ejercicio{}
	for rows.Next() {
		var e Ejercicio
		err := rows.Scan(&e.ID, &e.Nombre, &e.Descripcion, &e.UsuarioID)
		if err != nil {
			return nil, err
		}
		ejercicios = append(ejercicios, e)
	}
	return ejercicios, nil
}

func CrearEjercicio(e *Ejercicio) error {
	stmt, err := DB.Prepare("INSERT INTO ejercicios (nombre, descripcion,repeticiones, usuario_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(e.Nombre, e.Descripcion, e.Repeticiones, e.UsuarioID)
	return err
}

func EliminarEjercicio(id int, userID int) error{
	query := `DELETE FROM
	ejercicios WHERE id = ? AND usuario_id = ?`
	stmt, err := DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	

	res, err := stmt.Exec(id, userID)
	if err != nil{
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil{
		return err
	}

	if affected == 0 {
		return fmt.Errorf("ejercicio no encontrado o no autorizado")
	}

	return nil
}