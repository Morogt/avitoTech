package service

import (
	"avitoTech"
	"avitoTech/pkg/repository"
	"errors"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetInfo(userId int) (user avitoTech.User, err error) {
	return s.repo.GetInfo(userId)
}

func (s *BalanceService) AddMoney(userId int, count float64) (user avitoTech.User, err error) {
	info, err := s.repo.GetInfo(userId)
	if err != nil {
		return user, err
	}
	lastBalance := info.Balance
	return s.repo.AddMoney(userId, lastBalance+count, count)
}

func (s *BalanceService) WithdrawalOfMoney(userId int, count float64) (user avitoTech.User, err error) {
	info, err := s.repo.GetInfo(userId)
	if err != nil {
		return user, err
	}
	lastBalance := info.Balance
	if count > lastBalance {
		return user, errors.New("count should be less or equal to the balance")
	}
	return s.repo.WithdrawalOfMoney(userId, lastBalance-count, count)
}

func (s *BalanceService) Transfer(senderId, recipientId int, count float64) (senderLat avitoTech.User, recipientLat avitoTech.User, err error) {
	sender, err := s.repo.GetInfo(senderId)
	if err != nil {
		return senderLat, recipientLat, err
	}
	recipient, err := s.repo.GetInfo(recipientId)
	if err != nil {
		return senderLat, recipientLat, err
	}
	if sender.Balance < count {
		return senderLat, recipientLat, errors.New("count should be less or equal to the sender balance")
	}
	return s.repo.Transfer(senderId, recipientId, sender.Balance-count, recipient.Balance+count, count)
}

func (s *BalanceService) Pay(userId, serviceId int, count float64, description string) (user avitoTech.User, orderId int, err error) {
	info, err := s.repo.GetInfo(userId)
	if err != nil {
		return user, 0, err
	}
	if info.Balance < count {
		return user, 0, errors.New("count should be less or equal to the balance")
	}

	return s.repo.Pay(userId, serviceId, info.Balance-count, count, description)
}

func (s *BalanceService) History(id int) (history []avitoTech.History, err error) {
	return s.repo.History(id)
}
