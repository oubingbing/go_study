package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	var db *sql.DB

	var err error
	db,err = sql.Open("mysql","root:bingbing@tcp(127.0.0.1:3306)/aden_bi")
	if err != nil{
		panic("数据连接失败:"+err.Error())
	}

	row := db.QueryRow("select name from user where id = ?",11)

	var name *string

	scanError := row.Scan(&name)
	if scanError != nil{
		fmt.Printf("匹配数据错误 %v\n",scanError.Error())
	}

	fmt.Printf("数据 %v \n",*name)

	defer db.Close()
}

