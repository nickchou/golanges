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
	"strconv"
	"time"

	"github.com/nickchou/golanges/comm"
	"github.com/nickchou/golanges/dataaccess"
	"github.com/nickchou/golanges/entity"
)

const ES_URL string = "http://192.168.56.101:9200"

func main() {
	//ExportData() //从sqlserver数据库中导数据到elastic
	TodoUpdate() //更新数据测试
}
func TodoUpdate() {
	begin := time.Now()
	//创建一个channel，goroute执行完后塞一个值
	count := 6
	c := make(chan int, count)
	//	UpdateByID("111", strconv.Itoa(1))
	for i := 1; i <= count; i++ {
		//go UpdateByQuery("111", strconv.Itoa(i), c)
		//go UpdateByIndexScript("abc123", strconv.Itoa(i), c)
		go UpdateByIndexDoc("abc123", strconv.Itoa(i), c)
	}
	for j := 0; j < count; j++ {
		<-c //从信号量中去值出来，直到所有goroutine都执行完毕
	}
	end := time.Now()
	fmt.Println("结束时间:", end, time.Since(begin))
}

//
func UpdateByIndexDoc(id string, prex string, cc chan int) {
	strUpdate := fmt.Sprintf(`{
  "doc" : {
        "name": "lisi%s"
    }
}`, prex)
	//fmt.Printf(strUpdate)
	postdata := []byte(strUpdate)
	byteRes, err := comm.HttpPost(fmt.Sprintf("%s/stu/person/%s/_update?retry_on_conflict=5", ES_URL, id), postdata)
	if err != nil {
		log.Println("http.NewRequest,[err=%s]", err)
	}
	fmt.Println(prex + string(byteRes[:]))
	cc <- 1
}

//Too many dynamic script compilations within, max: [75/5m]; please use indexed, or scripts with parameters instead; this limit can be changed by the [script.max_compilations_rate] setting
func UpdateByIndexScript(id string, prex string, cc chan int) {
	strUpdate := fmt.Sprintf(`{
  "script" : {
        "source": "ctx._source.name = \"lisi%s\";",
        "lang": "painless"
    }
}`, prex)
	//fmt.Printf(strUpdate)
	postdata := []byte(strUpdate)
	byteRes, err := comm.HttpPost(fmt.Sprintf("%s/stu/person/%s/_update?refresh=true", ES_URL, id), postdata)
	if err != nil {
		log.Println("http.NewRequest,[err=%s]", err)
	}
	fmt.Println(prex + string(byteRes[:]))
	cc <- 1
}

//用 _update_by_query这里如果并发执行会报错
//[script] Too many dynamic script compilations within, max: [75/5m]; please use indexed, or scripts with parameters instead; this limit can be changed by the [script.max_compilations_rate] setting
func UpdateByQuery(id string, prex string, cc chan int) {
	strUpdate := fmt.Sprintf(`{
  "script": {
   "source": "ctx._source.name=\"lisi%s\";"
  },
  "query": {
  	"bool": {
    	"filter": [
      		{"term": { "id": "%s" } }
    	]
  	}
  }
}`, prex, id)
	//fmt.Printf(strUpdate)
	postdata := []byte(strUpdate)
	byteRes, err := comm.HttpPost(fmt.Sprintf("%s/stu/_update_by_query?conflicts=proceed&refresh=true&timeout=1s", ES_URL), postdata)
	if err != nil {
		log.Println("http.NewRequest,[err=%s]", err)
	}
	fmt.Println(prex + string(byteRes[:]))
	cc <- 1
}
func ExportData() {
	pageIndex := 201
	pageSize := 1000
	pageCount := 500
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
