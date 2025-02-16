package models

type LtUserday struct {
	Id         uint `xorm:"not null pk autoincr UNSIGNED INT"`
	Uid        uint `xorm:"not null default 0 comment('用户ID') unique(uid_day) UNSIGNED INT"`
	Day        uint `xorm:"not null default 0 comment('日期，如：20180725') unique(uid_day) UNSIGNED INT"`
	Num        uint `xorm:"not null default 0 comment('次数') UNSIGNED INT"`
	SysCreated uint `xorm:"not null default 0 comment('创建时间') UNSIGNED INT"`
	SysUpdated uint `xorm:"not null default 0 comment('修改时间') UNSIGNED INT"`
}

func (m *LtUserday) TableName() string {
	return "lt_userday"
}
