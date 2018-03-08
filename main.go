package main

/*
功能点
1、SQL Server查询
2、Elasticsearch的数据插入
3、es查询性能测试
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nickchou/golanges/comm"
	"github.com/nickchou/golanges/dataaccess"
	"github.com/nickchou/golanges/entity"
)

const ES_URL string = "http://192.168.56.101:9200"

func main() {
	pageIndex := 1
	pageSize := 1000
	pageCount := 200
	for pageIndex <= pageCount {
		var users []*entity.User = dataaccess.GetUser(pageIndex, pageSize)
		for _, u := range users {
			//fmt.Print(u.Id, "\t", u.Name, "\t", u.Gender)
			//fmt.Println()
			u.City = "PEK"
			PostElasticSearch(u)
		}
		fmt.Println("page:", pageIndex, pageCount, time.Now())
		pageIndex++
	}
	fmt.Printf("end")
	//json格式化
	//json, _ := json.Marshal(users)
	//fmt.Println(string(json))
}

//新增数据到elasticsearch保存
func PostElasticSearch(u *entity.User) {
	//json
	byteReq, _ := json.Marshal(u)
	//fmt打印请求参数
	//fmt.Println(string(byteReq))
	//response URL+索引+Type 新增
	byteRes, err := comm.HttpPost(fmt.Sprintf("%s/user/person", ES_URL), byteReq)
	if err != nil {
		log.Println("http.NewRequest,[err=%s]", err)
	}
	var EsRes entity.EsResponse
	//rejson
	err = json.Unmarshal(byteRes, &EsRes)
	if EsRes.Status == 0 {
		//fmt.Println("es success,indexID:", EsRes.Id)
	} else {
		fmt.Println("Error.Type:", EsRes.Error.Type)
		fmt.Println("Error.Reason:", EsRes.Error.Reason)
	}
	//fmt.Println(string(byteRes))
	//
}
