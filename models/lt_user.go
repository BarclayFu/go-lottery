package models

type LtUser struct {
	Id         uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Blacktime  uint   `xorm:"not null default 0 comment('黑名单限制到期时间') UNSIGNED INT"`
	Realname   string `xorm:"not null default '' comment('联系人') VARCHAR(50)"`
	Mobile     string `xorm:"not null default '' comment('手机号') VARCHAR(50)"`
	Address    string `xorm:"not null default '' comment('联系地址') VARCHAR(255)"`
	SysCreated uint   `xorm:"not null default 0 comment('创建时间') UNSIGNED INT"`
	SysUpdated uint   `xorm:"not null default 0 comment('修改时间') UNSIGNED INT"`
	SysIp      string `xorm:"not null default '' comment('IP地址') VARCHAR(50)"`
}

func (m *LtUser) TableName() string {
	return "lt_user"
}
