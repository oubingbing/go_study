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
	response,err := http.Get("https://kyfw.12306.cn/otn/leftTicket/queryA?leftTicketDTO.train_date=2019-12-20&leftTicketDTO.from_station=SZQ&leftTicketDTO.to_station=GZQ&purpose_codes=ADULT")
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

