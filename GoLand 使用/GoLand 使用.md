```
GO111MODULE='on'
GONOPROXY='*.xiaoduoai.com'
GONOSUMDB='*.xiaoduoai.com'
GOPRIVATE='*.xiaoduoai.com,*.xiaoduotech.com'
GOPROXY='https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct'
```





# go mod 引入指定的分支

方法1:

```sh
go get git地址@分支名
如：go get github.com/golang/go@master
```

分支的代码必须是合入的，如果还在评审中的，是不会拉取的

方法2:

![image-20240921172120829](GoLand%20%E4%BD%BF%E7%94%A8.assets/image-20240921172120829.png)

# 更新依赖

```sh
$ go mod tidy
  go: downloading pkg.xx.com/service/lib/db/v6 v6.0.13
$ go mod vendor
```

如果通过上面清除缓存的方式，vendor依赖仓库的内容还没有更新到最新的提交，可以在最新提交分支上重新创建一个新的tag，修改go.mod中依赖的tag为新创建的tag，再使用上面的`go mod tidy` 和 `go mod vendor`分别更新go mod和vendor依赖。
