package services

import (
	"go-lottery/dao"
	"go-lottery/models"
)

type BlackipService interface {
	GetAll() []models.LtBlackip
	CountAll() int64
	Get(id uint) *models.LtBlackip
	Delete(id uint) error
	Update(data *models.LtBlackip, columns []string) error
	Create(data *models.LtBlackip) error
}

type blackipService struct {
	dao *dao.BlackipDAO
}

func NewblackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDAO(nil),
	}
}

func (s *blackipService) GetAll() []models.LtBlackip {
	return s.dao.GetAll()
}

func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *blackipService) Get(id uint) *models.LtBlackip {
	return s.dao.Get(id)
}

func (s *blackipService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *blackipService) Create(data *models.LtBlackip) error {
	return s.dao.Create(data)
}
