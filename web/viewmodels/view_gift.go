package viewmodels

type ViewGift struct {
	Id           uint   `form:"id"`
	Title        string `form:"title"`
	PrizeNum     int    `form:"prize_num"`
	PrizeCode    string `form:"prize_code"`
	PrizeTime    uint   `form:"prize_time"`
	Img          string `form:"img"`
	Displayorder uint   `form:"displayorder"`
	Gtype        uint   `form:"gtype"`
	Gdata        string `form:"gdata"`
	TimeBegin    string `form:"time_begin"`
	TimeEnd      string `form:"time_end"`
}
