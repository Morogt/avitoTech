package repository

import (
	"avitoTech"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user avitoTech.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2) RETURNING ID", userTable)
	row := r.db.QueryRow(query, user.ID, user.Balance)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) GetUser(id int) (avitoTech.User, error) {
	var user avitoTech.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
