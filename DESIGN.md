# service registry 设计

## 0. 说明

service registry是用来管理各种抽象的资源，比如：机器、监控采集项、报警策略等

## 1. 实现目标

* 存储各种资源
* Restful API
* agent启动后，根据预定规则，服务自发现，自动注册到registry
* 底层存储高可用
* 实现cache，抵抗大量读取请求
* 节点变更通知（服务注册、服务注销）
* 全网服务器心跳上报检测 （自动注销）
* 非leader节点不处理写请求，提供API接口来查询当前node是否是leader，后期可以做无需人工干预的读写分离和自动切换。
* 读性能要求低延迟
* backup restore
* 可以从任意节点进行读写，集群模式对前端透明

## 2. 技术选型

```

语言      ：golang
web框架   ：http
存储      ：bolt
分布式算法 ：raft

```

## 3. 架构

```

               +------------+
               |            |
               |    user    |
               |            |
               +------------+

                     +   read  : api.registry.test.com
                     |   write : api.registry.test.com
                     v

   +------------------------------------+
   |                                    |
   |                LB                  |
   |                                    |
   +------------------------------------+

      +                 +               +
      |                 |               |
      v  write & read   v write & read  v write & read

+-------------+  +-------------+  +-------------+
|             |  |             |  |             |
|    Leader   |  |  Follower   |  |  Follower   |
|             |  |             |  |             |
+-------------+  +-------------+  +-------------+

```


## 4. 模型

* 一个bucket对应一个节点
* 节点的属性和资源以key-value存储其中
* 节点的层级关系存储在某一个节点中
* 每个节点都有唯一ID
* 每个key对应一种资源类型
* []map[string]string 来存放该类资源下的所有数据
* 数据存储前的编码自己完成，对搜索优化


```

                                              +---------------+
                                              |               |
                                 +------------+ router (key)  |------------ []map[string]string    (value)
                                 |            |               |
                                 |            +---------------+
                                 |
                                 |
+----------------------+         |
|                      |         |
|                      |         |           +----------------+
|                      |         |           |                |
|      Bucket (ns)     +---------------------+     Metrics    |
|                      |         |           |                |
|                      |         |           +----------------+
|                      |         |
+----------------------+         |
                                 |
                                 |
                                 |           +----------------+
                                 +-----------+                |
                                 |           |    alert       |
                                 |           |                |
                                 |           +----------------+
                                 |
                                 |
                                 |
                                 |
                                 |           +----------------+
                                 +-----------+                |
                                             |    meachine    |
                                             |                |
                                             +----------------+
                                 



```



用一个特殊的bucket来存储全局的东西，比如节点的关联关系，用这个只存了`bucket ID`的map来关联所有的bucket.


```


                                                                        +-----------+
                                                                        |           |
                                                                +-------+    node10 |
                                      +-------------+           |       |           |
                           +----------+             |           |       +-----------+
                           |          |    node0    | +---------+
                           |          |             |           |
                           |          +-------------+           |       +-----------+
+-------------+            |                                    |       |           |
|             |            |                                    +-------+    node11 |
|    map      | +--------> |                                            |           |
|             |            |                                            +-----------+
+-------------+            |
                           |
                           |
                           |
                           |          +-------------+
                           |          |             |
                           +----------+    node1    |
                                      |             |
                                      +-------------+
                        



```
