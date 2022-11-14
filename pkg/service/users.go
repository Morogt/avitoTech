package service

import (
	"avitoTech"
	"avitoTech/pkg/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(user avitoTech.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UsersService) GetUsers(id int) (avitoTech.User, error) {
	return s.repo.GetUser(id)
}
