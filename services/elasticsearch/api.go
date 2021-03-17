package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"reflect"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

//初始化
func init() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

/*下面是简单的CURD*/

//创建
func create() {

	//使用结构体
	e1 := Employee{"zhao", "jun", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("db1").
		Type("table1").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"wang","last_name":"er","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("db1").
		Type("table1").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"zhao","last_name":"yike","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("db1").
		Type("table1").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

//删除
func delete() {

	res, err := client.Delete().Index("db1").
		Type("table1").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func update() {
	res, err := client.Update().
		Index("db1").
		Type("table1").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

//查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("db1").Type("table1").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		//fmt.Printf("%+v",get1)
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

//搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("db1").Type("table1").Do(context.Background())
	printEmployee("取所有", res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:yike")
	res, err = client.Search("db1").Type("table1").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee("字段相等", res, err)

	if res.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {

			var t Employee
			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}

			fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("db1").Type("table1").Query(q).Do(context.Background())
	printEmployee("条件查询", res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock")
	res, err = client.Search("db1").Type("table1").Query(matchPhraseQuery).Do(context.Background())
	printEmployee("短语搜索", res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("db1").Type("table1").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee("分析", res, err)

}

//简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("db1").
		Type("table1").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee("分页", res, err)

}

//打印查询到的Employee
func printEmployee(desc string, res *elastic.SearchResult, err error) {
	fmt.Printf("%#v\n", desc)
	if err != nil {
		print(err.Error() + "\n")
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func main() {
	//create()
	//delete()
	//update()
	//gets()
	query()
	list(5, 1)
}
