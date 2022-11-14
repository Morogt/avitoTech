package repository

import (
	"avitoTech"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) GetInfo(userId int) (user avitoTech.User, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userTable)
	err = r.db.Get(&user, query, userId)
	return user, err
}

func (r *BalancePostgres) AddMoney(userId int, balance, count float64) (user avitoTech.User, err error) {
	query := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id = $2 RETURNING *", userTable)
	row := r.db.QueryRow(query, balance, userId)
	if err = row.Scan(&user.ID, &user.Balance); err != nil {
		return user, err
	}
	query = fmt.Sprintf("INSERT INTO %s(id_user, amount, description, refill) VALUES ($1, $2, $3, true)", orderTable)
	row = r.db.QueryRow(query, userId, count, "Пополнение счёта")
	return user, err
}

func (r *BalancePostgres) WithdrawalOfMoney(userId int, balance, count float64) (user avitoTech.User, err error) {
	query := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id = $2 RETURNING *", userTable)
	row := r.db.QueryRow(query, balance, userId)
	if err = row.Scan(&user.ID, &user.Balance); err != nil {
		return user, err
	}
	query = fmt.Sprintf("INSERT INTO %s(id_user, amount, description, refill) VALUES ($1, $2, $3, false)", orderTable)
	row = r.db.QueryRow(query, userId, count, "Вывод средств")
	return user, err
}

func (r *BalancePostgres) Transfer(senderId, recipientId int, balanceSend, balanceRec, count float64) (sender avitoTech.User, recipient avitoTech.User, err error) {
	query := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id = $2 RETURNING *", userTable)
	row := r.db.QueryRow(query, balanceSend, senderId)
	if err = row.Scan(&sender.ID, &sender.Balance); err != nil {
		return sender, recipient, err
	}
	query = fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id = $2 RETURNING *", userTable)
	row = r.db.QueryRow(query, balanceRec, recipientId)
	if err = row.Scan(&recipient.ID, &recipient.Balance); err != nil {
		return sender, recipient, err
	}
	query = fmt.Sprintf("INSERT INTO %s(id_user, amount, description, refill) VALUES ($1, $2, $3, false)", orderTable)
	row = r.db.QueryRow(query, senderId, count, "Перевод пользователю")
	query = fmt.Sprintf("INSERT INTO %s(id_user, amount, description, refill) VALUES ($1, $2, $3, true)", orderTable)
	row = r.db.QueryRow(query, recipientId, count, "Перевод от пользователя")
	return sender, recipient, err
}

func (r *BalancePostgres) Pay(userId, serviceId int, balance, count float64, description string) (user avitoTech.User, orderId int, err error) {
	query := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE id = $2 RETURNING *", userTable)
	row := r.db.QueryRow(query, balance, userId)
	if err = row.Scan(&user.ID, &user.Balance); err != nil {
		return user, 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s(id_service, id_user, amount, description, refill) VALUES ($1, $2, $3, $4, false) RETURNING id", orderTable)
	row = r.db.QueryRow(query, serviceId, userId, count, description)
	if err = row.Scan(&orderId); err != nil {
		return user, 0, err
	}
	return user, orderId, err
}

func (r *BalancePostgres) History(id int) (history []avitoTech.History, err error) {
	query := fmt.Sprintf("SELECT id, id_service, amount, description, refill, order_time FROM %s WHERE id_user=$1", orderTable)

	row, err := r.db.Query(query, id)
	if err != nil {
		return history, err
	}

	for row.Next() {
		var his avitoTech.History
		err := row.Scan(&his.OrderId, &his.ServiceId, &his.Amount, &his.Description, &his.Refill, &his.Time)
		if err != nil {
			return history, err
		}
		history = append(history, his)
	}

	return history, nil
}
