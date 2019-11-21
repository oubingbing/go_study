package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main()  {
	db,err := conn()
	if err != nil{
		fmt.Println(err)
	}

	rows,rowErr := db.Query("select * from worlds")
	if rowErr != nil{
		fmt.Printf("获取行错误：%v\n",rowErr.Error())
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([][]byte, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	var dataRow []interface{}
	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, v := range values {
			if v != nil {
				//fmt.Println(reflect.TypeOf(col))
				record[columns[i]] = string(v)
			}
		}
		dataRow = append(dataRow,record)
	}

	fmt.Println(dataRow)

	/*cols,_ := rows.Columns()
	dataRow := make([]interface{},0)
	for rows.Next() {
		//将行数据保存到record字典
		values := make([][]byte, len(cols))
		fields := make([]interface{}, len(cols))
		for index,_ := range cols {
			fields[index] = &values[index]
		}

		err := rows.Scan(fields...)
		if err != nil{
			fmt.Printf("获取字段值错误：%v\n",err.Error())
		}

		fieldMap := make(map[string]string)
		for i,v := range cols  {
			fieldMap[v] = string(values[i])
		}

		dataRow = append(dataRow,fieldMap)
	}*/

}

func conn() (*sql.DB,error) {
	var db *sql.DB
	var err error
	db,err = sql.Open("mysql", "root:bingbing@tcp(127.0.0.1:3306)/bi_stoneage?charset=utf8")
	return db,err
}