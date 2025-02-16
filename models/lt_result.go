package models

type LtResult struct {
	Id         uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	GiftId     uint   `xorm:"not null default 0 comment('奖品ID，关联lt_gift表') index UNSIGNED INT"`
	GiftName   string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	GiftType   uint   `xorm:"not null default 0 comment('奖品类型，同lt_gift. gtype') UNSIGNED INT"`
	Uid        uint   `xorm:"not null default 0 comment('用户ID') index UNSIGNED INT"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	PrizeCode  uint   `xorm:"not null default 0 comment('抽奖编号（4位的随机数）') UNSIGNED INT"`
	GiftData   string `xorm:"not null default '' comment('获奖信息') VARCHAR(255)"`
	SysCreated uint   `xorm:"not null default 0 comment('创建时间') UNSIGNED INT"`
	SysIp      string `xorm:"not null default '' comment('用户抽奖的IP') VARCHAR(50)"`
	SysStatus  uint   `xorm:"not null default 0 comment('状态，0 正常，1删除，2作弊') UNSIGNED SMALLINT"`
}

func (m *LtResult) TableName() string {
	return "lt_result"
}
