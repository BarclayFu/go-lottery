package dao

import (
	"go-lottery/models"
	"log"
	"xorm.io/xorm"
)

type BlackipDAO struct {
	engine *xorm.Engine
}

func NewBlackipDAO(engine *xorm.Engine) *BlackipDAO {
	return &BlackipDAO{
		engine: engine,
	}
}

func (d *BlackipDAO) Get(id uint) *models.LtBlackip {
	data := &models.LtBlackip{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *BlackipDAO) GetAll() []models.LtBlackip {
	dataList := make([]models.LtBlackip, 0)
	err := d.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&dataList)
	if err != nil {
		log.Println("dao.GetAll err:", err)
		return dataList
	}
	return dataList
}

func (d *BlackipDAO) CountAll() int64 {
	count, err := d.engine.Count(&models.LtBlackip{})
	if err != nil {
		return 0
	} else {
		return count
	}
}

func (d *BlackipDAO) Delete(id uint) error {
	data := &models.LtBlackip{Id: id}
	_, err := d.engine.ID(data.Id).Update(data)
	return err
}

func (d *BlackipDAO) Update(data *models.LtBlackip, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackipDAO) Create(data *models.LtBlackip) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackipDAO) GetByIp(ip string) *models.LtBlackip {
	dataList := make([]models.LtBlackip, 0)
	err := d.engine.Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
