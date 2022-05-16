package service

import (
	"IM/model"
	"IM/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

// Register 注册函数
func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	// 检查手机号是否存在
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	// 如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	// 否则拼接插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	// passwd=
	//md5 加密
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	// 插入
	_, err = DbEngin.InsertOne(&tmp)

	// 返回新用户信息
	return tmp, err
}

// Login 登录函数
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	// 首先通过手机号查询用户
	tmp := model.User{}
	DbEngin.Where("mobile=?", mobile).Get(&tmp)
	// 如果没有找到
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	// 查询到了比对密码
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	// 如果密码对了 刷新token 为了安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.Md5Encode(str)
	tmp.Token = token
	//更新数据库token
	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	// 返回数据
	return tmp, nil
}
