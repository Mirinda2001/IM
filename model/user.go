package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MAN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	// 用户的ID
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	// 用户的手机
	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	// 用户的密码=f(plainpwd + salt)
	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	// 随机数 加密因子  每个用户的salt不一样
	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`
	Online int    `xorm:"int(10)" form:"online" json:"online"`
	// 用户鉴权使用 chat?id=1&token=x
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
