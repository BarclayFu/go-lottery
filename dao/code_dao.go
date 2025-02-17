package dao

import (
	"go-lottery/models"
	"log"
	"xorm.io/xorm"
)

type CodeDAO struct {
	engine *xorm.Engine
}

func NewCodeDAO(engine *xorm.Engine) *CodeDAO {
	return &CodeDAO{
		engine: engine,
	}
}

func (d *CodeDAO) Get(id uint) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDAO) GetAll() []models.LtGift {
	dataList := make([]models.LtGift, 0)
	err := d.engine.
		Desc("id").
		Find(&dataList)
	if err != nil {
		log.Println("dao.GetAll err:", err)
		return dataList
	}
	return dataList
}

func (d *CodeDAO) CountAll() int64 {
	count, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *CodeDAO) Delete(id uint) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.ID(data.Id).Update(data)
	return err
}

func (d *CodeDAO) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDAO) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}
