package main

import (
	"bytes"
	"fmt"
)

func main()  {
	readSingle()
}

func testBuffer()  {
	var buffer1 bytes.Buffer

	contents := "Simple byte buffer for marshaling data."

	fmt.Printf("Writing contents %q ...\n", contents)

	buffer1.WriteString(contents)

	fmt.Printf("The length of buffer: %d\n", buffer1.Len())

	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())

	p1 := make([]byte, 7)

	n, _ := buffer1.Read(p1)

	p2 := make([]byte,7)
	buffer1.Read(p2)

	fmt.Printf("%d bytes were read. (call Read)\n", n)

	fmt.Printf("The length of buffer: %d\n", buffer1.Len())

	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())

	fmt.Printf("读取的内容: %v\n", string(p2))
}

func readSingle()  {
	str := "hello world"

	for _,s := range str{
		fmt.Println(string(s))
	}
}


