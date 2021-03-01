package main

import "fmt"

func main() {
	/*data := []byte("test")
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)
	fmt.Println(md5str)*/

	sli := make([]string,1,2)

	sli = append(sli,"test")
	sli = append(sli,"test2")
	sli = append(sli,"test3")

	fmt.Println(sli)

}
