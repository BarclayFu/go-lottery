package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"math/rand"
	"time"
)

var packageList map[uint32][]uint32 = make(map[uint32][]uint32)

type lotteryConotroller struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryConotroller{})

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 返回全部红包地址
// http://localhost:8080
func (c *lotteryConotroller) Get() map[uint32][2]int {
	rs := make(map[uint32][2]int)
	for id, list := range packageList {
		var money int
		for _, v := range list {
			money += int(v)
		}
		rs[id] = [2]int{len(list), money}
	}
	return rs
}

// http://localhost:8080/set?uid=1&money=100&num=100
func (c *lotteryConotroller) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamInt("money")
	num, errNum := c.Ctx.URLParamInt("num")
	if errUid != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("errUid=%d, errMonry=%d, errNum=%d\n", errUid, errMoney, errNum)
	}
	moneyTotal := int(money * 100)
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%d, num=%d", uid, money, num)
	}
	// 金额分配算法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.55 //随机分配的最大值
	list := make([]uint32, num)
	leftNum := num
	leftMoney := moneyTotal
	for leftNum > 0 {
		if leftNum == 1 {
			list[num-1] = uint32(leftMoney)
			break
		}
		if leftMoney == leftNum {
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}
		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint32(m)
		leftMoney -= m
		leftNum--
	}
	//红包的唯一ID
	id := r.Uint32()
	packageList[id] = list
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)

}
