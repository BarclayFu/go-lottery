package services

import (
	"go-lottery/dao"
	"go-lottery/models"
)

type CodeService interface {
	GetAll() []models.LtCode
	CountAll() int64
	Get(id uint) *models.LtCode
	Delete(id uint) error
	Update(data *models.LtCode, columns []string) error
	Create(data *models.LtCode) error
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(nil),
	}
}

func (s *codeService) GetAll() []models.LtCode {
	return s.dao.GetAll()
}

func (s *codeService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *codeService) Get(id uint) *models.LtCode {
	return s.dao.Get(id)
}

func (s *codeService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *codeService) Update(data *models.LtCode, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *codeService) Create(data *models.LtCode) error {
	return s.dao.Create(data)
}
