package main

import (
	"bytes"
	"fmt"
	"io"
)

func main()  {
	//compare()
	//container()
	equal()
 }

func equal()  {
	a := "hello world"
	b := "hello world"

	result := bytes.Equal([]byte(b),[]byte(a))
	fmt.Println(result)//true
}

//判断a在b中字符串出现的次数
func count()  {
	a := "o"
	b := "hello world"

	result := bytes.Count([]byte(b),[]byte(a))
	fmt.Println(result)//2
}

//判断b是否包含a
func container()  {
	a := "o"
	b := "hello world"
	result := bytes.Contains([]byte(b),[]byte(a))
	fmt.Println(result)//true
}

//比较两个字符串序列，如果两个字符串不一样就返回1，一致则返回0
func compare()  {
	var a string = "123"
	var b string = "你好呀"

	result := bytes.Compare([]byte(a),[]byte(b))
	fmt.Println(result)//1
}

func t()  {
	var reader io.Reader
	str := []byte("test")
	fmt.Println(reader.Read(str))
}
