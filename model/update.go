package model

import "github.com/rafael-men/rest-api-with-golang/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$2 , decsription=$3, done=$4 WHERE id=$1`, todo.Title, todo.Done, todo.Description, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
