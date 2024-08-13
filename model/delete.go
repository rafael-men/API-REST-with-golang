package model

import "github.com/rafael-men/rest-api-with-golang/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE title=$2 ,id`)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
