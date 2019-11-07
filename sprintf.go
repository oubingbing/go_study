package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	//fmt.Printf("p是啥 %v\n",p)
	return p, err
}

func sampleReadFromString() {
	var p *byte

	// 从字符串读取
	data, _ := ReadFrom(strings.NewReader("from string"), 5)

	p = &data[1]

	fmt.Println(p)

	fmt.Printf("指针地址 %v\n", &data[1])
	fmt.Printf("指针地址 %v\n", &p)
}

func sampleReadStdin() {

	fmt.Println("please input from stdin:")

	// 从标准输入读取
	data, _ := ReadFrom(os.Stdin, 11)

	fmt.Println(data)
}

func sampleReadFile() {

	file, _ := os.Open("main.go")
	defer file.Close()

	// 从普通文件读取，其中file是os.File的实例
	data, _ := ReadFrom(file, 9)

	fmt.Println(data)
}

func main() {

	//ampleReadFromString()

	//sampleReadStdin()

	//sampleReadFile()

	reader := bufio.NewWriter(os.Stdout)
	fmt.Fprint(reader,"测试")
	reader.Flush()
}
