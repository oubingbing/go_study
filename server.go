package main

import (
	"net/http"
)

func main()  {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Controller{}.user(writer,request)
	})

	http.HandleFunc("/test", func(response http.ResponseWriter,r *http.Request) {
		response.Write([]byte("测试路由"));
	})


	http.ListenAndServe("localhost:8080", nil)

}

type Controller struct {

}

func (c Controller) user(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("测试"))
}
