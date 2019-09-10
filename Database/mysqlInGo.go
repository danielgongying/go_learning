
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
/*选择：select * from table1 where 范围
插入：insert into table1(field1,field2) values(value1,value2)
删除：delete from table1 where 范围
更新：update table1 set field1=value1 where 范围
查找：select * from table1 where field1 like ’%value1%’ ---like的语法很精妙，查资料!
排序：select * from table1 order by field1,field2 [desc]
总数：select count as totalcount from table1
求和：select sum(field1) as sumvalue from table1
平均：select avg(field1) as avgvalue from table1
最大：select max(field1) as maxvalue from table1
最小：select min(field1) as minvalue from table1*/
func insertValues(db *sql.DB ,value string)  {
	insert,err:=db.Query("INSERT INTO users VALUES (value)")
	if err != nil {
		panic(err.Error())

	}
	defer  insert.Close()
	fmt.Print("insert succeed")
}

func main() {
	fmt.Println("connect to mysql-daniel")
	db,err:=sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/mysql")//root:password@tcp(127.0.0.1:3306)/name
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected now")
	defer  db.Close()
	user,_ := db.Query("create table user('name','age') values ('daniel',12)")
	fmt.Print(user)
	fmt.Print("create newtable succeed")
	//insert,err:=db.Query("insert into user values('daniel')")
	//if err != nil {
	//	panic(err.Error())
	//
	//}
	//defer  insert.Close()
	//fmt.Print("insert succeed")

}
