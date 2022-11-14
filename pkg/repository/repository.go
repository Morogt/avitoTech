package repository

import (
	"avitoTech"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	GetInfo(userId int) (user avitoTech.User, err error)
	AddMoney(userId int, balance, count float64) (user avitoTech.User, err error)
	WithdrawalOfMoney(userId int, balance, count float64) (user avitoTech.User, err error)
	Transfer(senderId, recipientId int, balanceSend, balanceRec, count float64) (sender avitoTech.User, recipient avitoTech.User, err error)
	Pay(userId, serviceId int, balance, count float64, description string) (user avitoTech.User, orderId int, err error)
	History(id int) (history []avitoTech.History, err error)
}

type Users interface {
	CreateUser(user avitoTech.User) (int, error)
	GetUser(id int) (avitoTech.User, error)
}

type Services interface {
	GetReportByServ() (posts []avitoTech.Report, err error)
}

type Repository struct {
	Balance
	Users
	Services
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users:    NewUserPostgres(db),
		Balance:  NewBalancePostgres(db),
		Services: NewServicesPostgres(db),
	}
}
