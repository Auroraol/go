[【数据库】——MongoDB常用语法](https://blog.csdn.net/weixin_44697562/article/details/110122105)

http://try.mongodb.org/







# 查询数量

```
db.sdk_goods_pdd.count({"platform": "pdd", "plat_shop_id":"468582923"})

```





```
Opt>

没有资产

Opt> r

                 刘丰洁, 欢迎使用Jumpserver开源跳板机系统

        1) 输入 ID 直接登录 或 输入部分 IP,主机名,备注 进行搜索登录(如果唯一).
        2) 输入 / + IP, 主机名 or 备注 搜索. 如: /ip
        3) 输入 p 显示您有权限的主机.
        4) 输入 g 显示您有权限的节点.
        5) 输入 g + 节点ID 显示节点下主机. 如: g1
        6) 输入 s 中/英文切换.
        7) 输入 h 帮助.
        0) 输入 r 刷新最新的机器和节点信息.
        0) 输入 q 退出.

Opt>

没有资产

Opt> r

                 刘丰洁, 欢迎使用Jumpserver开源跳板机系统

        1) 输入 ID 直接登录 或 输入部分 IP,主机名,备注 进行搜索登录(如果唯一).
        2) 输入 / + IP, 主机名 or 备注 搜索. 如: /ip
        3) 输入 p 显示您有权限的主机.
        4) 输入 g 显示您有权限的节点.
        5) 输入 g + 节点ID 显示节点下主机. 如: g1
        6) 输入 s 中/英文切换.
        7) 输入 h 帮助.
        0) 输入 r 刷新最新的机器和节点信息.
        0) 输入 q 退出.

Opt> r

                 刘丰洁, 欢迎使用Jumpserver开源跳板机系统

        1) 输入 ID 直接登录 或 输入部分 IP,主机名,备注 进行搜索登录(如果唯一).
        2) 输入 / + IP, 主机名 or 备注 搜索. 如: /ip
        3) 输入 p 显示您有权限的主机.
        4) 输入 g 显示您有权限的节点.
        5) 输入 g + 节点ID 显示节点下主机. 如: g1
        6) 输入 s 中/英文切换.
        7) 输入 h 帮助.
        0) 输入 r 刷新最新的机器和节点信息.
        0) 输入 q 退出.

Opt> [A

没有资产

Opt>
ID   主机名                     IP               登录用户 备注
1    cd-032114-lane-mini-db-all 10.248.32.114    [worker]
2    cd-033114-lane-tb-db-all   10.248.33.114    [worker]

页码: 1, 数量: 2, 总页数: 1, 总数量: 2

Opt> 2

开始连接到 worker@cd-033114-lane-tb-db-all 1.4
===================================================================
 OS: CentOS Linux release 7.6.1810 (Core)
 UP: 6 days 23 hours 33 minutes 34 seconds
 System load:   3.29            Memory usage:   67.0%
 Usage on /:    58%             Swap usage:
 Login users:   2               Processes:      530
===================================================================

          _                    _                       _
         (_)                  | |                     (_)
   __  __ _   __ _   ___    __| | _   _   ___    __ _  _
   \ \/ /| | / _` | / _ \  / _` || | | | / _ \  / _` || |
    >  < | || (_| || (_) || (_| || |_| || (_) || (_| || |
   /_/\_\|_| \__,_| \___/  \__,_| \__,_| \___/  \__,_||_|

[worker@cd-033114-lane-tb-db-all ~]$  sudo docker exec -it mongodb-online-27017 bash
root@9fc97df81141:/#
root@9fc97df81141:/#
root@9fc97df81141:/#
root@9fc97df81141:/#
root@9fc97df81141:/#
root@9fc97df81141:/#
root@9fc97df81141:/# mongo  -u root -p 3SqzSt65  --authenticationDatabase admin
MongoDB shell version v4.2.2
connecting to: mongodb://127.0.0.1:27017/?authSource=admin&compressors=disabled&gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("cd52ba52-543b-4d45-87e3-6e3816ced627") }
MongoDB server version: 4.2.2
Server has startup warnings:
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten]
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten]
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten]
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten] ** WARNING: /sys/kernel/mm/transparent_hugepage/enabled is 'always'.
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten] **        We suggest setting it to 'never'
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten]
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten] ** WARNING: /sys/kernel/mm/transparent_hugepage/defrag is 'always'.
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten] **        We suggest setting it to 'never'
2024-09-19T09:14:36.609+0000 I  CONTROL  [initandlisten]
---
Enable MongoDB's free cloud-based monitoring service, which will then receive and display
metrics about your deployment (disk utilization, CPU, operation statistics, etc).

The monitoring data will be available on a MongoDB website with a unique URL accessible to you
and anyone you share the URL with. MongoDB may use this information to make product
improvements and to suggest MongoDB products and deployment options to you.

To enable free monitoring, run the following command: db.enableFreeMonitoring()
To permanently disable this reminder, run the following command: db.disableFreeMonitoring()
---

xd-online:PRIMARY> use xdmp
switched to db xdmp
xd-online:PRIMARY> db.smkd_whitelist.insert( { "_id" : ObjectId("66e2596a87dd113f0421fcf4"), "shop_id" : ObjectId("66e25725cb9a2ff7c5869428") })
WriteResult({ "nInserted" : 1 })
xd-online:PRIMARY>
```

