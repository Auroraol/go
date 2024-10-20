# Go语言中使用gopsutil进行系统信息采集

## gopsutil包介绍

`psutil`是一个跨平台进程和系统监控的Python库，而`gopsutil`是其Go语言版本的实现。

安装`go get github.com/shirou/gopsutil`

## gopsutil包的使用

### 采集CPU相关信息

```bash
 1package main
 2
 3import (
 4    "fmt"
 5    "github.com/shirou/gopsutil/cpu"
 6)
 7
 8func main() {
 9    c, _ := cpu.Info()
10    fmt.Println("cpu信息:",c)
11    /*用户CPU时间/系统CPU时间/空闲时间。。。等等
12    用户CPU时间：就是用户的进程获得了CPU资源以后，在用户态执行的时间。
13    系统CPU时间：用户进程获得了CPU资源以后，在内核态的执行时间。
14    */
15    c1,_ := cpu.Times(false)
16    fmt.Println("cpu1:",c1)
17    
18    // CPU使用率，每秒刷新一次
19    for{
20     c2, _ := cpu.Percent(time.Duration(time.Second), false)
21     fmt.Println(c2)
22    }
23}
```

Copy

运行后输出结果如下

```bash
1cpu信息: [{"cpu":0,"vendorId":"GenuineIntel","family":"6","model":"61","stepping":4,"physicalId":"","coreId":"","cores":2,"modelName":"Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz","mhz":1800,"cacheSize":256,"flags":["fpu","vme","de","pse","tsc","msr","pae","mce","cx8","apic","sep","mtrr","pge","mca","cmov","pat","pse36","clfsh","ds","acpi","mmx","fxsr","sse","sse2","ss","htt","tm","pbe","sse3","pclmulqdq","dtes64","mon","dscpl","vmx","smx","est","tm2","ssse3","fma","cx16","tpr","pdcm","sse4.1","sse4.2","x2apic","movbe","popcnt","aes","pcid","xsave","osxsave","seglim64","tsctmr","avx1.0","rdrand","f16c","rdwrfsgs","tsc_thread_offset","bmi1","hle","avx2","smep","bmi2","erms","invpcid","rtm","fpu_csds","rdseed","adx","smap","ipt","mdclear","ibrs","stibp","l1df","ssbd","syscall","xd","1gbpage","em64t","lahf","lzcnt","prefetchw","rdtscp","tsci"],"microcode":""}]
2cpu1: [{"cpu":"cpu-total","user":141059.1,"system":180386.0,"idle":685311.2,"nice":0.0,"iowait":0.0,"irq":0.0,"softirq":0.0,"steal":0.0,"guest":0.0,"guestNice":0.0}]
3[22.71604938271605]
```

Copy

### 采集内存信息

```go
 1package main
 2
 3import (
 4	"fmt"
 5	"github.com/shirou/gopsutil/mem"
 6)
 7
 8func main() {
 9	//获取物理内存和交换区内存信息
10	m1, _ := mem.VirtualMemory()
11	fmt.Println("m1:", m1)
12	m2, _ := mem.SwapMemory()
13	fmt.Println("m2:", m2)
14}
```

Copy

运行后输出结果如下

```bash
1m1: {"total":8589934592,"available":1712226304,"used":6877708288,"usedPercent":80.06706237792969,"free":102297600,"active":1632382976,"inactive":1609928704,"wired":4049559552,"laundry":0,"buffers":0,"cached":0,"writeback":0,"dirty":0,"writebacktmp":0,"shared":0,"slab":0,"sreclaimable":0,"sunreclaim":0,"pagetables":0,"swapcached":0,"commitlimit":0,"committedas":0,"hightotal":0,"highfree":0,"lowtotal":0,"lowfree":0,"swaptotal":0,"swapfree":0,"mapped":0,"vmalloctotal":0,"vmallocused":0,"vmallocchunk":0,"hugepagestotal":0,"hugepagesfree":0,"hugepagesize":0}
2m2: {"total":7516192768,"used":6074400768,"free":1441792000,"usedPercent":80.81752232142857,"sin":0,"sout":0,"pgin":0,"pgout":0,"pgfault":0}
```

Copy

### 采集主机信息

```go
 1package main
 2
 3import (
 4    "fmt"
 5    "github.com/shirou/gopsutil/host"
 6)
 7
 8func main() {
 9  h,_ := host.Info()
10  fmt.Println("本机信息：",h)
11}
```

Copy

运行结果如下

```bash
1本机信息： {"hostname":"huangzhgdedeAir.lan","uptime":714985,"bootTime":1582686932,"procs":469,"os":"darwin","platform":"darwin","platformFamily":"Standalone Workstation","platformVersion":"10.15.3","kernelVersion":"19.3.0","kernelArch":"x86_64","virtualizationSystem":"","virtualizationRole":"","hostid":"a8dde75c-cd97-3c37-b35d-1070cc50d2ce"}
```

Copy

### 采集磁盘信息

```go
 1package main
 2
 3import(
 4	"fmt"
 5	"github.com/shirou/gopsutil/disk"
 6)
 7
 8func main(){
 9	 //可以通过psutil获取磁盘分区、磁盘使用率和磁盘IO信息
10	 d1, _ := disk.Partitions(true)  //所有分区
11	 fmt.Println("d1:",d1)
12	 d2, _ := disk.Usage("E:")  //指定某路径的硬盘使用情况
13	 fmt.Println("d2:",d2)
14	 d3, _ := disk.IOCounters()  //所有硬盘的io信息
15	 fmt.Println("d3:",d3)
16	}
```

Copy

运行结果输出如下

```bash
1d1: [{"device":"/dev/disk1s5","mountpoint":"/","fstype":"apfs","opts":"ro,journaled,multilabel"} {"device":"devfs","mountpoint":"/dev","fstype":"devfs","opts":"rw,nobrowse,multilabel"} {"device":"/dev/disk1s1","mountpoint":"/System/Volumes/Data","fstype":"apfs","opts":"rw,nobrowse,journaled,multilabel"} {"device":"/dev/disk1s4","mountpoint":"/private/var/vm","fstype":"apfs","opts":"rw,nobrowse,journaled,multilabel"} {"device":"map auto_home","mountpoint":"/System/Volumes/Data/home","fstype":"autofs","opts":"rw,nobrowse,automounted,multilabel"} {"device":"/dev/disk1s3","mountpoint":"/Volumes/Recovery","fstype":"apfs","opts":"rw,nobrowse,journaled,multilabel"}]
2d2: <nil>
3d3: map[disk0:{"readCount":29285277,"mergedReadCount":0,"writeCount":13452048,"mergedWriteCount":0,"readBytes":784495985152,"writeBytes":503051173888,"readTime":16338850,"writeTime":6971027,"iopsInProgress":0,"ioTime":23309878,"weightedIO":0,"name":"disk0","serialNumber":"","label":""}]
```

Copy

### 采集网络信息

```go
 1package main
 2
 3import (
 4	"fmt"
 5
 6	"github.com/shirou/gopsutil/net"
 7)
 8
 9func main() {
10	//获取当前网络连接信息
11	n1, _ := net.Connections("all") //可填入tcp、udp、tcp4、udp4等等
12	fmt.Println("n1:", n1)
13
14	//获取网络读写字节／包的个数
15	n2, _ := net.IOCounters(false)
16	fmt.Println("n2:", n2)
17}
```

Copy

输出结果如下

```bash
1n1: [{"fd":4,"family":2,"type":1,"localaddr":{"ip":"*","port":58299},"remoteaddr":{"ip":"","port":0},"status":"LISTEN","uids":null,"pid":560} {"fd":5,"family":30,"type":1,"localaddr":{"ip":"*","port":58299},"remoteaddr":{"ip":"","port":0},"status":"LISTEN","uids":null,"pid":560} {"fd":12,"family":2,"type":1,"localaddr":{"ip":"192.168.10.117","port":58299},"remoteaddr":{"ip":"192.168.10.139","port":54623},"status":"ESTABLISHED","uids":null,"pid":560} {"fd":13,"family":2,"type":2,"localaddr":{"ip":"*","port":3722},"remoteaddr":{"ip":"","port":0},"status":"","uids":null,"pid":560} {"fd":15,"family":2,"type":1,"localaddr":{"ip":"192.168.10.117","port":58299},"remoteaddr":
2...
3n2: [{"name":"all","bytesSent":7849600010,"bytesRecv":13449618362,"packetsSent":13952479,"packetsRecv":14294339,"errin":0,"errout":0,"dropin":0,"dropout":636,"fifoin":0,"fifoout":0}]
```

Copy

### 采集进程相关信息

```go
 1package main
 2
 3import (
 4	"fmt"
 5
 6	"github.com/shirou/gopsutil/process"
 7)
 8
 9func main() {
10	//获取到所有进程的详细信息
11	p1, _ := process.Pids() //获取当前所有进程的pid
12	fmt.Println("p1:", p1)
13	ifExists, _ := process.PidExists(10086) // 判断进程是否存在
14	fmt.Println("ifExists:", ifExists)
15
16}
```

Copy

输出结果

```bash
1p1: [1 125 126 129 130 131 132 135 138 139 141 146 151 155 156 157 164 165 166 167 168 170 171 172 174 177 178 180 181 186 187 188 191 192 194 195 197 198 200 201 202 203 204 206 207 209 215 222 230 239 247 248 254 257 258 259 260 286 384 385 391 414 444 446 447 449 451 471 483 493 495 502 516 519 525 528 533 534 535 536 538 541 542 546 548 549 552 555 558 559 560 561 562 563 565 566 568 571 572 574 576 577 578 580 581 582 583 586 590 595 597 600 601 621 622 624 625 628 629 632 633 639 641 642 643 644 652 654 658 659 660 661 662 665 667 670 671 682 685 686 688 689 691 696 697 698 713 718 719 720 737 793 802 1277 1278 1279 1282 1283 1286 1287 1313 1315 1379 1381 ...]
2ifExists: false
```

Copy

## 应用领域

### 服务器性能监控

写一个程序，周期性的获取服务器的各项指标数据，存入数据库中，然后通过界面进行查询，方便排查故障。

- **原文作者：**[黄忠德](https://huangzhongde.cn/)
- **原文链接：**https://huangzhongde.cn/post/2020-03-06-golang_introduce_gopsutil/