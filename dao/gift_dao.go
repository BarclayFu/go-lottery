package dao

import (
	"go-lottery/models"
	"xorm.io/xorm"
)

type GiftDAO struct {
	engine *xorm.Engine
}

func NewGiftDAO(engine *xorm.Engine) *GiftDAO {
	return &GiftDAO{
		engine: engine,
	}
}

func (d *GiftDAO) Get(id uint) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *GiftDAO) GetAll() []models.LtGift {
	dataList := make([]models.LtGift, 0)
	err := d.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&dataList)
}
