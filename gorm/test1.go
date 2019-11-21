package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := gorm.Open("mysql", "root:bingbing@tcp(127.0.0.1:3306)/bi_stoneage?charset=utf8")
	defer db.Close()
	if err != nil{
		fmt.Printf("链接数据库出错:%v\n",err.Error())
	}

	type Worlds struct {
		Id int `gorm:"column:id"`
		Name string `gorm:"column:name"`
		World string `gorm:"column:i_world"`
	}
	var world Worlds
	db.First(&world)

	fmt.Println(world)
}
