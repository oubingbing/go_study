package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"sync"
	"time"
)

type Result struct {
	Day time.Time
	Platform string
	Media string
	Agent string
	AccountName string
}

func main()  {

	chn := make(chan int,10)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1;i<=10; i++ {
		go getData(chn,wg,i)
	}

	go wait(chn,&wg)

	for v := range chn  {
		fmt.Println(v)
	}

	fmt.Println("结束")

	//getData2(1)

}

/**
 * 等待关闭渠道
 */
func wait(chn chan int,wg *sync.WaitGroup)  {
	wg.Wait()
	close(chn)
}

func getData2( pageNumber int)  {
	db, err := gorm.Open("mssql", "sqlserver://bi_china:FYv6y3O+UVtG@129.211.84.25:31051?database=Work_CN")
	if err != nil {
		fmt.Printf("连接错误：%v\n",err.Error())
	}
	defer db.Close()

	//var result Result
	rows,selectErrorr := db.Raw("select * from [stonemmocn].[bi_report_media_pid_creative_day_info_whole] as [t] left join [stonemmocn].[bi_relation_account_agent_discount] as [a] on [a].[start_date] <= [t].[day] and [a].[end_date] >= [t].[day] and [a].[account] = [t].[account_name] and [a].[media] = [t].[media] and [a].[agent] = [t].[agent] where [t].[day] >= ? and [t].[day] <= ? order by t.day OFFSET ? ROWS FETCH NEXT 1000 ROWS ONLY","2019-11-01","2019-11-02",pageNumber).Rows()
	if selectErrorr != nil{
		fmt.Printf("查询错误：%v\n",err.Error())
	}

	col,_ := rows.Columns()

	sli := make([][]byte,len(col))
	args := make([]interface{},len(col))
	fields := make(map[string]interface{})
	var result []map[string]interface{}
	for i,field := range col {
		args[i] = &sli[i]
		fields[field] = nil
	}

	for rows.Next() {
		e := rows.Scan(args...)
		if e != nil {
			fmt.Println(e)
		}

		for i,v := range col {
			fields[v] = string(sli[i])
		}

		result = append(result,fields)

	}
}

func getData( chn chan int,wg sync.WaitGroup,pageNumber int)  {
	db, err := gorm.Open("mssql", "")
	if err != nil {
		fmt.Printf("连接错误：%v\n",err.Error())
	}
	defer db.Close()

	//var result Result
	rows,selectErrorr := db.Raw("select * from [stonemmocn].[bi_report_media_pid_creative_day_info_whole] as [t] left join [stonemmocn].[bi_relation_account_agent_discount] as [a] on [a].[start_date] <= [t].[day] and [a].[end_date] >= [t].[day] and [a].[account] = [t].[account_name] and [a].[media] = [t].[media] and [a].[agent] = [t].[agent] where [t].[day] >= ? and [t].[day] <= ? order by t.day OFFSET ? ROWS FETCH NEXT 1000 ROWS ONLY","2019-11-01","2019-11-02",pageNumber).Rows()
	if selectErrorr != nil{
		fmt.Printf("查询错误：%v\n",err.Error())
	}

	col,_ := rows.Columns()

	sli := make([][]byte,len(col))
	args := make([]interface{},len(col))
	fields := make(map[string]interface{})
	var result []map[string]interface{}
	for i,field := range col {
		args[i] = &sli[i]
		fields[field] = nil
	}

	for rows.Next() {
		e := rows.Scan(args...)
		if e != nil {
			fmt.Println(e)
		}

		for i,v := range col {
			fields[v] = string(sli[i])
		}

		result = append(result,fields)

	}

	chn <- pageNumber
	wg.Done()
}

func sql1()  {
	db, err := gorm.Open("mssql", "")
	if err != nil {
		fmt.Printf("连接错误：%v\n",err.Error())
	}
	defer db.Close()

	//var result Result
	rows,selectErrorr := db.Raw("select * from stonemmocn.bi_report_media_pid_account_day_info_whole where account_name = ?","baidu-石器-B19KA03627").Rows()
	if selectErrorr != nil{
		fmt.Printf("查询错误：%v\n",err.Error())
	}

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
}
