package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func userLogin(writer http.ResponseWriter, request *http.Request) {
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
	loginok := false
	if mobile == "18600000000" && passwd == "123456" {
		loginok = true
	}
	if loginok {
		// {"id":1, "token" : "xx"}
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "")
	} else {
		Resp(writer, -1, nil, "密码不正确")
	}
}

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	// 设置header为json  默认的为text/html  所以特别指出返回的为application/json
	writer.Header().Set("Content-Type", "application/json")
	// 设置200状态
	writer.WriteHeader(http.StatusOK)
	// 定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// 将结构体转化为JSON字符串  用json的Marshal方法
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	// 输出
	writer.Write(ret)
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)
	// 提供静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))
	// 提供指定目录的静态文件支持  使得直接加载index.html
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// 启动服务器
	http.ListenAndServe(":8080", nil)
}
