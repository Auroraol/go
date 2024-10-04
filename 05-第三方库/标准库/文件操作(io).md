# 一 文件操作

## 1.1 获取文件描述信息 os.Stat()

Go的os包中定义了file类，封装了文件描述信息，同时也提供了Read、Write的实现。  

```go
	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println("stat err: ", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)		// *os.fileStat
```

获取到的fileInfo内部包含 `文件名Name()`、`大小Size()`、`是否是目录IsDir()` 等操作。  


## 1.2 路径、目录操作

```go
	// 路径操作
	fmt.Println(filepath.IsAbs("./test.txt"))	// false：判断是否是绝对路径
	fmt.Println(filepath.Abs("./test.txt"))		// 转换为绝对路径

	// 创建目录
	err := os.Mkdir("./test", os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err: ", err)
		return
	}

	// 创建多级目录
	err = os.MkdirAll("./dd/rr", os.ModePerm)
	if err != nil {
		fmt.Println("mkdirAll err: ", err)
		return
	}
```

贴士：Openfile()可以用于打开目录。

## 1.3 删除文件

```go
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println("remove err:", err)
		return 
	}
```

该函数也可用于删除目录（只能删除空目录）。如果要删除非空目录，需要使用 `RemoveAll()` 函数

## 1.4 新建文件

新建文件可以通过如下两个方法

- func Create(name string) (file *File, err Error)

	根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。

- func NewFile(fd uintptr, name string) *File
	
	根据文件描述符创建相应的文件，返回一个文件对象

```go
    f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f)              // 打印文件指针
    f.Close()                   // 打开的资源在不使用时必须关闭
```

使用Create()创建文件时：

- 如果文件不存在，则创建文件。
- 如果文件存在，则清空文件内内容。  
- Create创建的文件任何人都可以读写。 

## 1.5 打开文件

- func Open(name string) (file *File, err Error)

	该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。

- func OpenFile(name string, flag int, perm uint32) (file *File, err Error)	该函数的三个参数：

	- 参数1：要打开的文件路径
	- 参数2：文件打开模式，如 `O_RDONLY`，`O_WRONGLY`，`O_RDWR`，还可以通过管道符来指定文件不存在时创建文件
	- 参数3：文件创建时候的权限级别，在0-7之间，常用参数为6

```go
	f, err := os.OpenFile("test.txt", os.O_APPEND | os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	f.Close()
```

常用的文件打开模式：

```go
	O_RDONLY 	int = syscall.O_RDONLY		// 只读
	O_WRONGLY	int = syscall.O_WRONGLY		// 只写
	O_RDWR 		int = syscall.O_RDWR		// 读写
	O_APPEND 	int = syscall.O_APPEND		// 写操作时将数据追加到文件末尾
	O_CREATE 	int = syscall.O_CREATE		// 如果不存在则创建一个新文件
	O_EXCL 		int = syscall.O_EXCL		// 打开文件用于同步I/O
	O_TRUNC		int = syscall.O_TRUNC		// 如果可能，打开时清空文件
```

## 1.6 写文件
写文件函数：

- func (file *File) Write(b []byte) (n int, err Error)

	写入byte类型的信息到文件

- func (file *File) WriteAt(b []byte, off int64) (n int, err Error)

	在指定位置开始写入byte类型的信息

- func (file *File) WriteString(s string) (ret int, err Error)

	写入string信息到文件
	

**写文件的示例代码**

```Go

package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "test.txt"
	fout, err := os.Create(userFile)		
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
    
	for i := 0; i < 10; i++ {
         // 按字符串写 WriteString()：
		fout.WriteString("Just a test!\r\n")
        
        // 按字节写 Write()：
		fout.Write([]byte("Just a test!\r\n"))
	}
}

```
**带缓冲的写入：**

```go
file, err := os.Openfile(path, O_WRONLY | O_CREATE, 0666)
if err != nil {
	fmt.Printf("%v", err)
	return
}
defer file.Close()
writer := bufio.NewWriter(file)
for l := 0; i < 5; i++ {
	writer.Writetring("hello\n")
}

writer.Flush()
```
**修改文件的读写指针位置 `Seek()`，包含两个参数：**

- 参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
- 参数2：偏移的开始位置，包括：
  - io.SeekStart：从文件起始位置开始
  - io.SeekCurrent：从文件当前位置开始
  - io.SeekEnd：从文件末尾位置开始

`Seek()`函数返回

```go
	f, _ := os.OpenFile("test.txt",os.O_RDWR, 6)
	off, _ := f.Seek(5, io.SeekStart)
	fmt.Println(off)							// 5
	n, _ := f.WriteAt([]byte("111"), off)
	fmt.Println(n)
	f.Close()
```

## 1.7 文件读取

文件读写的接口位于io包，file文件类是这些接口的实现类。

### 一 文件读取

### 1.1 直接读取 read()

read() 实现的是按字节数读取：

```go
	readByte := make([]byte, 128)			// 指定要读取的长度
	for {
		n, err := f.Read(readByte)			// 将数据读取如切片，返回值 n 是实际读取到的字节数
		if err != nil && err != io.EOF{		// 如果读到了文件末尾：EOF 即 end of file
			fmt.Println("read file : ", err)
			break
		}

		fmt.Println("read: ", string(readByte[:n]))
		if n < 128 {
			fmt.Println("read end")
			break
		}
	}
```

### 1.2 bufio的写操作

bufio封装了io.Reader、io.Writer接口对象，并创建了另一个也实现了该接口的对象：bufio.Reader、bufio.Writer。通过该实现，bufio实现了文件的缓冲区设计，可以大大提高文件I/O的效率。  

使用bufio读取文件时，先将数据读入内存的缓冲区（缓冲区一般比要比程序中设置的文件接收对象要大），这样就可以有效降低直接I/O的次数。  

`bufio.Read([]byte)`相当于读取大小`len(p)`的内容：

- 当缓冲区有内容时，将缓冲区内容全部填入p并清空缓冲区
- 当缓冲区没有内容且`len(p)>len(buf)`，即要读取的内容比缓冲区还要大，直接去文件读取即可
- 当缓冲区没有内容且`len(p)<len(buf)`，即要读取的内容比缓冲区小，读取文件内容并填满缓冲区，并将p填满
- 以后再次读取时，缓冲区有内容，将缓冲区内容全部填入p并清空缓冲区（和第一步一致）

示例：

```go
	// 创建读对象
	reader := bufio.NewReader(f)

	// 读一行数据
	byt, _ := reader.ReadBytes('\n')			
	fmt.Println(string(byt))
```

ReadString() 函数也具有同样的功能，且能直接读取到字符串数据，无需转换，示例：读取大文件的全部数据

```go
	reader := bufio.NewReader(f)
	for {										// 按照缓冲区读取：读取到特定字符结束
		str, err := reader.ReadString('\n')		// 按行读取
		if err != nil && err != io.EOF {
			fmt.Println("read err: ", err)
			break
		}
		fmt.Println("str = ", str)
		if err == io.EOF {
			fmt.Print("read end")
			break
		}
	}
```

在Unix设计思想中，一切皆文件，命令行输入也可以作为文件读入：

```go
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString("-")		// 假设命令行以 - 开始
```

缓冲的思想：通过bufio，数据被写入用户缓冲，再进入系统缓冲，最后由操作系统将系统缓冲区的数据写入磁盘。  

### 1.3 io/ioutil 包文件读取

ioutil直接读取文件：

```go
	ret, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("read err :", err)
		return
	}
	fmt.Println(string(ret))
```

### 二 文件写入

### 2.1 直接写

```go
	f, err := os.OpenFile("test.txt", os.O_CREATE | os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("write err:", err)
	}
	fmt.Println(n)								// 每次都会从头开始重新写入
```

上述案例中，如果我们不想每次写入都会从头开始重新写入，那么需要将打开模式修改为：`os.O_CREATE | os.O_WRONLY | os.O_APPEND`

### 2.2 bufio的写操作

```go
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString("hello world!")
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	writer.Flush()		// 必须刷新缓冲区：将缓冲区的内容写入文件中。如果不刷新，则只会在内容超出缓冲区大小时写入
```

### 2.3  io/ioutil 包文件写入

```go
	s := "你好世界"
	err := ioutil.WriteFile("test.txt", []byte(s), os.ModePerm)
```

### 三 文件读取偏移量

文件读取时，是可以控制光标位置的：

```go
	f, err := os.OpenFile("test.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f.Close()

	// 读取前五个字节，假设读取的文件内容为： hello world!
	bs := []byte{0}		// 创建1个字节的切片
	_, err = f.Read(bs)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("读到的数据是：", string(bs))	// h

	// 移动光标
	_, err = f.Seek(4, io.SeekStart)		// 光标从开始位置(h之前)，移动4位，到达o之前
	if err != nil {
		fmt.Println("seek err:", err)
		return
	}
	_, err = f.Read(bs)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("读到的数据是：", string(bs))		// o
```

通过记录光标的位置，可以实现断点续传：假设已经下载了1KB文件，即本地临时文件存储了1KB，此时断电，重启后通过本地文件大小、Seek()方法获取到上次读取文件的光标位置即可实现继续下载！
