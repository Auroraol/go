#  json

> 1. **JSON 编码**: 将 Go 的数据结构转换为 JSON 格式的字符串。通过 `json.Marshal` 函数实现。
> 2. **JSON 解码**: 将 JSON 格式的字符串转换为 Go 的数据结构。通过 `json.Unmarshal` 函数实现。

相关的包：

```go
import {
	"encoding/json"
	"fmt"
}
```

## 一、序列化结构体

**使用 json.Marshal 将数据序列化为字节数组。适用于结构体、Map。**

```go
type Person struct {
	Name string
	Age int
	Sex int
}

func main() {
	person := Person{Name: "xiaoyu", Age: 18, Sex: 1}
	b, err := json.Marshal(person) // 将结构体序列号为json字节数组
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(b)) // 将json字节数组转化为字符串
}
```

打印出：

```
{"Name":"xiaoyu","Age":18,"Sex":1}
```

<span style="color:red">注意：如果结构体中的字段为小写字母开头，将不能被序列化，因为这相当于是一个局部私有属性</span>

如果想要输出的结果为小写字母开头的key，可以为结构体添加Tag：

```go
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}
```

这样的话，就可以输出：

```
{"name":"xiaoyu","age":18,"sex":1}
```

## 二、序列化Map

```go
func main() {
    person := make(map[string]interface{}) // value为interface{}表示值可以是多种类型
    person["name"] = "xiaoyu"
    person["age"] = 18
    person["sex"] = 1

    b, err := json.Marshal(person) // 将Map序列号为json字节数组
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(string(b)) // 将json字节数组转化为字符串
}
```

打印出：

```
{"age":18,"name":"xiaoyu","sex":1}
```

## 三、反序列化为结构体

**使用 json.Unmarshal 将字节数组反序列化。**

```go
func main() {
	personStr := `{"name":"xiaoyu","age":18,"sex":1}`
	person := new(Person)
	err := json.Unmarshal([]byte(personStr), &person) // 将json字符串转化为byte数组，再填充到person
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(person)
	fmt.Println(person.Name)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}
```

打印出：

```
&{xiaoyu 18 1}
xiaoyu
```

## **四、反序列化为Map**

```go
func main() {
    personStr := `{"name":"xiaoyu","age":18,"sex":1}`
    person := make(map[string]interface{})
    err := json.Unmarshal([]byte(personStr), &person) // 将json字符串转化为byte数组，再填充到person
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(person)
    fmt.Println(person["name"])
}
```

打印出：

```
map[age:18 name:xiaoyu sex:1]
xiaoyu
```
