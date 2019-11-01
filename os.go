package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	appendString()
	read()
}

func appendString()  {
	f,err := os.OpenFile("test.log",os.O_APPEND,777)
	if err != nil{
		fmt.Printf("%v\n",err)
	}

	f.WriteString("apend content \n")
}

func read(){

	data,err := ioutil.ReadFile("test.log")
	if err != nil{
		fmt.Printf("错误 %v\n",err.Error())
	}

	fmt.Printf("data %v\n",string(data))

}

func write(){
	f,e := os.Create("test.log")
	if e != nil {
		fmt.Printf("错误 %v \n",e.Error())
	}

	result,wError := f.WriteString("hello world")
	if wError != nil{
		fmt.Printf("%v\n",wError.Error())
	}

	fmt.Printf("%v\n",result)
}
