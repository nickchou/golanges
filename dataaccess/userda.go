package dataaccess

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/nickchou/golanges/entity"
)

const connStr string = "server=127.0.0.1;port=1433;database=shifenzheng;user id=sa;password=123"

func GetUser(index int, size int) []*entity.User {
	if index < 1 {
		index = 1
	}
	var users []*entity.User
	//创建连接
	conn, err := sql.Open("mssql", connStr)
	defer conn.Close()
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	//拼接SQL
	strSQL := fmt.Sprintf(`
	SELECT id,Name,CtfId,Gender,Birthday,Address,Zip,Mobile,Tel,Fax,EMail 
	FROM [shifenzheng].[dbo].[cdsgus]
	ORDER BY id OFFSET %d ROW  FETCH NEXT %d ROW ONLY`, (index-1)*size, size)

	//打印SQL
	//fmt.Println(strSQL)
	//stmt
	stmt, _ := conn.Prepare(strSQL)
	defer stmt.Close()
	//查看结果
	rows, _ := stmt.Query()
	defer rows.Close()
	//从rows中获取数据
	for rows.Next() {
		var row = new(entity.User)
		rows.Scan(&row.Id, &row.Name, &row.Ctfid, &row.Gender, &row.Birthday, &row.Address, &row.Zip, &row.Mobile, &row.Tel, &row.Fax, &row.Email)
		users = append(users, row)
	}
	return users
}
