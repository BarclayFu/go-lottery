package services

import (
	"go-lottery/dao"
	"go-lottery/models"
)

type ResultService interface {
	GetAll(page, size int) []models.LtResult
	CountAll() int64
	Get(id uint) *models.LtResult
	Delete(id uint) error
	Update(data *models.LtResult, columns []string) error
	Create(data *models.LtResult) error
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(nil),
	}
}

func (s *resultService) GetAll(page, size int) []models.LtResult {
	return s.dao.GetAll(page, size)
}

func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *resultService) Get(id uint) *models.LtResult {
	return s.dao.Get(id)
}

func (s *resultService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *resultService) Update(data *models.LtResult, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *resultService) Create(data *models.LtResult) error {
	return s.dao.Create(data)
}
