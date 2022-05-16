package controller

import (
	"IM/model"
	"IM/service"
	"IM/util"
	"fmt"
	"math/rand"
	"net/http"
)

var userService service.UserService

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW
	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOK(writer, user, "")
	}
}

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	// 数据库操作
	//逻辑处理
	//Restful/API json/XML的返回
	// 1. 获取前端传递的参数
	// mobile passwd
	// 解析参数
	// 如何获得参数
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOK(writer, user, "")
	}
}
