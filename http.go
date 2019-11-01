package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	CurlGet()
}

func CurlGet()  {
	var response *http.Response
	response,err := http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Printf("错误 %v\n",err.Error())
	}

	defer response.Body.Close()
	result,readError := ioutil.ReadAll(response.Body)
	if readError != nil{
		fmt.Printf("读取错误 %v\n",readError.Error())
	}

	fmt.Printf("结果：%v\n",string(result))
}
