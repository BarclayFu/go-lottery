package dao

import (
	"go-lottery/models"
	"log"
	"xorm.io/xorm"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

func (d *CodeDao) Get(id uint) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll() []models.LtCode {
	dataList := make([]models.LtCode, 0)
	err := d.engine.
		Desc("id").
		Find(&dataList)
	if err != nil {
		log.Println("dao.GetAll err:", err)
		return dataList
	}
	return dataList
}

func (d *CodeDao) CountAll() int64 {
	count, err := d.engine.Count(&models.LtCode{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *CodeDao) Delete(id uint) error {
	data := &models.LtCode{Id: id, SysStatus: 1}
	_, err := d.engine.ID(data.Id).Update(data)
	return err
}

func (d *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Insert(data)
	return err
}
