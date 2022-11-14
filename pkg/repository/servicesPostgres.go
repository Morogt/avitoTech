package repository

import (
	"avitoTech"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ServicesPostgres struct {
	db *sqlx.DB
}

func NewServicesPostgres(db *sqlx.DB) *ServicesPostgres {
	return &ServicesPostgres{db: db}
}

func (r *ServicesPostgres) GetReportByServ() (posts []avitoTech.Report, err error) {
	query := fmt.Sprintf("SELECT id_service, sum(amount) FROM %s WHERE id_service != 0 GROUP BY id_service ", orderTable)
	row, err := r.db.Query(query)
	if err != nil {
		return posts, err
	}

	for row.Next() {
		var rep avitoTech.Report
		err := row.Scan(&rep.ID, &rep.Amount)
		if err != nil {
			return posts, err
		}
		posts = append(posts, rep)
	}

	return posts, nil
}
