package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main(){

	http.HandleFunc("/",handle)
	http.HandleFunc("/api/user",ApiHandle)
	http.HandleFunc("/tpl", vendorHtml)
	http.HandleFunc("/index",loadTpl)

	serveError := http.ListenAndServe(":8000",nil)
	if serveError != nil{
		fmt.Println(serveError.Error())
	}
}

/**
获取路由参数
 */
func ApiHandle(w http.ResponseWriter,r *http.Request)  {
	_,err := w.Write([]byte("api 路由"))
	if err != nil{
		fmt.Println(err.Error())
	}

	name := r.URL.Query()
	fmt.Println(name["name"])
}

func handle(w http.ResponseWriter,r *http.Request)  {
	_,err := w.Write([]byte("测试"))
	if err != nil{
		fmt.Println(err.Error())
	}
}

/**
渲染html模板
 */
func vendorHtml(w http.ResponseWriter,r *http.Request){
	_,err := fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		"标题", "标题", "内容")
	if err != nil{
		fmt.Println(err.Error())
	}
}

type Page struct {
	Title string
}

func loadTpl(w http.ResponseWriter,r *http.Request)  {
	tpl,err := template.ParseFiles("index.html")
	if err != nil{
		fmt.Println(err.Error())
	}

	page := &Page{Title:"叶子"}

	executeError := tpl.Execute(w,page)
	if executeError != nil {
		fmt.Println(executeError)
	}

}