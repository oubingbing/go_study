package main

import (
	"fmt"
	"os"
)

func main()  {
	//获取环境变量
	env := os.Getenv("GOENV")
	os.Setenv("GOTEST","test")
	test := os.Getenv("GOTEST")
	fmt.Printf("go环境变量: %v\n",env)
}
