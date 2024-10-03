# go-elasticsearch

安装go-elasticsearch包

　　　　go get -u github.com/elastic/go-elasticsearch

　　**elasticsearch** 包将两个单独的包联系在一起：**esapi** 和 **estransport**，分别用于调用 Elasticsearch API 和通过 HTTP 传输数据

　　简单示例：

```
import (
    "github.com/elastic/go-elasticsearch/v7"
    "log"
)

func main() {
    // 默认连接地址：http://localhost:9200
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        log.Fatalf("Error creating the client: %s", err)
    }

    res, err := es.Info()
    if err != nil {
        log.Fatalf("Error getting response: %s", err)
    }

    log.Println(res)
}[![复制代码](elasticsearch.assets/copycode.gif)](javascript:void(0);)
```

 　连接并查询示例：

```
import (
    "context"
    "encoding/json"
    "github.com/elastic/go-elasticsearch/v7"
    "log"
    "strings"
)

func main() {
    // es连接配置
    cfg := elasticsearch.Config{
        Addresses: []string{
            "http://localhost:9200",
        },
        Username: "elastic",
        Password: "XcF8EbPPmgRgiLqoVAcI",
    }
    es, err := elasticsearch.NewClient(cfg)
    if err != nil {
        log.Fatalf("Error creating the client: %s", err)
    }

    // 搜索
    res, err := es.Search(es.Search.WithContext(context.Background()),
        es.Search.WithIndex("megacorp"),
        es.Search.WithBody(strings.NewReader(`{"query" : { "match" : { "last_name" : "Smith" } }}`)),
        es.Search.WithTrackTotalHits(true),
        es.Search.WithPretty(),
    )
    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }
    // 反序列化结果到map中
    var result map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        log.Fatalf("Error parsing the response body: %s", err)
    }

    // Print the ID and document source for each hit.
    for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
        log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
    }

    log.Println(result)

    defer res.Body.Close()
}
```

**bulk**

go elasticsearch bulk 操作：

```go
res, err := es.Bulk(
	strings.NewReader(`
        { "index" : { "_index" : "test", "_id" : "1" } }
        { "field1" : "value1" }
        { "delete" : { "_index" : "test", "_id" : "2" } }
        { "create" : { "_index" : "test", "_id" : "3" } }
        { "field1" : "value3" }
        { "update" : {"_id" : "1", "_index" : "test"} }
        { "doc" : {"field2" : "value2"} }
    `),
)
fmt.Println(res, err)
```

附录：

　　Elasticsearch Go Client：https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/index.html

　　Go 官方文档：https://pkg.go.dev/github.com/elastic/go-elasticsearch

 