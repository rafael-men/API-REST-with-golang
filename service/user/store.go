package user

import (
	"database/sql"
	"github.com/rafael-men/rest-api-with-golang/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	u := &types.User{}

	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}
