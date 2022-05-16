package util

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}
func RespOK(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}
