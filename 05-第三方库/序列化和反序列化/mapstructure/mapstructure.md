# mapstructure 库

## 1、基础使用

安装方式：

```sh
go get github.com/mitchellh/mapstructure@v1.5.0
```

在日常开发中，接受的数据可能不是固定的格式，而是会根据某个值的不同有不同的内容。

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
    // json数据
	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)  //反序列化为结构体
		if err != nil {
			log.Fatal(err)
		}
		
        //
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person:", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat:", cat)
		}
	}
}
```

运行结果：

```css
person: {dj 18 programmer}
cat: {kitty 1 Ragdoll}
```

我们定义了两个结构体`Person`和`Cat`，他们的字段有些许不同。现在，我们约定通信的 JSON 串中有一个`type`字段。当`type`的值为`person`时，该 JSON 串表示的是`Person`类型的数据。当`type`的值为`cat`时，该 JSON 串表示的是`Cat`类型的数据。

上面代码中，我们先用`json.Unmarshal`将字节流解码为`map[string]interface{}`类型。然后读取里面的`type`字段。根据`type`字段的值，再使用`mapstructure.Decode`将该 JSON 串分别解码为`Person`和`Cat`类型的值，并输出。

实际上，Google Protobuf 通常也使用这种方式。在协议中添加消息 ID 或**全限定消息名**。接收方收到数据后，先读取协议 ID 或**全限定消息名**。然后调用 Protobuf 的解码方法将其解码为对应的`Message`结构。从这个角度来看，`mapstructure`也可以用于网络消息解码，**如果你不考虑性能的话**。

## 2、详细使用

### 2.1、Field Tags (字段标签)

默认情况下，`mapstructure`使用结构体中字段的名称做这个映射，例如我们的结构体有一个`Name`字段，`mapstructure`解码时会在`map[string]interface{}`中查找键名`name`。

注意，这里的`name`是大小写不敏感的！

```go
type Person struct {
  Name string
}
```

当然，我们也可以指定映射的字段名。设置`mapstructure`标签。例如下面使用`username`代替上例中的`name`：

```go
type Person struct {
  Name string `mapstructure:"username"`
}
```

示例：

```go
type Person struct {
  Name string `mapstructure:"username"`
  Age  int
  Job  string
}

type Cat struct {
  Name  string
  Age   int
  Breed string
}

func main() {
  datas := []string{`
    { 
      "type": "person",
      "username":"dj",           // 字段映射为username
      "age":18,                  // 大小写不区分 
      "job": "programmer"
    }
  `,
    `
    {
      "type": "cat",
      "name": "kitty",
      "Age": 1,
      "breed": "Ragdoll"
    }
  `,
    `
    {
      "type": "cat",
      "Name": "rooooose",
      "age": 2,
      "breed": "shorthair"
    }
  `,
  }

  for _, data := range datas {
    var m map[string]interface{}
    err := json.Unmarshal([]byte(data), &m)
    if err != nil {
      log.Fatal(err)
    }

    switch m["type"].(string) {
    case "person":
      var p Person
      mapstructure.Decode(m, &p)
      fmt.Println("person", p)

    case "cat":
      var cat Cat
      mapstructure.Decode(m, &cat)
      fmt.Println("cat", cat)
    }
  }
}
```

上面代码中，我们使用标签`mapstructure:"username"`将`Person`的`Name`字段映射为`username`，在 JSON 串中我们需要设置`username`才能正确解析。另外，注意到，我们将第二个 JSON 串中的`Age`和第三个 JSON 串中的`Name`首字母大写了，但是并没有影响解码结果。`mapstructure`处理字段映射是大小写不敏感的。

### 2.2、Renaming Fields

在实际使用过程中，我们可能需要重命名 `mapstructure` 查找的键，这个时候，可以使用 "mapstructure" 标签并直接设置一个值。例如，要将上面的 "username" 示例更改为 "user"：

```go
type User struct {
    Username string `mapstructure:"user"`
}
```

### 2.3、Embedded Structs and Squashing（内嵌结构）

结构体可以任意嵌套，嵌套的结构被认为是拥有该结构体名字的另一个字段。

例如，下面两种`Friend`的定义方式对于`mapstructure`是一样的：

```go
type Person struct {
  Name string
}

// 方式一
type Friend struct {
  Person
}

// 方式二
type Friend struct {
  Person Person
}
```

为了正确解码，`Person`结构的数据要在`person`键下：

```go
map[string]interface{} {
  "person": map[string]interface{}{"name": "dj"},
}
```

我们也可以设置`mapstructure:",squash"`将该结构体的字段提到父结构中：

```go
type Friend struct {
  Person `mapstructure:",squash"`
}
```

这样只需要这样的 JSON 串，无效嵌套`person`键：

```go
map[string]interface{}{
  "name": "dj",
}
```

看**例子1**：

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
}

type Friend1 struct {
	Person
}

type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
    { 
      "type": "friend1",
      "person": {
        "name":"dj"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Println("friend1", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Println("friend2", f2)
		}
	}
}
```

结果：

```fsharp
friend1 {{dj}}
friend2 {{dj2}}
Exiting.
```

注意对比`Friend1`和`Friend2`使用的 JSON 串的不同。

接着看这个**例子2**：

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Type string
}

type Friend1 struct {
	Type string
	Person
}

type Friend2 struct {
	Type   string
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
    { 
      "type": "friend1",
      "person": {
        "name":"dj"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Printf("friend1: %+v \n", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Printf("friend2: %+v \n", f2)
		}
	}
}
```

结果：

```css
friend1: {Type:friend1 Person:{Name:dj Type:}} 
friend2: {Type:friend2 Person:{Name:dj2 Type:friend2}} 
```

注意对比`Friend1`和`Friend2`使用的 JSON 串的不同。

另外需要注意一点，如果父结构体中有同名的字段，那么`mapstructure`会将JSON 中对应的值**同时设置到这两个字段中**，即这两个字段有相同的值。

### 2.4、Remainder Values (未映射的值)

如果源数据中有未映射的值（即结构体中无对应的字段），`mapstructure`默认会忽略它。

解决方法:

1. 可以通过在 `DecoderConfig` 中设置 `ErrorUnused` 来引发错误。如果正在使用元数据（Metadata），还可以维护一个未使用键的切片（slice）。
2. 可以在结构体中定义一个字段，为其设置`mapstructure:",remain"`标签。这样未映射的值就会添加到这个字段中。注意，这个字段的类型只能为`map[string]interface{}`或`map[interface{}]interface{}`这两种类型之一。

示例:

```go
package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Friend struct {
	Name  string
	Other map[string]interface{} `mapstructure:",remain"`
}

func main() {
	m := map[string]interface{}{
		"name":    "bob",
		"address": "123 Maple St.",
	}

	var f Friend
	err := mapstructure.Decode(m, &f) //err != nil 表示正确响应
	fmt.Println("err->", err)
	fmt.Printf("friend: %+v", f)
}
```

结果：

```css
err-> <nil>
friend: {Name:bob Other:map[address:123 Maple St.]}
```

### 2.5、Omit Empty Values(忽略空值)

我们在使用 json 库时，对于空值我们不需要展示的时候，可以使用 `"json:,omitempty" `来忽略。 mapstructure 也是一样的。

当从结构体解码到其他任何值时，你可以在标签上使用 ",omitempty" 后缀，以便在该值等于零值时省略它。所有类型的零值在 Go 规范中有明确定义。

例如，数值类型的零值是零（"0"）。如果结构体字段的值为零且是数值类型，该字段将为空，且不会被编码到目标类型中。

```go
type Source struct {
    Age int `mapstructure:",omitempty"`
}
```

### 2.6、Unexported fields

Go 中规定了 未导出的（私有的）结构体字段不能在定义它们的包之外进行设置，解码器将直接跳过它们。

通过以下例子来进行讲解：

```go
package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Exported struct {
	private string // 首字母小写就表示私有的
	Public  string
}

func main() {
	m := map[string]interface{}{
		"private": "I will be ignored",
		"Public":  "I made it through!",
	}

	var e Exported
	_ = mapstructure.Decode(m, &e)
	fmt.Printf("e: %+v", e)
}
```

结果：

```css
e: {private: Public:I made it through!}
```

###  2.7、Other Configuration

mapstructure是高度可配置的。有关支持的其他功能和选项，请参阅 [DecoderConfig](https://pkg.go.dev/github.com/mitchellh/mapstructure#DecoderConfig) 结构。

### 2.8、逆向转换

在反向解码时，可以为某些字段设置`mapstructure:",omitempty"`。这样当这些字段为空值时，就不会出现在结构的`map[string]interface{}`中：

```go
type Person struct {
  Name string
  Age  int
  Job  string `mapstructure:",omitempty"`
}

func main() {
  p := &Person{
    Name: "dj",
    Age:  18,
  }

  var m map[string]interface{}
  mapstructure.Decode(p, &m)

  data, _ := json.Marshal(m)
  fmt.Println(string(data))
}
```

结果：

```go
$ go run main.go 
{"Age":18,"Name":"dj"}
```

### 2.9、Metadata

解码时会产生一些有用的信息，`mapstructure`可以使用`Metadata`收集这些信息。`Metadata`结构如下：

```go
// Metadata 包含关于解码结构的信息，这些信息通常通过其他方式获取起来会比较繁琐或困难。
type Metadata struct {
	// Keys 是成功解码的结构的键
	Keys []string

	// Unused 是一个键的切片，在原始值中被找到，但由于在结果接口中没有匹配的字段，所以未被解码
	Unused []string

	// Unset 是一个字段名称的切片，在结果接口中被找到，
	// 但在解码过程中未被设置，因为在输入中没有匹配的值
	Unset []string
}
```

`Metadata`只有3个导出字段：

- `Keys`：解码成功的键名；
- `Unused`：在源数据中存在，但是目标结构中不存在的键名。
- `Unset`：在目标结构中存在，但是源数据中不存在。

使用`DecodeMetadata`方法：

```go
package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Sex  bool
}

func main() {
	m := map[string]interface{}{
		"name": "dj",
		"age":  18,
		"job":  "programmer",
	}

	var p Person
    //定义一个Metadata结构
	var metadata mapstructure.Metadata
    //传入DecodeMetadata收集解码的信息
	mapstructure.DecodeMetadata(m, &p, &metadata)

	fmt.Printf("成功解码的结构的键keys:%#v 源数据未被解码unused:%#v, 结果接口中被找到unset: %#v \n", metadata.Keys, metadata.Unused, metadata.Unset)
}
```

结果

```css
成功解码的结构的键keys:[]string{"Name", "Age"} 源数据未被解码unused:[]string{"job"}, 结果接口中被找到unset: []string{"Sex"} 
```

### 2.10、字段类型错误处理

`mapstructure`执行转换的过程中不可避免地会产生错误，例如 JSON 中某个键的类型与对应 Go 结构体中的字段类型不一致。`Decode/DecodeMetadata`会返回这些错误：

```go
type Person struct {
  Name   string
  Age    int
  Emails []string
}

func main() {
  m := map[string]interface{}{
    "name":   123,
    "age":    "bad value",
    "emails": []int{1, 2, 3},
  }

  var p Person
  err := mapstructure.Decode(m, &p)
  if err != nil {
    fmt.Println(err.Error())
  }
}
```

上面代码中，结构体中`Person`中字段`Name`为`string`类型，但输入中`name`为`int`类型；字段`Age`为`int`类型，但输入中`age`为`string`类型；字段`Emails`为`[]string`类型，但输入中`emails`为`[]int`类型。

结果：

```sh
$ go run main.go 
5 error(s) decoding:

* 'Age' expected type 'int', got unconvertible type 'string'
* 'Emails[0]' expected type 'string', got unconvertible type 'int'
* 'Emails[1]' expected type 'string', got unconvertible type 'int'
* 'Emails[2]' expected type 'string', got unconvertible type 'int'
* 'Name' expected type 'string', got unconvertible type 'int'
```

### 2.11、弱类型输入

不想对结构体字段类型和`map[string]interface{}`的对应键值做强类型一致的校验。

使用`WeakDecode/WeakDecodeMetadata`方法，它们会尝试做类型自动转换。

注意: 如果类型转换失败了，`WeakDecode`同样会返回错误。

```go
type Person struct {
  Name   string
  Age    int
  Emails []string
}

func main() {
  m := map[string]interface{}{
    "name":   123,
    "age":    "18",
    "emails": []int{1, 2, 3},
  }

  var p Person
  err := mapstructure.WeakDecode(m, &p)
  if err == nil {
    fmt.Println("person:", p)   // 打印此行代码
  } else {
    fmt.Println(err.Error())
  }
}
```

### 2.12、解码器

`mapstructure`还提供了更灵活的解码器（`Decoder`）。可以通过配置`DecoderConfig`实现上面介绍的任何功能：

```go
// DecoderConfig 是用于创建新解码器的配置，允许自定义解码的各个方面。
type DecoderConfig struct {
	// DecodeHook，如果设置了，将在任何解码和任何类型转换（如果 WeaklyTypedInput 打开）之前调用。
	// 这允许你在将值设置到结果结构之前修改它们的值。
	// DecodeHook 会为输入中的每个映射和值调用一次。这意味着如果结构体具有带有 squash 标签的嵌入字段，
	// 解码钩子只会一次使用所有输入数据进行调用，而不是为每个嵌入的结构体分别调用。
	//
	// 如果返回错误，整个解码将以该错误失败。
	DecodeHook DecodeHookFunc

	// 如果 ErrorUnused 为 true，则表示在解码过程中存在于原始映射中但未被使用的键是错误的（多余的键）。
	ErrorUnused bool

	// 如果 ErrorUnset 为 true，则表示在解码过程中存在于结果中但未被设置的字段是错误的（多余的字段）。
	// 这仅适用于解码为结构体。这还将影响所有嵌套结构体。
	ErrorUnset bool

	// ZeroFields，如果设置为 true，在写入字段之前将字段清零。
	// 例如，一个映射在放入解码值之前将被清空。如果为 false，映射将会被合并。
	ZeroFields bool

	// 如果 WeaklyTypedInput 为 true，则解码器将进行以下“弱”转换：
	//
	//   - 布尔值转换为字符串（true = "1"，false = "0"）
	//   - 数字转换为字符串（十进制）
	//   - 布尔值转换为 int/uint（true = 1，false = 0）
	//   - 字符串转换为 int/uint（基数由前缀隐含）
	//   - int 转换为布尔值（如果值 != 0 则为 true）
	//   - 字符串转换为布尔值（接受：1、t、T、TRUE、true、True、0、f、F、
	//     FALSE、false、False。其他任何值都是错误的）
	//   - 空数组 = 空映射，反之亦然
	//   - 负数转换为溢出的 uint 值（十进制）
	//   - 映射的切片转换为合并的映射
	//   - 单个值根据需要转换为切片。每个元素都会被弱解码。
	//     例如："4" 如果目标类型是 int 切片，则可以变为 []int{4}。
	//
	WeaklyTypedInput bool

	// Squash 将压缩（squash）嵌入的结构体。也可以通过使用标签将 squash 标签添加到单个结构体字段中。例如：
	//
	//  type Parent struct {
	//      Child `mapstructure:",squash"`
	//  }
	Squash bool

	// Metadata 是将包含有关解码的额外元数据的结构。
	// 如果为 nil，则不会跟踪任何元数据。
	Metadata *Metadata

	// Result 是指向将包含解码值的结构体的指针。
	Result interface{}

	// 用于字段名称的标签名称，mapstructure 会读取它。默认为 "mapstructure"。
	TagName string

	// IgnoreUntaggedFields 忽略所有没有明确 TagName 的结构字段，类似于默认行为下的 `mapstructure:"-"`。
	IgnoreUntaggedFields bool

	// MatchName 是用于匹配映射键与结构体字段名或标签的函数。
	// 默认为 `strings.EqualFold`。可以用来实现区分大小写的标签值、支持蛇形命名等。
	MatchName func(mapKey, fieldName string) bool
}
```

示例：

```go
type Person struct {
  Name string
  Age  int
}

func main() {
  m := map[string]interface{}{
    "name": 123,
    "age":  "18",
    "job":  "programmer",
  }

  var p Person
  var metadata mapstructure.Metadata

  decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
    WeaklyTypedInput: true,
    Result:           &p,
    Metadata:         &metadata,
  })

  if err != nil {
    log.Fatal(err)
  }

  err = decoder.Decode(m)
  if err == nil {
    fmt.Println("person:", p)
    fmt.Printf("keys:%#v, unused:%#v\n", metadata.Keys, metadata.Unused)
  } else {
    fmt.Println(err.Error())
  }
}
```

