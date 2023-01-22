package models

import "api-postgres/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `DELETE FROM todos WHERE id = $1`
	res, err := conn.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
