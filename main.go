package main

import (
	"fmt"
	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
	"html/template"
	"net/http"
)

func main(){

	app := tview.NewApplication()
	wm := winman.NewWindowManager()

	content := tview.NewTextView().
		SetText("Hello, world!").       // set content of the text view
		SetTextAlign(tview.AlignCenter) // align text to the center of the text view

	window := wm.NewWindow(). // create new window and add it to the window manager
		Show().                   // make window visible
		SetRoot(content).         // have the text view above be the content of the window
		SetDraggable(true).       // make window draggable around the screen
		SetResizable(true).       // make the window resizable
		SetTitle("Hi!").          // set the window title
		AddButton(&winman.Button{ // create a button with an X to close the application
			Symbol:  'X',
			OnClick: func() { app.Stop() }, // close the application
		})

	window.SetRect(5, 5, 30, 10) // place the window

	// now, execute the application:
	if err := app.SetRoot(wm, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	/*http.HandleFunc("/",handle)
	http.HandleFunc("/api/user",ApiHandle)
	http.HandleFunc("/tpl", vendorHtml)
	http.HandleFunc("/index",loadTpl)

	serveError := http.ListenAndServe(":8000",nil)
	if serveError != nil{
		fmt.Println(serveError.Error())
	}*/
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