package services

import (
	"go-lottery/dao"
	"go-lottery/models"
)

type UserService interface {
	GetAll(page, size int) []models.LtUser
	CountAll() int
	Get(id uint) *models.LtUser
	Delete(id uint) error
	Update(data *models.LtUser, columns []string) error
	Create(data *models.LtUser) error
	Search(country string) []models.LtUser
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(nil),
	}
}

func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}

func (s *userService) CountAll() int {
	return s.dao.CountAll()
}

func (s *userService) Get(id uint) *models.LtUser {
	return s.dao.Get(id)
}

func (s *userService) Delete(id uint) error {
	return s.dao.Delete(id)
}

func (s *userService) Update(data *models.LtUser, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userService) Create(data *models.LtUser) error {
	return s.dao.Create(data)
}

func (s *userService) Search(country string) []models.LtUser {
	return s.dao.Search(country)
}
