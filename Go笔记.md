

#  :scroll: **Go 开发环境搭建**

## 一、概述 

  Go 编程语言是一个开源项目，它使程序员更具生产力。 Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。 <img src="Go%E7%AC%94%E8%AE%B0.assets/1608370194961-929008e0-0f51-43af-8895-0ccb5adedee1.png" alt="go128.png" style="zoom:50%;" />

Go 语言具有很强的表达能力，它简洁、清晰而高效。得益于其并发机制，用它编写的程序能够非常有效地利用多核与联网的计算机，其新颖的类型系统则使程序结构变得灵活而模块化。 Go 代码编译成机器码不仅非常迅速，还具有方便的垃圾收集机制和强大的运行时反射机制。 它是一个快速的、静态类型的编译型语言，感觉却像动态类型的解释型语言。 



Go语言自己的早期源码使用C语言和汇编语言写成。从 Go 1.5 版本后，完全使用Go语言自身进行编写。Go语言的源码对了解Go语言的底层调度有极大的参考意义。

●官网：[https://golang.org](https://golang.org/)

●GitHub：https://github.com/golang/go

●中文镜像：[https://golang.google.cn](https://golang.google.cn/)

### **Golang生态拓展介绍“站在巨人的肩膀上”**

![image-20231114142919414](Go%E7%AC%94%E8%AE%B0.assets/image-20231114142919414.png)

## 二、下载与安装:crossed_swords:

> 本文基于go version go1.14.2 linux/amd64

### 下载安装包

Go官网下载地址：https://golang.org/dl/

Go官方镜像站（推荐）：https://golang.google.cn/dl/

如果是window系统 推荐下载可执行文件版,一路 Next

这里以linux为例

![image-20231114143155098](Go%E7%AC%94%E8%AE%B0.assets/image-20231114143155098.png)

### 解压安装包

Linux 从 https://golang.org/dl/ 下载tar⽂件，并将其解压到 /usr/local。

将/usr/local/go/bin添加到PATH环境变量中。

### 在/home下新建go文件夹

```shell
[root@iZ2ze505h9bgsbp83ct28pZ src]# ll
总用量 131008
drwxr-xr-x. 2 root root         6 5月  11 2019 debug
-rw-r--r--  1 root root 123658438 4月   9 06:12 go1.14.2.linux-amd64.tar.gz
drwxr-xr-x. 3 root root        41 3月  29 12:13 kernels
[root@iZ2ze505h9bgsbp83ct28pZ src]# sudo su root
[root@iZ2ze505h9bgsbp83ct28pZ src]# tar -xvf go1.14.2.linux-amd64.tar.gz -C /usr/local/
[root@iZ2ze505h9bgsbp83ct28pZ src]# cd /usr/local/
[root@iZ2ze505h9bgsbp83ct28pZ local]# ls
aegis  bin  etc  games  go  include  lib  lib64  libexec  mysql  sbin  share  src
```

在/home/go目录里新建下面三个文件夹：

```
cd /home
mkdir go
cd /home/go
mkdir bin
mkdir src
mkdir pkg
```

![img](Go%E7%AC%94%E8%AE%B0.assets/1650470398898-b212fc73-7077-430d-bdee-7c04e15af23c.png)

### 配置GOROOT

把/usr/local/go/bin目录配置GOROOT 到环境变量里

```
sodu vim /etc/profile
```

![WH8IZUB_GC4F0MT~D7RBW~D](Go%E7%AC%94%E8%AE%B0.assets/WH8IZUB_GC4F0MTD7RBWD.png)

ctrl+shift+v

```shell
export GOROOT="/usr/local/go"
export GOPATH=$HOME/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
```

测试

```shell
source /etc/profile
go version
go env
```

![](Go%E7%AC%94%E8%AE%B0.assets/image-20231114145305482.png)

如果系统变量还是不能生效 每次新打开一个命令窗口都要重新输入 source /etc/profile 才能使go env 等配置文件生效： 那就加到用户变量,当前用户一登录就会加载到 解决方法：

在 ~/.bashrc 中添加语句（在root账号和子账号里都加一次）

```
source /etc/profile
```

保存退出

```
source /etc/profile 
或者
source $HOME/.profile
```

### GOPROXY

Go1.14版本之后，都推荐使用go mod模式来管理依赖了，也不再强制我们把代码必须写在GOPATH下面的src目录了，你可以在你电脑的任意位置编写go代码。

默认GoPROXY配置是：GOPROXY=https://proxy.golang.org,direct， 由于国内访问不到 https://proxy.golang.org 所以我们需要换一个PROXY，这里推荐使用https://goproxy.io 或 https://goproxy.cn。

可以执行下面的命令修改GOPROXY：

```shell
`go env -w GOPROXY=https://goproxy.cn,direct`
```

###  开发工具

+ vscode

+ Goland

## **三、Golang语言特性**

### Golang的优势

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650470888012-e20eedfd-9064-4d4e-b040-d6878aaa96ad.png" alt="3-golang优势1.png" style="zoom: 33%;" />

静态类型语言

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471318257-8884275c-9fe9-41de-8251-1bb828d50aa6.png" alt="7-golang优势2.png" style="zoom: 33%;" />

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471330308-12d4e2a3-7355-46f6-9af4-1ed10e39538e.png" alt="6-golang优势2.png" style="zoom: 80%;" />

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471363769-9ded1e7a-acb0-4d6c-b4c1-61f9b77589b0.png" alt="5-golan优势1.png" style="zoom: 50%;" />

```go
package main
  
import (
    "fmt"
    "time"
)

func goFunc(i int) {
    fmt.Println("goroutine ", i, " ...")
}

func main() {
    for i := 0; i < 10000; i++ {
        go goFunc(i) //开启一个并发协程
    }

    time.Sleep(time.Second)
}
```

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471374246-adf62ac5-7eba-45c9-bcd7-ccc2f8141640.png" alt="8-golang优势3.png" style="zoom: 80%;" />

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471446832-5722e0a9-5522-469b-9ea9-296c373e3d66.png" alt="9-golang优势4.png" style="zoom:67%;" />

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471454878-bf9c4abc-62c5-42f8-b595-b16d99d10743.png" alt="10-golang优势5.png" style="zoom:67%;" />

![11-golang优势6.png](Go%E7%AC%94%E8%AE%B0.assets/1650471465058-b5db8451-e1d8-4ce4-a572-cc3d8be9bdc1.png)

### Golang适合做什么

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471474504-6de5ec53-2447-4bdd-8a08-36ae40283ece.png" alt="12-golang优势7.png" style="zoom:67%;" />

(1)、云计算基础设施领域



代表项目：docker、kubernetes、etcd、consul、cloudflare CDN、七牛云存储等。



(2)、基础后端软件



代表项目：tidb、influxdb、cockroachdb等。



(3)、微服务



代表项目：go-kit、micro、monzo bank的typhon、bilibili等。



(4)、互联网基础设施



代表项目：以太坊、hyperledger等。

### Golang明星作品

DOCKER

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471498432-166e36fd-6294-460c-bbcd-96f6e784f8a9.png" alt="13-golang优势8.png" style="zoom: 25%;" />

KUBERNETES

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471506905-c3bf704e-d2fc-41e1-8e01-0a4907ae28fc.png" alt="14-golang优势9.png" style="zoom: 25%;" />





GITHUB

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650471515654-27569b7e-d67a-45f6-8d8f-54bd616da975.png" alt="15-golang优势10.png" style="zoom:25%;" />

### Golang的不足

1、包管理，大部分包都在github上

2、无泛化类型

(Golang 1.18+已经支持泛型)

3、所有Excepiton都用Error来处理(比较有争议)。

4、对C的降级处理，并非无缝，没有C降级到asm那么完美(序列化问题

## 四、VSCode 开发环境:crossed_swords:

远程连接服务器

> 类似Xshell似功能

**添加扩展：**

Remote-SSH插件

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114151658843.png" alt="image-20231114151658843" style="zoom:67%;" />

 **首次登录设置**

1. 请单击VS Code首页左下角的**远程链接符号**：

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114162914592.png" alt="image-20231114162914592" style="zoom:80%;" />

2. 选择：**Connect Current Window to Host**

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114151818522.png" alt="image-20231114151818522" style="zoom:67%;" />

3. 请输入用户名和IP地址（用户名以root为例）：

   ```
   ssh root@您的ip地址
   ```

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114151849531.png" alt="image-20231114151849531" style="zoom:67%;" />

​          3.0.  选择第一项：C:\User\hrwei\.ssh\config;  可以查看配置

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114151906344.png" alt="image-20231114151906344" style="zoom: 80%;" />

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114152227210.png" alt="image-20231114152227210" style="zoom:67%;" />

​         3.2 首次登录需要输入您的服务器SSH密码。

​         3.3 如果您能在VS Code终端看到服务器并可进行命令操作，则证明远程连接成功。

![image-20231114154001758](Go%E7%AC%94%E8%AE%B0.assets/image-20231114154001758.png)

设置文件权限

```shell
sudo chmod 777 -R GolangStudy/
```

vscoed打开该文件夹

![image-20231114154206028](Go%E7%AC%94%E8%AE%B0.assets/image-20231114154206028.png)

效果

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114154244652.png" alt="image-20231114154244652" style="zoom: 67%;" />

**一个经典的Hello World程序**

在线测试:  https://godbolt.org/

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

其中，main为入口函数，package 声明了包名，import 引入其他依赖包。

![image-20231113204547577](Go%E7%AC%94%E8%AE%B0.assets/image-20231113204547577.png)

## 五、常用命令

### 运行程序

```shell
$ go run index.go
Hello, World!
```

go run 命令的常用标记：

●-a：强制编译相关代码，不论它们的编译结果是否已是最新的

●-x：打印编译过程中所需运行的命令，并执行它们

●-n：打印编译过程中所需运行的命令，但并不执行

●-p n：并行编译，其中n为并行的数量

●-v：列出被编译的代码包的名称

●-a -v：列出所有被编译的代码包的名称

○go v1.3 中的所有：包含Go语言自带的标准库的代码包

○go v1.4 中的所有：不包含Go语言自带的标准库的代码包

●-work：显示编译时创建的临时工作目录的路径，并且不删除它

### 编译程序

```shell
$ go build index.go
```

在Windows下可以看到编译出了 exe 文件，可直接运行：

```shell
$ ./index.exe
Hello, World!
```

参考资料

●[Go语言入门教程](http://c.biancheng.net/golang/)

●[Go语言中文网](https://studygolang.com/)

●[Golang标准库文档](https://studygolang.com/pkgdoc)

●[Golang文档中文版镜像](http://docscn.studygolang.com/)

●[Google公布实现Go 1.5自举的计划](https://studygolang.com/articles/2419)【[英文原文地址](https://www.infoq.com/news/2015/01/golang-15-bootstrapped/)】

# **Golang语法新奇**

> 注意一下每个章节和GolangStudy文件夹内容一一对应

## **1、从一个main函数初见golang语法**

> 1-firstGolang

```go
package main //程序的包名

/* 方式1
import "fmt"
import "time"
*/

// 方式2(推荐)
import (
	"fmt"
	"time"
)


//main函数
func main() { //函数的{  一定是 和函数名在同一行的，否则编译错误
	//golang中的表达式，加";", 和不加 都可以，建议是不加
	fmt.Println(" hello Go!")

	time.Sleep(1 * time.Second)
}
```

终端运行

```shell
$ go run test1_hello.go 
Hello Go
```

![](Go%E7%AC%94%E8%AE%B0.assets/image-20231114163209768.png)

go run 表示 直接编译go语言并执行应用程序，一步完成

可以先编译，然后再执行

```shell
 $go build test1_hello.go 
 $./test1_hello
 Hello Go
```

![image-20231114163140423](Go%E7%AC%94%E8%AE%B0.assets/image-20231114163140423.png)

## **2、变量的声明**

> 2-var

声明变量的一般形式是使用 var 关键字

```go
package main

/*
	四种变量的声明方式
*/

import (
	"fmt"
)

//声明全局变量 方法一、方法二、方法三是可以的
var gA int = 100
var gB = 200

//用方法四来声明全局变量
// := 只能够用在 函数体内来声明
//gC := 200

func main() {
	//方法一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：(常用的方法) 省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	// =====
	fmt.Println("gA = ", gA, ", gB = ", gB)
	//fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	//多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
```

## **3、常量**

> 3-const_iota

```go
package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING = 10*iota	 //iota = 0
	SHANGHAI 		  //iota = 1
	SHENZHEN          //iota = 2
)

const (
	a, b = iota+1, iota+2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d				  // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f				  // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota *3  // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9 
	i, k					   // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)

func main() {
	//常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的。

	fmt.Println("BEIJIGN = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const() 一起使用， iota只有在const进行累加效果。
	//var a int = iota 

}
```

常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

```go
package main


import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)


func main(){
    println(a, b, c)
}

//unsafe.Sizeof(a)输出的结果是16 。
//字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
```

输出结果为：abc, 3, 16

### iota和表达式

iota总是用于 increment，但是它可以用于表达式，在常量中的存储结果值。

```go
type Allergen int


const (
    IgEggs Allergen = 1 << iota         // 1 << 0 which is 00000001
    IgChocolate                         // 1 << 1 which is 00000010
    IgNuts                              // 1 << 2 which is 00000100
    IgStrawberries                      // 1 << 3 which is 00001000
    IgShellfish                         // 1 << 4 which is 00010000
)
```


这个工作是因为当你在一个const组中仅仅有一个标示符在一行的时候，它将使用增长的iota取得前面的表达式并且再运用它。在 Go 语言的[spec](https://legacy.gitbook.com/book/aceld/how-do-go/edit#)中， 这就是所谓的隐性重复最后一个非空的表达式列表.

如果你对鸡蛋，巧克力和海鲜过敏，把这些 bits 翻转到 “on” 的位置（从左到右映射 bits）。然后你将得到一个 bit 值00010011，它对应十进制的 19。

```go
fmt.Println(IgEggs | IgChocolate | IgShellfish)


// output:
// 19
type ByteSize float64


const (
    _           = iota                   // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)       // 1 << (10*1)
    MB                                   // 1 << (10*2)
    GB                                   // 1 << (10*3)
    TB                                   // 1 << (10*4)
    PB                                   // 1 << (10*5)
    EB                                   // 1 << (10*6)
    ZB                                   // 1 << (10*7)
    YB                                   // 1 << (10*8)
)
```

## 4、函数

### **函数返回多个值**

> 4-function

```go
package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

//返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

//返回多个返回值， 有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)



	//r1 r2 属于foo3的形参，  初始化默认的值是0
	//r1 r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)


	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)


	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)

	ret1, ret2 = foo4("foo4", 444)
	fmt.Println("ret1 = ", ret1, " ret2 = ", ret2)
}

// a =  abc
// b =  555
// c =  100
// a =  haha
// b =  999
// ret1 =  666  ret2 =  777
// ---- foo3 ----
// a =  foo3
// b =  333
// r1 =  0
// r2 =  0
// ret1 =  1000  ret2 =  2000
// ---- foo4 ----
// a =  foo4
// b =  444
// ret1 =  1000  ret2 =  2000
```

![image-20231114165146843](Go%E7%AC%94%E8%AE%B0.assets/image-20231114165146843.png)

### init函数与import

> 5-init

**init函数：** init 函数可在package main中，可在其他package中，可在同一个package中出现多次。

**main函数:**  main 函数只能在package main中

**执行顺序**

+ golang里面有两个保留的函数(这两个函数在定义时不能有任何的参数和返回值)：
  + init函数（能够应用于所有的package）
  + main函数（只能应用于package main）
+ 程序的初始化和执行都起始于main包(注意是main包,不是面函数)

+ 如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。

+ 当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。

+ 等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。

下图详细地解释了整个执行过程：

<img src="Go%E7%AC%94%E8%AE%B0.assets/1650528765014-63d3d631-428e-4468-bc95-40206d8cd252.png" alt="img" style="zoom:80%;" />

虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。

go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

示例:

代码结构：<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114170019583.png" alt="image-20231114170019583" style="zoom: 67%;" />

Lib1.go

```go
package InitLib1

import "fmt"

func init() {
    fmt.Println("lib1")
}
```

Lib2.go

```go
package InitLib2

import "fmt"

func init() {
    fmt.Println("lib2")
}
```

main.go

```go
package main

import (
    "fmt"
	_"GolangStudy/5-init/lib1"
	_"GolangStudy/5-init/lib2"
)

func init() {
    fmt.Println("libmain init")
}

func main() {
    fmt.Println("libmian main")
}
```

运行结果

```
lib1
lib2
libmain init
libmian main
```

改动一个地方，Lib1包导入Lib2，main包不管

```go
package InitLib1

import (
    "fmt"
    _ "GolangTraining/InitLib2"
)

func init() {
    fmt.Println("lib1")
}
```

运行结果

```
lib2
lib1
libmain init
libmian main
```

总结:  

1. 程序的初始化和执行都起始于main包(注意是main包,不是面函数)
2. 先按包的顺序依次自动执行他们的init函数
3. 再执行main包init函数, 然后继续执行main函数

### import匿名及别名导包方式

+ 导入: 包路径

+ 使用 包名.方法

```go
package main

import (
    // 导入 包路径
    // "GolangStudy/5-init/lib2"   //默认,当没有调用导入包里的函数时会报错
	// _ "GolangStudy/5-init/lib1"  //匿名
	mylib2 "GolangStudy/5-init/lib2"  //别名
	// ."GolangStudy/5-init/lib2"   //最好不用
)

func main() {
    
    // 使用 包名.方法
	// lib1.lib1Test()
	//lib2.Lib2Test()
	mylib2.Lib2Test()
	//Lib2Test()
}
```

###  函数参数

> 6-pointer

函数如果使用参数，该变量可称为函数的形参。

形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数：

#### 值传递

值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

值传递来调用 swap() 函数：

```go
package main


import "fmt"


func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200


   fmt.Printf("交换前 a 的值为 : %d\n", a )
   fmt.Printf("交换前 b 的值为 : %d\n", b )


   /* 通过调用函数来交换值 */
   swap(a, b)


   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}


/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int


   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/


   return temp;
}
```

运行结果

```
交换前 a 的值为 : 100
交换前 b 的值为 : 200
交换后 a 的值 : 100
交换后 b 的值 : 200
```

#### **引用传递(指针传递)**

指针

Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务。

接下来让我们来一步步学习 Go 语言指针。

我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

以下实例演示了变量在内存中地址：

```go
package main

import "fmt"

func main() {
   var a int = 10   
   fmt.Printf("变量的地址: %x\n", &a  )  //变量的地址: 20818a220
}
```

引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：

```go
package main


import "fmt"


func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200


   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )


   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)


   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}


func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```

运行结果

```
交换前，a 的值 : 100
交换前，b 的值 : 200
交换后，a 的值 : 200
交换后，b 的值 : 100
```

## 5、defer(延迟函数)

> 7-defer

defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。

defer作用：

+ 释放占用的资源
+ 捕捉处理异常
+ 输出日志


结果

+ 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。

```go
package main

import "fmt"

func main() {
	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")


	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
```

![image-20231114202523826](Go%E7%AC%94%E8%AE%B0.assets/image-20231114202523826.png)

#### recover错误拦截

运行时panic异常一旦被引发就会导致程序崩溃。

Go语言提供了专用于“拦截”运行时panic的内建函数“recover”。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。


注意：recover只有在defer调用的函数中有效。

```go
package main

import "fmt"

func Demo(i int) {
	//定义10个元素的数组
	var arr [10]int
	//错误拦截要在产生错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产生panic异常  打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()
	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10

}

func main() {
	Demo(10)
	//产生错误后 程序继续
	fmt.Println("程序继续执行...")
}
```

运行结果

```
runtime error: index out of range
程序继续执行...
```

## 6、slice和map

### slice

> 8-slice

+ Go 语言切片是对数组的抽象。

+ Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片

```
var identifier []type
```

切片不需要说明长度。

或使用make()函数来创建切片:

```go
var slice1 []type = make([]type, len)

//也可以简写为
slice1 := make([]type, len)
```

也可以指定容量，其中capacity为可选参数

```
make([]T, len, capacity)
```

这里 len 是数组的长度并且也是切片的初始长度

### map

> 9-map

## **7、面向对象特征**

> 10-OOP





11









# Go 模块管理

## 一、模块管理基础命令

### 初始化模块

```go
go mod init <模块名>
```

比如：

```go
go mod init demo
```

初始化后，会在工程目录生成一个 go.mod 和 go.sum 文件。

![image-20231113204903724](Go%E7%AC%94%E8%AE%B0.assets/image-20231113204903724.png)

### 获取模块

使用 get 命令获取模块。

比如获取gin：

```go
go get github.com/gin-gonic/gin
```

模块将安装到 %GOPATH% 的 pkg 下。

安装好之后，打开  go.mod 看到：

```go
module demo
go 1.15
require github.com/gin-gonic/gin v1.6.3
```

go get 命令的常用标记：

●-d：只执行下载动作，而不执行安装动作
●-fix：在下载代码包后先执行修正动作，而后再进行编译和安装
●-u：利用网络来更新已有的代码包及其依赖包

可以到 https://pkg.go.dev/ 搜索需要获取的模块

### 查看依赖图

使用以下命令查看当前项目的依赖图：

```shell
$ go mod graph
demo github.com/gin-gonic/gin@v1.6.3
```

### 安装依赖图

如果是从远程仓库克隆的项目，里面包含依赖图，我们需要手动执行以下命令安装依赖图中的模块：

```shell
go mod download
```

## 二、换源

如果拉取依赖缓慢，可以换源到Goproxy中国：

●[Goproxy中国 - GitHub](https://github.com/goproxy)

●[Goproxy中国 - 官网](https://goproxy.cn/)

执行以下命令即可：

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

查看所有已经配置的环境变量：

```
$ go env
set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\quanzaiyu\AppData\Local\go-build
set GOENV=C:\Users\quanzaiyu\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=D:\Users\quanzaiyu\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=D:\Users\quanzaiyu\go
set GOPRIVATE=
set GOPROXY=https://goproxy.cn,direct
set GOROOT=D:\Development\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=D:\Development\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=D:\Workplace\temp\go_learn\go.mod
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0 -fdebug-prefix-map=C:\Users\QUANZA~1\AppData\Local\Temp\go-build015280026=/tmp/go-build -gno-record-gcc-s
witches
```

除此之外，还可以使用包管理工具：[gopm](https://github.com/gpmgo/gopm)

## 三、构建和安装

使用以下命令将在工程目录下构建一个 exe 文件：

```shell
go build
```

使用以下命令将在 %GOPAHT%/bin 下安装构建好的 exe 文件：

```shell
go install
```

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231113205532883.png" alt="image-20231113205532883" style="zoom:50%;" />

## 四、其他命令

### go mod tidy

整理依赖，比如清理掉无用的模块，添加用到的依赖。

### go mod verify

验证依赖，比如在依赖中包含了错误的版本号，会给出错误提示。

```go
// 未通过验证的情况
$ go mod verify
go: github.com/gin-gonic/gin@v1.6.4: reading github.com/gin-gonic/gin/go.mod at revision v1.6.4: unknown revision v1.6.4

// 通过验证的情况
$ go mod verify
all modules verified
```

verify 还会检查依赖包中的文件是否被修改，若被修改也会给出错误提示。

### go mod why

询问某个依赖在项目中使用与否：

```go
go mod why -m github.com/gin-gonic/gin
```

### go mod edit

先查看一下 go mod edit 的用法：

```go
$ go help mod edit
usage: go mod edit [editing flags] [go.mod]
...
```

修改当前模块的名字为test：

```go
go mod edit -module test
```

修改go的版本号：

```go
go mod edit -go=1.12
```

格式化 go.mod 文件：

```go
go mod edit -fmt
```

将某个依赖添加到项目中：

```go
go mod edit -require github.com/gin-gonic/gin@v1.6.3
```

排除某个依赖，被排除的依赖不能被拉取和安装：

```go
go mod edit -exclude github.com/gin-gonic/gin@v1.6.3
```

### go mod vendor

将项目中的依赖在 vendor 文件夹中复制一份

```go
go mod vendor
```

### go list

列出当前项目用到的所有依赖：

```go
go list -m all
```

# Go 语言基础

## 基础语法

### 一、输出语句

使用 fmt.Println 输出：

```go
fmt.Println("hello")
```

也可以使用 fmt.Printf 格式化输出：

```go
fmt.Println("%v %v %v %q\n", i, f, b, s)
```

输出而不换行使用：

```go
fmt.Println(10)
```

使用 println() 和 print() 也可以打印数据。

### 二、变量

```go
package main

import "fmt"

func main() {
    //同时定义多个变量：
    var name string = "xiaoyu"
    fmt.Println(name)

    var a1 int = 0
    fmt.Println(a1)

    var b1, c1 int = 1, 2
    fmt.Println(b1, c1)

    //可以将变量声明放于括号之中：
    var (
        a string = "1"
        b int = 2
        c int
    )
    fmt.Println(a, b, c)
}
```

运行结果

```go
xiaoyu
0
1 2
1 2 0
```

#### 零值

零值就是变量没有做初始化时系统默认设置的值

```go

// 数值类型（包括complex64/128）为 0
var a int
var a int8
var a int16
var a int32
var a int64
var a float32
var a float64
var a complex64 // 0+0i
var a complex128 // 0+0i

// 字符串为 ""（空字符串）
var a string

// 布尔类型为 false
var a bool

// 以下几种类型为 nil：
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
```

#### 通过 := 声明变量

可以省略 var ，使用 := 声明并为变量赋值

```go
a := 1
fmt.Print(a)       //必须是声明新的变量, 后进行变量赋值

a = 2              //只能赋值
fmt.Print(a)

nn := "mmmmm么么么么"
fmt.Println(nn)
```

同时声明多个变量：

```go
a, b := 1, 2 
println(a, b)
```

### 空白标识符

空白标识符 _ 常常被用于抛弃值，如值 5 在 _, b = 5, 7 中被抛弃。这在丢弃函数的某些返回值时非常有用。

### 三、常量

定义常量跟变量差不多，只是将关键字改为const

```go
const LENGTH int = 10
```

**iota**

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

```go
 const (
    a = iota   //0
    b          //1
    c          //2
    d = "ha"   //独立值，iota += 1
    e          //"ha"   iota += 1
    f = 100    //iota +=1
    g          //100  iota +=1
    h = iota   //7,恢复计数
    i          //8
  )
  fmt.Println(a,b,c,d,e,f,g,h,i) // 0 1 2 ha ha 100 100 7 8
```

再看一个有趣的例子：

```go
 const (
    i=1<<iota // 1: 1b; 1<<0 1b(1)
    j=3<<iota // 3: 11b; 3<<1: 110b(6)
    k // 3<<2: 1100b(12)
    l // 3<<3: 11000b(24)
  )
  fmt.Println(i,j,k,l) // 1 6 12 24
```

注：<<n==*(2^n)

### 四、关键字

下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

|          |             |        |           |        |
| -------- | ----------- | ------ | --------- | ------ |
| break    | default     | func   | interface | select |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

|        |         |         |         |        |         |           |            |         |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

### 五、运算符

+ 算术运算符：+ - * / % ++ --
+ 关系运算符：== != > < >= <=
+ 逻辑运算符：&& || !
+ 位运算符：& | ^ << >>
+ 赋值运算符：= += -= *= /= %= <<= >>= &= ^= |=

引用：

+ & 返回变量存储地址。示例：&a 将给出变量的实际地址。

指针：

+ 指针变量。示例：*a 是一个指针变量

**运算符优先级**

| 优先级 | 运算符           |
| ------ | ---------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |

## 数据类型

### 数据类型概述

Go 语言按类别有以下几种数据类型：

+ **布尔型**
  + 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。

+ **数字类型**
  + 整型 int、uint 和 uintptr 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
  + uint：32 或 64 位
  + uint8：无符号 8 位整型 (0 到 255)
  + uint16：无符号 16 位整型 (0 到 65535)
  + uint32：无符号 32 位整型 (0 到 4294967295)
  + uint64：无符号 64 位整型 (0 到 18446744073709551615)
  + int：与 uint 一样大小
  + int8：有符号 8 位整型 (-128 到 127)
  + int16：有符号 16 位整型 (-32768 到 32767)
  + int32：有符号 32 位整型 (-2147483648 到 2147483647)
  + int64：有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
  + uintptr：无符号整型，用于存放一个指针
  + float32：IEEE-754 32位浮点型数
  + float64：IEEE-754 64位浮点型数
  + complex64：32 位实数和虚数
  + complex128：64 位实数和虚数
  + byte：类似 uint8
  + rune：类似 int32

**字符串类型**

+ 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

**派生类型**

+ （a） 指针类型（pointer）

+    (b) 数组类型（array）
+ （c）结构化类型（struct）
+ （d) 通道类型（channel）
+ （e） 函数类型（func）
+ （f) 切片类型（slice）
+ （g）接口类型（interface）
+    (h) Map类型(map)
+    (a) 指针类型（pointer） 

### **数字**

#### **整数**

如果声明类型为 int 或 uint ，则其宽度（宽度即指存储一个某类型的值所需要的空间。空间的单位可以是比特，也可以是字节（byte）。）与计算机的计算架构有关

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231113213624853.png" alt="image-20231113213624853" style="zoom:80%;" />

除了这两个计算架构相关的整数类型之外，还有8个可以显式表达自身宽度的整数类型。如下表所示。 

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231113213736181.png" style="zoom: 67%;" />

整数类型值的表示范围

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231113213825107.png" style="zoom:80%;" />

#### 进制

●十进制数表示方法： num := 12 ，格式化参数使用 %d 表示

●八进制数表示方法： num := 014

●十六进制数表示方法： num := 0xC ，格式化参数使用 %x 表示

#### 浮点数

浮点数类型有两个，即float32和float64。存储这两个类型的值的空间分别需要4个字节和8个字节。

浮点数类型的值一般由整数部分、小数点“.”和小数部分组成。其中，整数部分和小数部分均由10进制表示法表示。不过还有另一种表示方法。那就是在其中加入指数部分。指数部分由“E”或“e”以及一个带正负号的10进制数组成。比如，3.7E-2表示浮点数0.037。又比如，3.7E+1表示浮点数37。

有时候，浮点数类型值的表示也可以被简化。比如，37.0可以被简化为37。又比如，0.037可以被简化为.037。

有一点需要注意，在Go语言里，浮点数的相关部分只能由10进制表示法表示，而不能由8进制表示法或16进制表示法表示。比如，03.7表示的一定是浮点数3.7。

在格式化参数中：

●%E 用于以带指数部分的表示法显示浮点数类型值

●%f 用于以通常的方法显示浮点数类型值

#### 复数

复数类型同样有两个，即complex64和complex128。存储这两个类型的值的空间分别需要8个字节和16个字节。实际上，complex64类型的值会由两个float32类型的值分别表示复数的实数部分和虚数部分。而complex128类型的值会由两个float64类型的值分别表示复数的实数部分和虚数部分。

复数类型的值一般由浮点数表示的实数部分、加号“+”、浮点数表示的虚数部分，以及小写字母“i”组成。比如，3.7E+1 + 5.98E-2i。正因为复数类型的值由两个浮点数类型值组成，所以其表示法的规则自然需遵从浮点数类型的值表示法的相关规则。

举例：

```go
num := 3.7E+1 + 5.98E-2i
fmt.Println(num) // (37+0.0598i)
```

在格式化参数中，复数仍然可以用 %E 和 %f 表示。

#### byte和rune

 byte与rune类型有一个共性，即：

+  它们都属于别名类型
  + byte是uint8的别名类型
  + rune则是int32的别名类型

byte类型的值需用8个比特位表示，其表示法与uint8类型无异。

一个rune类型的值即可表示一个Unicode字符。Unicode是一个可以表示世界范围内的绝大部分字符的编码规范。详细信息可以参看：http://unicode.org/，https://unicode-table.com/cn/。用于代表Unicode字符的编码值也被称为Unicode代码点。一个Unicode代码点通常由“U+”和一个以十六进制表示法表示的整数表示。例如，英文字母“A”的Unicode代码点为“U+0041”。

rune类型的值需要由单引号“'”包裹。例如，'A'或'昱'。这种表示方法一目了然。不过，我们还可以用另外几种形式表示rune类型值。见下表。  

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231113214158015.png" style="zoom:80%;" />

在rune类型值的表示中支持几种特殊的字符序列，即：转义符。它们由“\”和一个单个英文字符组成。如下表所示。

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231113214220627.png" style="zoom:67%;" />

举例：

```go
var char rune = '昱'
fmt.Printf("字符 %c 的 unicode 为 %x", char, char) // 字符 昱 的 unicode 为 6631
```

这意味着可以这样输出这个汉字：

```go

fmt.Printf("\u6631 is %x", '昱') // 昱 is 6631

if '\u6631' == '昱' {
    fmt.Println(1)
}
```

### 字符串

一个字符串类型的值可以代表一个字符序列。这些字符必须是被Unicode编码规范支持的。虽然从表象上来说是字符序列，但是在底层，一个字符串值却是由若干个字节来表现和存储的。一个字符串（也可以说字符序列）会被Go语言用Unicode编码规范中的UTF-8编码格式编码为字节数组。

注意：我们在一个字符串值或者一个字符串类型的变量之上应用Go语言的内置函数len将会得到代表它的那个字节数组的长度。这可能与我们看到的表象是不同的。

字符串的表示法有两种，即：原生表示法和解释型表示法。

+ 若用原生表示法，需用反引号“`”把字符序列包裹起来。
+ 若用解释型表示法，则需用双引号“"”包裹字符序列。

二者的区别是，前者表示的值是所见即所得的（除了回车符）。在那对反引号之间的内容就是该字符串值本身。而后者所表示的值中的转义符会起作用并在程序编译期间被转义。所以，如此表示的字符串值的实际值可能会与我们看到的表象不相同。

最后要注意，字符串值是不可变的。也就是说，我们一旦创建了一个此类型的值，就不可能再对它本身做任何修改。

**转义**

使用 \ 符号表示转义字符：

```go
var str string = "\\\""
fmt.Printf("%q ---> %s", str, str) // "\\\"" ---> \"
```

将转义字符放于 `` 中不会被转义：

```go
fmt.Println(`\\\"`) // \\\"
```

在格式化参数中：

+ %q 表示字符串的表象值，包括转移符及双引号
+ %s 表示字符串的真实值，转义后的值

### 数组

一个数组（Array）就是一个可以容纳若干类型相同的元素的容器。这个容器的大小（即数组的长度）是固定的，且是体现在数组的类型字面量之中的。

数组的声明：

```
var variable_name [SIZE] variable_type
```

数组的定义：

```
var variable_name = [SIZE]variable_type{初始化值1, 初始化值2, ...}
```

示例：

```go
nums := [...]int{1,2,3}
println(nums)
println(nums[0])
```

#### 定义新的数组类型

可以通过 type 关键字定义新的数组类型：

```go
type MyNumbers [3]int
nums := MyNumbers{1,2,3}

fmt.Println(nums)
```

#### 数组长度和容量

使用 len 方法获取数组长度， cap 方法获取数组的容量：

```go
nums := [3]int{1,2,3}
fmt.Println(len(nums), cap(nums))
```

### 切片

切片（Slice）与数组一样，也是可以容纳若干类型相同的元素的容器。与数组不同的是，无法通过切片类型来确定其值的长度。每个切片值都会将数组作为其底层数据结构。我们也把这样的数组称为切片的底层数组。

创建切片：

```
numbers := []int{1, 2, 3, 4, 5}
```

从数组中获取切片：

```
numbers := [5]int{1, 2, 3, 4, 5}
slice := numbers[1:4]
fmt.Println(slice, len(slice)) // [2 3 4] 3
```

### Map

Go语言的字典（Map）类型其实是哈希表（Hash Table）的一个实现。字典用于存储键-元素对（更通俗的说法是键-值对）的无序集合。注意，同一个字典中的每个键都是唯一的。如果我们在向字典中放入一个键值对的时候其中已经有相同的键的话，那么与此键关联的那个值会被新值替换。

Map的声明：

```go
* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)
```

Map的定义：

```go
var map_variable = map[key_data_type]value_data_type{key1: value1, key2: value2, ...}
```

```go
stringMap := map[string]string{}

stringMap["1"] = "A"
stringMap["2"] = "B"
stringMap["3"] = "V"
stringMap["4"] = "D"
```

### **类型转换**

```
type_name(expression)
```

## 控制流程

### 一、条件语句

#### **if...else**

```go
 var a int = 10

  if a < 0 {
    fmt.Printf("a < 0" )
  } else if a > 0 {
    fmt.Printf("a > 0" )
  } else {
    fmt.Println("a == 0")
  }
```

#### **switch**

go的switch语句不需要添加break

```go
  var grade string
  var marks int = 90

  switch marks {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50,60,70 : grade = "C"
    default: grade = "D"
  }

  switch {
    case grade == "A":
      fmt.Printf("优秀!\n" )
    case grade == "B", grade == "C":
      fmt.Printf("良好\n" )
    case grade == "D":
      fmt.Printf("及格\n" )
    case grade == "F":
      fmt.Printf("不及格\n" )
    default:
      fmt.Printf("差\n" )
  }
  fmt.Printf("你的等级是 %s\n", grade )
```

##### **Type Switch**

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

Type Switch 语法格式如下：

```go
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    /* 你可以定义任意个数的case */
    default: /* 可选 */
       statement(s);
}
```

示例：

```go
  var x interface{}

  switch i := x.(type) {
    case nil:
      fmt.Printf(" x 的类型 :%T",i)
    case int:
      fmt.Printf("x 是 int 型")
    case float64:
      fmt.Printf("x 是 float64 型")
    case func(int) float64:
      fmt.Printf("x 是 func(int) 型")
    case bool, string:
      fmt.Printf("x 是 bool 或 string 型" )
    default:
      fmt.Printf("未知型")
  }
```

输出：

```
 x 的类型 :<nil>
```

##### fallthrough

使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
示例：

```go
  var num = 50
  switch {
    case num > 10:
      fmt.Println("num > 10")
      fallthrough
    case num > 50:
      fmt.Println("num > 50")
      fallthrough
    default:
      fmt.Println("num <= 10")
  }
```

输出：

```
num > 10
num > 50
num <= 10
```

#### **select**

```go
// 定义多个通道
var chanInt chan int = make(chan int, 3)
var chanSting chan string = make(chan string)
var chanBool chan bool = make(chan bool)

func main() {
	go send()
	go receive()
	time.Sleep(time.Second * 10)
	chanSting <- "send over"
	chanBool <- true
}

// 向通道发送数据
func send() {
	time.Sleep(time.Second)
	chanInt <- 1
	time.Sleep(time.Second)
	chanInt <- 2
	time.Sleep(time.Second)
	chanInt <- 3
}

// 获取通道中的数据
func receive() {
	for {
		select {
			case num := <- chanInt:
				fmt.Println(num)
			case str := <- chanSting:
				fmt.Println(str)
			case <- chanBool:
				fmt.Println("运行结束")
            default:
                fmt.Println("unknown channel")
		}
	}
}
```

输出：

```go
1
2
3
send over
运行结束
```

### 二、循环语句

#### 循环数字

```go
//var sum int = 0;
sum := 0
for i := 0; i <= 10; i++ {
    sum += i
}
fmt.Println(sum)
```

#### 循环数组

如果指定了数组长度，而初始化时并未完全填充，则未填充的值为零值

```go
numbers := [6]int{1, 2, 3, 4}
for i, x:= range numbers {
    fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
}
```

输出

```
第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 4
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0
```

#### **循环Map**

```go
oldMap := map[int]string{1: "a", 2: "b"}
newMap := map[interface{}]interface{}{}
for key, value := range oldMap {
    newMap[key] = value
}
fmt.Println(newMap) // map[1:a 2:b]
```

#### **类while循环**

```go
sum := 1
for ; sum <= 10; {
    sum += sum
}
fmt.Println(sum)
```

```go
for sum <= 10{
    sum += sum
}
fmt.Println(sum)
```

#### **无限循环**

```go
for {
    fmt.Println("done")
}
```

```go
for true{
    fmt.Println("done")
}
```

#### **循环控制语句**

循环控制语句包括：

+ break
+ continue
+ goto

break 和 continue 不多说，重点说一下 goto 。
每一个语句都能对其添加标签，goto用于跳到指定标签处。

```go

/* 定义局部变量 */
var a int = 10

/* 循环 */
LOOP: for a < 20 {
    if a > 12 && a < 18 {
        /* 跳过迭代 */
        a = a + 1
        goto LOOP
    }
    fmt.Printf("a的值为 : %d\n", a)
    a++
}
```

```
a的值为 : 10
a的值为 : 11
a的值为 : 12
a的值为 : 18
a的值为 : 19
```

```go
/* 定义局部变量 */
var a int = 10

/* 循环 */
for a < 20 {
    if a > 12 {
        /* 跳过迭代 */
        a = a + 1
        goto DONE
    }
    fmt.Printf("a的值为 : %d\n", a)
    a++
}

DONE: fmt.Println("done")
```

```
a的值为 : 10
a的值为 : 11
a的值为 : 12
done
```

## 内键方法

### 一、make:crossed_swords:

使用make可以创建切片，array、map、chan类型的数据。其返回的是引用类型。

**创建切片：**

```go
mSlice := make([]string, 3)
mSlice[0] = "Dog"
mSlice[1] = "Cat"
mSlice[2] = "Tiger"
fmt.Println(mSlice) // [Dog Cat Tiger]
```

**创建Map：**

```go
mMap := make(map[int]string, 3)
mMap[10] = "Dog"
mMap[20] = "Cat"
mMap[30] = "Tiger"
fmt.Println(mMap) // map[10:Dog 20:Cat 30:Tiger]
```

**创建通道（chan）：**

```go
mChan := make(chan int)
close(mChan)
```

### 二、new

new返回传入类型的指针地址

```go
nMap := new(map[int]string)
fmt.Println(nMap)
fmt.Println(reflect.TypeOf(nMap))
```

输出：

```
&map[]
*map[int]string
```

可以看出，通过 new 方法出来的数据是一个指针类型的数据。

### 三、常见操作

#### **len、cap**

+ len：长度，支持的数据类型：string、array、slice、map、chan

+ cap：容量，支持的数据类型：slice、array、chan

make的第三个参数可以预指定切片的大小，但是随着append的扩容会动态增加：

```go
	mSlice := make([]int, 1, 2)
	mSlice[0] = 1
	fmt.Println(len(mSlice), cap(mSlice)) // 1 2

	mSlice = append(mSlice, 2)
	mSlice = append(mSlice, 3)
	fmt.Println(len(mSlice), cap(mSlice)) // 3 4
```

#### **close**

使用 close 关闭一个通道。

```go
mChan := make(chan int, 1) // 创建一个通道
	mChan <- 1 // 往channel中写数据
	close(mChan) // 关闭channel
```

如果怕害怕记最后写关闭通道的代码，可以将其放于前面，加上 defer 修饰，这样在代码执行到最后的时候都会执行关闭通道的操作：

```go
mChan := make(chan int, 1)
	defer close(mChan)
	mChan <- 1
```

#### **append**

append方法往切片中添加一个或多个元素。

```go
mSlice := make([]string, 1)
mSlice[0] = "dog"
mSlice = append(mSlice, "cat", "tiger")
fmt.Println(mSlice) // [dog cat tiger]
fmt.Println(len(mSlice)) // 3
fmt.Println(cap(mSlice)) // 3
```

len 方法返回切片的长度，cap 方法返回切片的容量。

#### **copy**

copy 可以往目标切片中拷贝一个已存在的切片。

```
mSlice := make([]string, 1)
mSlice[0] = "dog"
mSlice = append(mSlice, "cat", "tiger")

mSlice2 := make([]string, 2)
copy(mSlice2, mSlice)

fmt.Println(mSlice2) // [dog cat]
```

注意到，目标切片的长度设置为2，即使源切片长度为3，拷贝过来的仍然只有两个数据。

#### **delete**

delete 会删除map中指定key的元素

```
mMap := make(map[int]string)
mMap[10] = "dog"
mMap[20] = "cat"
delete(mMap, 10)
fmt.Println(mMap) // map[20:cat]
```

### 四、异常处理

在go语言中，使用 panic 和 recover 处理异常。
panic负责抛出异常：

```go
func main() {
	panicRecover()
}

func panicRecover() {
	panic("I am panic")
}
```

执行后，控制台可以看到异常信息：

![image-20231113224601663](Go%E7%AC%94%E8%AE%B0.assets/image-20231113224601663.png)

使用 recover 可以捕获异常：

```go

func main() {
	panicRecover()
}

func panicRecover() {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	panic("I am panic")
}
```

可以看到，控制台不报错了，只是打印出了异常信息：

```go
func main() {
	panicRecover()
}

func panicRecover() {
	defer coverPanic()
    // 创建一个error类型的异常
	panic(errors.New("I am a pannic!"))
}

func coverPanic() {
	msg := recover()
	switch msg.(type) {
	case string:
		fmt.Println("string panic: ", msg)
	case error:
		fmt.Println("error panic: ", msg)
	default:
		fmt.Println("unknown panic: ", msg)
	}
}
```

## 指针

### 一、指针的定义

跟C语言一样，使用 * 定义指针：

```go
  var num int = 10
  var numPoint *int = &num
  fmt.Printf("num 的值: %d \n", num)
  fmt.Printf("num 的地址: %x \n", &num)
  fmt.Printf("numPoint 指向的地址: %x \n", numPoint)
  fmt.Printf("numPoint 指向的地址的值: %d \n", *numPoint)
```

打印出：

```go
num 的值: 10
num 的地址: c0000140a0
numPoint 指向的地址: c0000140a0
numPoint 指向的地址的值: 10
```

如果定义一个指针却没给它赋值，此指针的值为 nil 

```go
  var numPoint *int
  fmt.Printf("%x \n", numPoint) // 0
  fmt.Println(numPoint) // <nil>
```

由此，判断一个指针时候为空就应该这样写：

```go
var numPoint *int
if numPoint != nil {
    fmt.Println("numPoint != nil")
} else {
    fmt.Println("numPoint is nil")
}
```

### 二、数组指针

数组指针是一个指向数组的指针：

```go
  arr := [...]int {1,2,3}
  arrPoint := &arr  //类似数组指针 int (*p)[5] = &arr
  fmt.Println("arrPoint 的值:", arrPoint)
  fmt.Println("arrPoint 的地址:", *arrPoint)
```

输出：

```
arrPoint 的值: &[1 2 3]
arrPoint 的地址: [1 2 3]
```

要取出数组指针中某一项指针指向的地址的值，需要添加括号后取出：

```
fmt.Println((*arrPoint)[0]) // 同arr[0]
fmt.Println((*arrPoint)[1]) // 同arr[1]
fmt.Println((*arrPoint)[2]) // 同arr[2]
```

要取出某项指针指向的地址，再在前面添加 & 符号：

```
fmt.Println(&(*arrPoint)[0]) // 同&arr[0]
fmt.Println(&(*arrPoint)[1]) // 同&arr[1]
fmt.Println(&(*arrPoint)[2]) // 同&arr[2]
```

### **三、指针数组**

指针数组是一个包含若干指针的数组：

```go
  a, b, c := 1, 2, 3
  pointArr := [...]*int {&a,&b,&c}
  fmt.Println("pointArr:", pointArr)
  fmt.Println("*pointArr[0]:", *pointArr[0])
  fmt.Println("*pointArr[1]:", *pointArr[1])
  fmt.Println("*pointArr[2]:", *pointArr[2])
```

输出：

```go
pointArr: [0xc0000140a0 0xc0000140a8 0xc0000140b0]
*pointArr[0]: 1
*pointArr[1]: 2
*pointArr[2]: 3
```

## 函数

```go
Test := func(a int) int{
    return a
}

b := Test(1000) 
fmt.Println(b)  //1000
```

### 一、闭包

Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。

匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

用法：

```go
var x int
inc := func() int {
    x++
    return x
}

fmt.Println(func() (a, b int) {
    return inc(), inc()
}())
```

输出值显而易见，是 1 2

### 二、可变长度的参数

使用 ... 指定参数的长度是可变的。

```go
var Test = func(values...int)  {
    for _, v := range values {
        fmt.Println(v)
    }
}

Test(1,2,3)
```

### 三、立即执行的函数

跟JS类似，go中也可以让函数立即执行，并且不需要像JS一样在函数之前添加 ; 

```go
func(values...int)  {
    for _, v := range values {
        fmt.Println(v)
    }
}(1,2,3)
```

```
1
2
3
```

### 四、返回多个值的函数

函数可以包含多个返回值：

```go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("A", "B")
	fmt.Println(a, b) // B A
}
```

### 五、值传递与引用传递

**按值传递**

```go
func swap(x int, y int) {
	var temp int

	temp = x
	x = y
	y = temp
}

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值为 : %d\n", a ) // 100
	fmt.Printf("交换前 b 的值为 : %d\n", b ) // 200

	/* 通过调用函数来交换值 */
	swap(a, b)
	fmt.Printf("交换后 a 的值 : %d\n", a ) // 100
	fmt.Printf("交换后 b 的值 : %d\n", b ) // 200
}
```

程序中使用的是值传递，所以两个值并没有实现交换。

**使用引用的示例：**(类似c++中的按指针传递)

```go
func swap(x *int, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值为 : %d\n", a ) // 100
	fmt.Printf("交换前 b 的值为 : %d\n", b ) // 200

	/* 通过调用函数来交换值 */
	swap(&a, &b)
	fmt.Printf("交换后 a 的值 : %d\n", a ) // 200
	fmt.Printf("交换后 b 的值 : %d\n", b ) // 100
}
```

由于传递的是引用，指向内存地址，所有值成功交换了。

### 六、延迟语句

使用 defer 关键字可以使函数中的某一条语句延迟执行，也就是在函数结束的时候才执行。

比如在读写文件时，可以将关闭文件写在前面，以防止遗忘关闭文件的操作：

```go
func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return ioutil.ReadAll(file)
}
```

#### defer匿名函数

当一个函数中存在多个defer语句时，可以将其放置于匿名函数中，它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序。

```go
func main() {
	fmt.Println(1)
	defer func() {
		fmt.Println(2)
		fmt.Println(3)
	}()
	fmt.Println("end")
}
```

#### 循环之中的defer

当defer放于循环之中，其输出值将反向输出。

```go
func main() {
	f := func(i int) int {
		fmt.Printf("%d ",i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}
```

输出：

```
1 2 3 4 40 30 20 10
```

#### 循环中defer匿名函数

当在循环中defer的是一个匿名函数，那么其输出值将会是最终循环值。

```go
for i := 1; i < 5; i++ {
    defer func() {
        fmt.Print(i)
    }()
}
```

输出

```
5555
```

原因是defer语句携带的表达式语句中的那个匿名函数包含了对外部（确切地说，是该defer语句之外）的变量的使用。注意，等到这个匿名函数要被执行（且会被执行4次）的时候，包含该defer语句的那条for语句已经执行完毕了。此时的变量i的值已经变为了5。因此该匿名函数中的打印函数只会打印出5。

正确的用法是：把要使用的外部变量作为参数传入到匿名函数中。（玩过JS的一定都懂）

```go
for i := 1; i < 5; i++ {
    defer func(n int) {
        fmt.Print(n)
    }(i)
}
```

### 七、递归函数

阶乘

```go
func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i int = 5
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
```

**斐波那契数列**

```go
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
```

## 面向对象

### 一、结构体

**声明结构体：**

```go
type Person struct {
	Name string
	Age int
	Sex int
}
```

**创建结构体：**

第一种方式

```go
person := Person{Name: "xiaoyu", Age: 18, Sex: 1}
fmt.Println(person)

// or
person := Person{"xiaoyu", 18, 1}
fmt.Println(person) // {xiaoyu 18 1}

// or
var person = Person{Name: "xiaoyu", Age: 18, Sex: 1}
fmt.Println(person) // {xiaoyu 18 1}
```

第二种方式

```go
var person Person
person.Name = "xiaoyu"
person.Age = 18
person.Sex = 1
fmt.Println(person) // {xiaoyu 18 1}
```

第三种方式：使用new方法

```go
person := new(Person)
person.Name = "xiaoyu"
person.Age = 18
person.Sex = 1
fmt.Println(person) // &{xiaoyu 18 1}
```

**结构体的方法:**

go语言中结构体的方法看起来像是在JS的原型中添加方法：

```go
func main() {
	person := Person{Name: "Xiaoyu", Age: 18, Sex: 1}
	person.Say("hello") // Xiaoyu say hello
}

// 声明
type Person struct {
	Name string
	Age int
	Sex int
}

func (person *Person) Say(words string) {
	fmt.Println(person.Name, "say", words)
}
```

尝试将Say换为say，发现会报错。go就是通过控制方法的大小写来控制其作用域，如果是小写只能内部调用。

**使用组合模拟类的继承**

在go中，通过组合的方式模拟面向对象中的 继承 ：

```go
func main() {
	person := Person{Name: "Xiaoyu", Age: 18, Sex: 1}
	person.Kind = "mammalia"
	person.Eat() // 好吃
	fmt.Println(person) // {{mammalia} Xiaoyu 18 1}
}

type Animal struct {
	Kind string
}

func (person *Animal) Eat() {
	fmt.Println("好吃")
}

type Person struct {
	Animal // 继承
	Name string
	Age int
	Sex int
}

func (person *Person) Say(words string) {
	fmt.Println(person.Name, "say", words)
}
```

**内嵌结构中的命名冲突**
如果一个结构体组合了多个结构体，而这些结构体中包含相同的字段名，那么就不能直接为这个字段赋值了，需要先指定是哪个结构体中的字段。

```go
type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func main() {
	c := &C{}
	c.A.a = 1 // 指定为A中的a字段
	c.B.a = 2 // 指定为B中的a字段
	fmt.Println(c) // &{{1} {2}}
}
```

**匿名结构体**
跟匿名函数一样，结构体也可以匿名化：

```go
msg := &struct {
		status   int
		data string
    }{ 200, "ok"}
	
	fmt.Println(msg)
```

**树状结构定义**

使用指针对结构体自引用，可定义一个树状结构的结构体：

```go
func main() {
	person := &Person{
		name: "Job Smith",
		children: []*Person{
			{
				name: "Bob Smith",
				children: []*Person{
					{
						name: "Joy Smith",
					},
				},
			},
			{
				name: "Bob Smith",
			},
		},
	}

	fmt.Println(person.children[0].name) // Bob Smith
}

type Person struct {
	name  string
	children []*Person
}
```

**构造函数**
go语言结构体的构造函数可以用一个 返回结构体自身类型的指针 函数进行模拟：

```go
func main() {
	person := NewPerson("xiaoyu", 18, 1)
	fmt.Println(person)
}

type Person struct {
	Name string
	Age int
	Sex int
}

func NewPerson(name string, age int, sex int) *Person {
	return &Person{name, age, sex}
}
```

### 二、接口

#### **定义接口**

```go
type Behavior interface {
	Eat() string
	Run() string
}
```

#### **实现接口**

在go中不需要像Java一样使用implements关键字来实现接口，只需要保证struct中的方法、并传入的参数类型、数量、返回值类型一致即可：

```go
type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name string
	Age int
	Place string
}

func (cat *Cat) Grow() {

}

func (cat *Cat) Move(str string) string {
	return ""
}

func main() {
	cat := Cat{"Kitty", 2, "House"}
	animal, ok := interface{}(&cat).(Animal)
	fmt.Printf("%v, %v \n", ok, animal) // true, &{Kitty 2 House}
}
```

接口的实现条件：

1. 接口的方法与实现接口的类型方法格式一致
2. 接口中所有方法均被实现

Go语言的接口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。

#### **通过接口实现多态**

```go
func main() {
	dog := Dog{Name: "wangwang", Age: 2}
	cat := Cat{Name: "mimi", Age: 2}
	fmt.Println(dog.Eat()) // dog: eat
	fmt.Println(dog.Run()) // dog: run
	fmt.Println(cat.Eat()) // cat: eat
	fmt.Println(cat.Run()) // cat: run
}

type Animal struct {
	Kind string
}

type Cat struct {
	Animal
	Name string
	Age int
}

func (cat *Cat) Eat() string {
	return "cat: eat"
}

func (cat *Cat) Run() string {
	return "cat: run"
}

type Dog struct {
	Animal
	Name string
	Age int
}

func (dog *Dog) Eat() string {
	return "dog: eat"
}

func (cat *Dog) Run() string {
	return "dog: run"
}
```

可以先声明一个接口类型的变量，然后通过 new 方法创建对应类型的实例：

```go
var b1 Behavior
b1 = new(Dog)
fmt.Println(b1.Eat())
fmt.Println(b1.Run())

var b2 Behavior
b2 = new(Cat)
fmt.Println(b2.Eat())
fmt.Println(b2.Run())
```

接口还可以当做方法的参数：

```go
func main() {
	var dog = new(Dog)
	action(dog)

	var cat = new(Cat)
	action(cat)
}

func action(b Behavior) {
	fmt.Println(b.Eat())
	fmt.Println(b.Run())
}
```

## 并发与多线程

### 一、并发

Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

降低加锁/解锁的频率去掉了元余的协程生命周期管理协程完成协程重用降低额外的延迟和开销Goroutine

![image-20231114210031346](Go%E7%AC%94%E8%AE%B0.assets/image-20231114210031346.png)

### 二、通道

通道（channel）是Go语言中一种非常独特的数据结构。它可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

```go
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据并把值赋给 v
```

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

```go
ch := make(chan int)
```

注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

使用 chose 关闭通道，关闭后的通道不能再向其输入：

```
close(ch)
```

通道工作示意图：

![image-20231114210209029](Go%E7%AC%94%E8%AE%B0.assets/image-20231114210209029.png)

#### 通道缓冲区

通道有带缓冲和非缓冲之分。缓冲通道中可以缓存N个数据。我们在初始化一个通道值的时候必须指定这个N。相对的，非缓冲通道不会缓存任何数据。发送方在向通道值发送数据的时候会立即被阻塞，直到有某一个接收方已从该通道值中接收了这条数据。非缓冲的通道值的初始化方法如下：





通过 make 的第二个参数设置缓冲区缓冲区大小：



带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。



不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。



注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。



示例：



如果将通道缓冲区改为1，则执行此段程序会报错：





#### 单向通道

默认情况下，通道都是双向的，即双向通道。如果数据只能在通道中单向传输，那么该通道就被称作单向通道。我们在初始化一个通道值的时候不能指定它为单向。但是，在编写类型声明的时候，我们却是可以这样做的。例如：



类型Receiver代表了一个只可从中接收数据的单向通道类型。这样的通道也被称为接收通道。在关键字chan左边的接收操作符<-形象地表示出了数据的流向。相对应的，如果我们想声明一个发送通道类型，那么应该这样：



这次<-被放在了chan的右边，并且“箭头”直指“通道”。我们可以把一个双向通道值赋予上述类型的变量，就像这样：



但是，反之则是不行的。像下面这样的代码是通不过编译的：



单向通道的主要作用是约束程序对通道值的使用方式。比如，我们调用一个函数时给予它一个发送通道作为参数，以此来约束它只能向该通道发送数据。又比如，一个函数将一个接收通道作为结果返回，以此来约束调用该函数的代码只能从这个通道中接收数据。



举个例子：





#### 通道阻塞

通道阻塞的条件：

●1.输入Channel的数据量>Channel能接受的量

●2. Channel输出的数据量>Channel内现有的数据量



不阻塞的情况：输入数据量<=channel缓冲区大小 && 输出数据量<=channel缓冲区大小





阻塞的情况1：输入的数据量 > channel缓冲区大小





阻塞的情况2：输出的数据量 > channel缓冲区大小





### 三、协程

#### 协程的相关概念

协程是轻量级的线程。

协程少一道手续：线程申请内存，需要走过内核协程申请内存，不需要走过内核上下文切换更快协程大概2KB的内存申请量一个线程可以包含多个协程线程大约8MB的内存申请量协程的内存消耗更小协程的优势









协程的创建

通常执行程序，都是顺序执行的：



上面的程序，无论执行多少次，都会依次输出：





在go语言中，使用 go 关键字启动一个协程。



可以看到，控制台输出将会两个循环依次执行：



而之所以在主程序中添加 time.Sleep ，是因为运行太快的话，协程还未将输出结果发送到控制台，主线程就结束了。



如果设置不同的间隔时间，将会发现两个协程一个快一个慢。



输出：





#### 与主线程并发执行

主线程实际上也可看做一条协程，比如如下程序，主线程中的循环将于协程中的循环一起执行：



输出：





设置CPU最大核心数





### 四、协程通讯

#### 协程与主线程之间的通讯

以下实例通过两个 goroutine 来计算数字之和：



在协程完成了一次相加操作之后，将相加的结果通过通道传递，再从主线程中取出通道中的数据。



#### 多个协程之间通讯

go中的协程通讯需要用到通道（chan），一个简单的例子如下：



send 和 receive 两个方法本身应该并发执行，由于send方法不停地向 chanInt 通道发送数据，而 receive 方法不停地读取chanInt 通道的数据，即可完成协程间的通讯。



#### 通过select从通道中读取数据

如果存在多个通道，可以使用select语句选择从不同的通道中读取数据。



输出：





### 五、协程同步

协程同步需要用到sync.WaitGroup 工具，相关的方法有：

●Add 添加协程记录

●Done 移出协程记录

●Wait 同步等待所有记录的携程全部结束

打印出的结果为：





注意Add的次数跟Done的次数得一致，否则会抛出一个死锁的错误：



比如将 read 方法中的 WG.Add(1) 换为 WG.Add(i) ，这样相当于加的delta为3，而只Done了2次。就会报上面的错误。从其源码中就可以看出，每次Done是将其delta减一。

DonedecrementstheWaitGroupcounter

byone.

*WaitGroup)DoneO

func

Wg

wg.Add(delta:-1)

![image.png](Go%E7%AC%94%E8%AE%B0.assets/1608542996337-f19203ad-2d6f-4fa0-b4c3-049c31ddfe7a.png)







并发的应用

抢票问题

经典的抢票程序：有10张票，有100个人来抢票，每个人最多只能购买1张。





执行后，每次输出会不一样：

```
第3个人抢到了第6张票
第0个人抢到了第0张票
第1个人抢到了第1张票
第5个人抢到了第2张票
第6个人抢到了第3张票
第2个人抢到了第4张票
第7个人抢到了第5张票
第11个人抢到了第8张票
第8个人抢到了第9张票
第4个人抢到了第7张票
```



## 包

### 一、创建并引入包

创建如下目录结构：<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114210541929.png" alt="image-20231114210541929" style="zoom: 67%;" />

在 math.go 中：

```
package utils

func Add(x,y int) int {
  return x + y
}

func Sub(x,y int) int {
  return x - y
}
```



在 index.go 中引入：

```
package main

import (
  "./utils"
  "fmt"
)

func main() {
  fmt.Println(utils.Add(1,1))
  fmt.Println(utils.Sub(1,1))
}
```

### 二、包的别名

如果想要引入一个包，而不使用其默认名称，可以为其添加别名。
比如：

```
import f "fmt"

f.Println(1)
```

如果连前缀都不想添加，可以使用本地化的导入，方法是在包前面添加 . ：

```
import ."fmt"

Println(v)
```

如果导入了某个包而在程序中没有调用，可以使用 _ 符号丢弃此包：

```
import _"fmt"
```

## 链表

#### 创建链表结构

```
ype Node struct {
	Data  int
	Next  *Node
}
```

使用：

```
linkedList := &Node{1, &Node{
		2, &Node{
			3,
			nil,
		},
	}}
```

**遍历链表**

```
func (p *Node) Traverse() {
	for p != nil {
		fmt.Println((*p).Data)
		p = p.Next
	}
}
```

使用：

```
linkedList.Traverse()
```

# **Go语言标准库概述**

## Go语言常用标准库列表

| Go语言标准库包名 | 功  能                                                       |
| ---------------- | ------------------------------------------------------------ |
| bufio            | 带缓冲的 I/O 操作                                            |
| bytes            | 实现字节操作                                                 |
| container        | 封装堆、列表和环形列表等容器                                 |
| crypto           | 加密算法                                                     |
| database         | 数据库驱动和接口                                             |
| debug            | 各种调试文件格式访问及调试功能                               |
| encoding         | 常见算法如 JSON、XML、Base64 等                              |
| flag             | 命令行解析                                                   |
| fmt              | 格式化操作                                                   |
| go               | Go语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改 |
| html             | HTML 转义及模板系统                                          |
| image            | 常见图形格式的访问及生成                                     |
| io               | 实现 I/O 原始访问接口及访问封装                              |
| math             | 数学库                                                       |
| net              | 网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等                |
| os               | 操作系统平台不依赖平台操作封装                               |
| path             | 兼容各操作系统的路径操作实用函数                             |
| plugin           | Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载        |
| reflect          | 语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值 |
| regexp           | 正则表达式封装                                               |
| runtime          | 运行时接口                                                   |
| sort             | 排序接口                                                     |
| strings          | 字符串转换、解析及实用函数                                   |
| time             | 时间接口                                                     |
| text             | 文本模板及 Token 词法器                                      |



##  testing

testing 包用于go的单元测试。

引入：

```
import "testing"
```

### 一、测试规则

要创建测试文件，需要将文件名设置为 xxx_test.go ，结尾必须是 _test.go 。

测试文件中的，结构必须为如下形式：

```go
func TestXXX(t *testing.T) {
    ...
}

func BenchmarkXXX(t *testing.B) {
    ...
}
```

+ 单元测试方法必须以 Test 开头，压力测试方法必须以 Benchmark 开头
+ 单元测试参数必须是 (t *testing.T) ，性能测试参数必须是 (t *testing.B)

执行测试：

```
go test 测试文件名
```

### 二、单元（功能）测试

demo.go 

```go
package test

func GetArea(weight int, height int) int {
	return weight * height
}
```

demo_test.go

```go
package test

import "testing"

func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}
```

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114212110556.png" style="zoom:67%;" />

### 三、性能（压力）测试

```
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
```

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114212317030.png" style="zoom:67%;" />

执行完压力测试输出：

<img src="Go%E7%AC%94%E8%AE%B0.assets/image-20231114212252373.png" alt="image-20231114212252373" style="zoom:67%;" />

##  sort

sort 包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。

相比之下，Go语言的 sort.Sort 函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型 sort.Interface 来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法：

```go
package sort
type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool // i，j是序列元素的指数。
    Swap(i, j int)        // 交换元素
}
```

### 一、自定义实现排序接口

比如我们自定义一个字符串数组类型，要对字符串数组中的每一条字符串进行排序：

```go
// 将[]string定义为MyStringList类型
type StringList []string

// 实现sort.Interface接口的获取元素数量方法
func (m StringList) Len() int {
	return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m StringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m StringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func StringSort() {
	// 准备一个内容被打乱顺序的字符串切片
	names := StringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
```

### 二、使用内置接口实现排序

sort 包中的 StringSlice 的代码与上面自定义的 StringList 的实现代码几乎一样。因此，只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序。

```
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)
	fmt.Println(names)
```

### 三、sort 包中内建的类型排序接口

Go语言中的 sort 包中定义了一些常见类型的排序方法，如下表所示。

| 类  型                | 实现 sort.lnterface 的类型 | 直接排序方法               | 说  明            |
| --------------------- | -------------------------- | -------------------------- | ----------------- |
| 字符串（String）      | StringSlice                | sort.Strings(a [] string)  | 字符 ASCII 值升序 |
| 整型（int）           | IntSlice                   | sort.Ints(a []int)         | 数值升序          |
| 双精度浮点（float64） | Float64Slice               | sort.Float64s(a []float64) | 数值升序          |



### 四、对结构体进行排序

除了基本类型的排序，也可以对结构体进行排序。结构体比基本类型更为复杂，排序时不能像数值和字符串一样拥有一些固定的单一原则。结构体的多个字段在排序中可能会存在多种排序的规则，例如，结构体中的名字按字母升序排列，数值按从小到大的顺序排序。一般在多种规则同时存在时，需要确定规则的优先度，如先按名字排序，再按年龄排序等。

示例：

```go
type HeroKind int
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

// 实现sort.Interface接口取元素数量方法
func (s Heros) Len() int {
	return len(s)
}

// 实现sort.Interface接口比较元素方法
func (s Heros) Less(i, j int) bool {
	// 如果英雄的分类不一致时, 优先对分类进行排序
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	// 默认按英雄名字字符升序排列
	return s[i].Name < s[j].Name
}

// 实现sort.Interface接口交换元素方法
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortStruct1() {
	// 准备英雄列表
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	// 使用sort包进行排序
	sort.Sort(heros)
	// 遍历英雄列表打印排序结果
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
```

### 五、sort.Slice(推荐)

从 Go 1.8 开始，Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法。sort.Slice() 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类似于JavaScript中的 Array.prototype.sort。

函数的定义如下：

```
func Slice(slice interface{}, less func(i, j int) bool)
```

示例：

```go
type HeroKind int
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

func SortStruct2() {
	heros := []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
```

##  json

相关的包：

```go
import {
	"encoding/json"
	"fmt"
}
```

### 一、序列化结构体

使用 json.Marshal 将数据序列化为字节数组。适用于结构体、Map。

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

注意：如果结构体中的字段为小写字母开头，将不能被序列化，因为这相当于是一个局部属性

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

### 二、序列化Map

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

### 三、反序列化为结构体

使用 json.Unmarshal 将字节数组反序列化。

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

### **四、反序列化为Map**

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

##  **os**

相关的包：

```
import{
	"fmt"
	"io/ioutil"
	"os"
}
```

### 一、获取当前工作路径

```
currentPath, _ := os.Getwd()
fmt.Println(currentPath)
```



### **二、获取文件（夹）信息**

```
	filePath := "c:/windows/notepad.exe"
	stat, _ := os.Stat(filePath)
	fmt.Println("文件（夹）名：", stat.Name())
	fmt.Println("文件大小：", stat.Size())
	fmt.Println("权限：", stat.Mode())
	fmt.Println("是否是文件夹：", stat.IsDir())
	fmt.Println("修改时间：", stat.ModTime())
```

输出：

```
文件名： notepad.exe
文件大小： 202240
权限： -rw-rw-rw-
是否是文件夹： false
修改时间： 2020-09-15 14:16:51.9560539 +0800 CST
```

先读取文件再获取文件信息：

```
	filePath := "c:/windows"
	file, _ := os.Open(filePath)
	fmt.Println(file.Name())
	fmt.Println(file.Stat())
```



### **三、列出子文件（夹）**

只列出子文件（夹）名字：

```
    filePath := "c:/windows"
	file, _ := os.Open(filePath)
	for i := 1; ; i++{
		names, err := file.Readdirnames(10 * i)
		if err != nil {
			break
		}
		for _, name := range names{
			fmt.Println(name)
		}
	}
```

获取所有子文件：

```
    filePath := "c:/windows"
	file, _ := os.Open(filePath)
	for i := 1; ; i++{
		names, err := file.Readdir(10 * i)
		if err != nil {
			break
		}
		for _, file := range names{
			if file.IsDir() {
				continue
			}
			fmt.Println(file.Name())
		}
	}
```



### **四、创建、修改、删除、拷贝、重命名**

创建文件夹:

```
    err := os.Mkdir("d:/test", 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建文件夹成功")
```

创建文件：

```go
    _, err := os.Create("d:/test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建文件成功")
```

注意：如果有同名的文件夹存在，创建同名文件将会失败

重命名（移动文件）：

```go
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("重命名成功")
```

删除文件：

```go
	err := os.Remove("d:/test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("删除文件成功")
```

拷贝文件：暂时没有发现现成的API，可以自己封装一个拷贝文件的方法。

```go
func CopyFile(srcPath string, destPath string) {
	data, err := ioutil.ReadFile(srcPath)
	fmt.Println(string(data))
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
		return
	}
	err = ioutil.WriteFile(destPath, data, 0666)
	if err != nil {
		fmt.Printf("文件拷贝失败=%v\n", err)
		return
	}
	fmt.Println("拷贝成功")
}

// 调用
CopyFile("d:/test", "d:/test1")
```

##  **io**

相关的包：

```
import (
	"archive/zip"
    "archive/tar"
	"bytes"
	"bufio"
    "encoding/gob"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"io"
	"os"
	"strings"
)
```

