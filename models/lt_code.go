package models

type LtCode struct {
	Id         uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	GiftId     uint   `xorm:"not null default 0 comment('奖品ID，关联lt_gift表') index UNSIGNED INT"`
	Code       string `xorm:"not null default '' comment('虚拟券编码') unique VARCHAR(255)"`
	SysCreated uint   `xorm:"not null default 0 comment('创建时间') UNSIGNED INT"`
	SysUpdated uint   `xorm:"not null default 0 comment('更新时间') UNSIGNED INT"`
	SysStatus  uint   `xorm:"not null default 0 comment('状态，0正常，1作废，2已发放') UNSIGNED SMALLINT"`
}

func (m *LtCode) TableName() string {
	return "lt_code"
}
