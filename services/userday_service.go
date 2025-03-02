package services

import (
	"go-lottery/dao"
	"go-lottery/models"
)

type UserdayService interface {
	Get(id uint) *models.LtUserday
	GetAll(page, size int) []models.LtUserday
	CountAll() int64
	Search(uid, day int) []models.LtUserday
	Count(uid, day uint) uint
	Delete(id uint) error
	Update(data *models.LtUserday, columns []string) error
	Create(data *models.LtUserday) error
}

type userdayService struct {
	dao *dao.UserdayDao
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao: dao.NewUserdayDao(nil),
	}
}

func (s *userdayService) Get(id uint) *models.LtUserday {
	return s.dao.Get(id)
}

func (s *userdayService) GetAll(page, size int) []models.LtUserday {
	return s.dao.GetAll(page, size)
}

func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *userdayService) Search(uid, day int) []models.LtUserday {
	return s.dao.Search(uid, day)
}

func (s *userdayService) Count(uid, day uint) uint {
	return s.dao.Count(uid, day)
}

func (s *userdayService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *userdayService) Update(data *models.LtUserday, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userdayService) Create(data *models.LtUserday) error {
	return s.dao.Create(data)
}
