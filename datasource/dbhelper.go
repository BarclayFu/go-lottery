package datasource

import (
	"fmt"
	"go-lottery/conf"
	"log"
	"sync"
	"xorm.io/xorm"
)

var dbLock sync.Mutex
var masterInstance *xorm.Engine

func InstanceMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}

func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database,
	)

	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("dbhelper.NewDbMaster NewEngine error", err)
		return nil
	}
	instance.ShowSQL(true)

	masterInstance = instance
	return instance

}
