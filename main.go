package main

import (
	"IM/controller"
	"html/template"
	"log"
	"net/http"
)

func RegisterView() {
	// 解析符合这种文件格式的
	tpl, err := template.ParseGlob("view/**/*")
	// 如果报错就不要继续了
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// 提供静态资源目录支持
	//http.Handle("/", http.FileServer(http.Dir(".")))
	// 提供指定目录的静态文件支持  使得直接加载index.html
	RegisterView()
	// 启动服务器
	http.ListenAndServe(":8080", nil)
}
