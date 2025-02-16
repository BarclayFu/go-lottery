package models

type LtGift struct {
	Id           uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	Title        string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	PrizeNum     int    `xorm:"not null default -1 comment('奖品数量，0 无限量，>0限量，<0无奖品') INT"`
	LeftNum      int    `xorm:"not null default 0 comment('剩余数量') INT"`
	PrizeCode    string `xorm:"not null default '' comment('0-9999表示100%，0-0表示万分之一的中奖概率') VARCHAR(50)"`
	PrizeTime    uint   `xorm:"not null default 0 comment('发奖周期，D天') UNSIGNED INT"`
	Img          string `xorm:"not null default '' comment('奖品图片') VARCHAR(255)"`
	Displayorder uint   `xorm:"not null default 0 comment('位置序号，小的排在前面') UNSIGNED INT"`
	Gtype        uint   `xorm:"not null default 0 comment('奖品类型，0 虚拟币，1 虚拟券，2 实物-小奖，3 实物-大奖') UNSIGNED INT"`
	Gdata        string `xorm:"not null default '' comment('扩展数据，如：虚拟币数量') VARCHAR(255)"`
	TimeBegin    int    `xorm:"not null default 0 comment('开始时间') INT"`
	TimeEnd      int    `xorm:"not null default 0 comment('结束时间') INT"`
	PrizeData    string `xorm:"comment('发奖计划，[[时间1,数量1],[时间2,数量2]]') MEDIUMTEXT"`
	PrizeBegin   int    `xorm:"not null default 0 comment('发奖计划周期的开始') INT"`
	PrizeEnd     int    `xorm:"not null default 0 comment('发奖计划周期的结束') INT"`
	SysStatus    uint   `xorm:"not null default 0 comment('状态，0 正常，1 删除') UNSIGNED SMALLINT"`
	SysCreated   uint   `xorm:"not null default 0 comment('创建时间') UNSIGNED INT"`
	SysUpdated   uint   `xorm:"not null default 0 comment('修改时间') UNSIGNED INT"`
	SysIp        string `xorm:"not null default '' comment('操作人IP') VARCHAR(50)"`
}

func (m *LtGift) TableName() string {
	return "lt_gift"
}
