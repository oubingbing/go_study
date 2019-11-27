package main

import "fmt"

func main()  {
	slic := make([]map[string]interface{},0)
	mp := make(map[string]interface{})
	mp["name"] = "冰冰"
	mp["age"] = 18
	mp["city"] = "深圳"
	slic = append(slic,mp)

	fmt.Println(slic)
	m:= slic[0]
	fmt.Println(m["name"])
}
