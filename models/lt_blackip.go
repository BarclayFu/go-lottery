package models

type LtBlackip struct {
	Id         uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	Ip         string `xorm:"not null default '' comment('IP地址') unique VARCHAR(50)"`
	Blacktime  uint   `xorm:"not null default 0 comment('黑名单限制到期时间') UNSIGNED INT"`
	SysCreated uint   `xorm:"not null default 0 comment('创建时间') UNSIGNED INT"`
	SysUpdated uint   `xorm:"not null default 0 comment('修改时间') UNSIGNED INT"`
}

func (m *LtBlackip) TableName() string {
	return "lt_blackip"
}
