package service

import (
	"avitoTech"
	"avitoTech/pkg/repository"
)

type ServicesService struct {
	repo repository.Services
}

func NewServicesService(repo repository.Services) *ServicesService {
	return &ServicesService{repo: repo}
}

func (s *ServicesService) GetReportByServ() (posts []avitoTech.Report, err error) {
	return s.repo.GetReportByServ()
}
