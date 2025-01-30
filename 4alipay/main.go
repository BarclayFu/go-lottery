package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type gift struct {
	id      int    //奖品ID
	name    string // 奖品名称
	pic     string // 奖品的图片
	link    string // 奖品的链接
	inuse   bool   //是否使用中
	rate    int    //中奖概率，万分之N， 0-9999
	rateMin int    // 大雨等于最早中奖编码
	rateMax int    //小于中奖编码
}

// 最大的中奖号码
const rateMax = 10000

var logger *log.Logger

type lotteryController struct {
	Ctx iris.Context
}

func initLog() {

	f, err := os.Create("/Users/fusizhe/Documents/GitHub/go-lottery/3wechatShake/demo.log")
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newGift() *[5]gift {
	giftList := new([5]gift)
	g1 := gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[0] = g1
	g2 := gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[1] = g2
	g3 := gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[2] = g3
	g4 := gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[3] = g4
	g5 := gift{
		id:      5,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "",
		inuse:   false,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[4] = g5
	return giftList
}

func giftRage(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)
	rateStart := 0
	for i, data := range giftList {
		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate
		if giftList[i].rateMax > rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftList=%v\n", giftList)
	return giftList
}

// GET http://localhost:8080/?rate=4,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRage(rate)
	return fmt.Sprintf("%v", giftList)
}

func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	giftList := giftRage(rate)
	for _, data := range giftList {
		if !data.inuse {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			// 中奖了，抽奖编码在奖品编码范围内
			// 开始发奖
			sendData := data.pic
			ok = true
			if ok {
				//中奖后，成功得到奖品
				// 生成中奖记录
				saveLuckyData(code, data.id, data.name, data.link, sendData)
				result["success"] = ok
				result["uid"] = uid
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				break
			}

		}
	}
	return result
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

// 记录用户的获奖信息
func saveLuckyData(code int32, id int, name string, link string, data string) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s\n",
		code, id, name, link, data)
}
