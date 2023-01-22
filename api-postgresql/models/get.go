package models

import "api-postgres/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM todos WHERE id=$1`
	err = conn.QueryRow(sql, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
