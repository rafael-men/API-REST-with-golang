package model

import "github.com/rafael-men/rest-api-with-golang/db"

func GetAll(id int64) (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
