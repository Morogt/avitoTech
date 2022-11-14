package service

import (
	"avitoTech"
	"avitoTech/pkg/repository"
)

type Balance interface {
	GetInfo(userId int) (user avitoTech.User, err error)
	AddMoney(userId int, count float64) (user avitoTech.User, err error)
	WithdrawalOfMoney(userId int, count float64) (user avitoTech.User, err error)
	Transfer(senderId, recipientId int, count float64) (senderLat avitoTech.User, recipientLat avitoTech.User, err error)
	Pay(userId, serviceId int, count float64, description string) (user avitoTech.User, orderId int, err error)
	History(id int) (history []avitoTech.History, err error)
}

type Users interface {
	CreateUser(user avitoTech.User) (int, error)
}

type Services interface {
	//CreateServ(serv avitoTech.Service) (int, error)
	GetReportByServ() (posts []avitoTech.Report, err error)
}

type Service struct {
	Balance
	Users
	Services
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Users:    NewUsersService(repos.Users),
		Balance:  NewBalanceService(repos.Balance),
		Services: NewServicesService(repos.Services),
	}
}
