package services

import (
	"go-lottery/dao"
	"go-lottery/models"
)

type GiftService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id uint) *models.LtGift
	Delete(id uint) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
}

type giftService struct {
	dao *dao.GiftDAO
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDAO(nil),
	}
}

func (s *giftService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}

func (s *giftService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *giftService) Get(id uint) *models.LtGift {
	return s.dao.Get(id)
}

func (s *giftService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *giftService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}
