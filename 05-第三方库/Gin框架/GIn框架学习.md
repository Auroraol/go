# Go Gin框架学习

**什么是Gin？**

　　Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。 它是一个类似于 martini 但拥有更好性能的 API 框架, 多亏了 httprouter，速度提高了近 40 倍。具有良好的性能和生产力。

　　而且封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点

　　对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错

　　文档：https://gin-gonic.com/zh-cn/docs/

**Gin安装使用：**

　　要求：Go 1.13 及以上版本

　　要安装Gin软件包，您需要安装Go并首先设置Go工作区。

　　　　1、首先需要安装Go（需要1.13+版本），然后可以使用下面的Go命令安装Gin。

　　　　　　**go get -u github.com/gin-gonic/gin**

　　　　2、将 gin 引入到代码中：

　　　　　　**import "github.com/gin-gonic/gin"**

　　　　3、（可选）如果使用诸如 http.StatusOK 之类的常量，则需要引入 net/http 包：

　　　　　　**import "net/http"**  

**快速开始：**

　　1、创建一个名为GinDemo.go的文件， touch GinDemo.go

　　2、GinDemo.go代码如下：

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

　　3、执行 go run GinDemo.go 命令来运行代码：

```go
# 运行 example.go 并且在浏览器中访问 0.0.0.0:8080/ping
$ go run example.go
```

 　启动日志：

```go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:    export GIN_MODE=release
 - using code:    gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

**Gin路由：**

　　**基本路由：**gin框架中采用的路由是基于httprouter做的，也支持Restful风格的API

　　API参数：可以通过Context的Param方法来获取API参数

　　URL参数：可以通过DefaultQuery()或Query()方法获取，DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串

　　表单参数：表单传输为post请求，http常见的传输格式为四种：

　　　　application/json

　　　　application/x-www-form-urlencoded

　　　　application/xml

　　　　application/form-data

　　　　表单参数可以通过PostFrom()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

　　**routes group:**routes group是为了管理一些相同的URL

```go
func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 路由组1 ，处理GET请求
   v1 := r.Group("/v1")
   // {} 是书写规范
   {
      v1.GET("/login", login)
      v1.GET("submit", submit)
   }
   v2 := r.Group("/v2")
   {
      v2.POST("/login", login)
      v2.POST("/submit", submit)
   }
   r.Run(":8000")
}
```

　　运行main方法即可启动gin项目：go run main.go

**Jsoniter：**

　　Gin 使用 encoding/json 作为默认的 json 包，但是你可以在编译中使用标签将其修改为 jsoniter。

　　　　go build -tags=jsoniter

**Goland新建Gin项目：**

  1、新建go modules，File-New-Project-Go modules

![img](GIn%E6%A1%86%E6%9E%B6%E5%AD%A6%E4%B9%A0.assets/1238257-20220606151941117-116756375.png)

    此时项目中只有一个go.mod文件，类似于Java中Maven的pom.xml

  2、下载Gin框架依赖

    命令行进入项目目录中，下载并安装gin：go get -u github.com/gin-gonic/gin

    此时go.mod：

```go
module filter-search

go 1.16

require (
	github.com/gin-gonic/gin v1.8.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
)
```

  3、新建main.go文件

   项目根目录新建cmd文件夹，cmd目录下新建 main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 路由请求到处理器
	// router.GET("/search",handler)
	// 指定服务端口
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// 监听并启动服务
	server.ListenAndServe()
}
```

  4、go run main.go 或者右键 go build filter-search

    访问：http://localhost:8080/ping

  5、新建目录

    internal目录，用于存放业务实现代码（类似于Java中的src），其下根据具体子业务新建子目录。handler放在此目录下

    init目录，用于存放项目启动即初始化的逻辑代码

    deploy目录存放部署相关代码

**Gin日志：**

  1、如何记录日志

```go
func main() {
    // 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
    gin.DisableConsoleColor()

    // 记录到文件。
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // 如果需要同时将日志写入文件和控制台，请使用以下代码。
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run(":8080")
}
```

  2、请求路由日志格式

   默认的请求路由日志格式：

```go
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
```

  如果你想要以指定的格式（例如 JSON，key values 或其他格式）记录信息，则可以使用 gin.DebugPrintRouteFunc 指定格式。

  在下面的示例中，我们使用标准日志包记录所有路由，但你可以使用其他满足你需求的日志工具  

```go
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
```

  3、自定义日志输出文件

```go
func main() {
	router := gin.New()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}
```