package main

import (
	"net/http"
	"io"
	"log"
	"os"
)

// beego 版本
//import "github.com/astaxie/beego"
//
//type HomeController struct {
//	beego.Controller
//}
//
//func (this *HomeController) Get()  {
//	this.Ctx.WriteString("hello world\n")
//}
//func main() {
//	beego.Router("/",&HomeController{})
//	beego.Run()
//}

//func main() {
//	http.HandleFunc("/", sayHello)
//	err := http.ListenAndServe(":8080", nil)
//	if (err != nil) {
//		log.Fatal(err)
//	}
//}
//
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello workd ,this is version 1\n")
//}

//第二个版本
//func main() {
//	mux := http.NewServeMux()
//	mux.Handle("/", &myHandler{})
//	mux.HandleFunc("/hello",sayHello)
//	err := http.ListenAndServe(":8080", mux)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//type myHandler struct {
//
//}
//
//func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "URL: "+r.URL.String())
//}
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello workd ,this is version 2\n")
//}

// 第三个版本

//func main() {
//	server := http.Server{
//		Addr:        ":8088",
//		Handler:     &myHandler{},
//		ReadTimeout: 5 * time.Second,
//	}
//
//	err := server.ListenAndServe()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//type myHandler struct {
//}
//
//func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "URL: "+r.URL.String())
//}

// 第四个版本 最底层的
//var mux map[string]func(http.ResponseWriter, *http.Request)
//
//func main() {
//	server := http.Server{
//		Addr:        ":8088",
//		Handler:     &myHandler{},
//		ReadTimeout: 5 * time.Second,
//	}
//	mux = make(map[string]func(http.ResponseWriter, *http.Request))
//	mux["/hello"] = sayHello
//	mux["/bye"] = sayBye
//	err := server.ListenAndServe()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//type myHandler struct {
//}
//
//func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if h,ok := mux[r.URL.String()]; ok {
//		h(w,r)
//		return
//	}
//	io.WriteString(w, "URL: "+r.URL.String())
//
//
//}
//
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello ,say hello")
//}
//func sayBye(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello, say Bye")
//
//}

// 静态文件服务器 必须绝对路径

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say hello")

}
