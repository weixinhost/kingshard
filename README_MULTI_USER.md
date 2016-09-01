### 多用户支持说明
-------

##### 为什么需要支持多用户

* 在语义上支持了Mysql的权限系统。
* 为后续做用户级别的资源隔离做基础。
* 在使用场景中需要支持多用户的。
	1. 旧系统迁移（旧系统中已经使用了权限管理的）
	2. 多租户系统的实现

##### 实现细节

* 在原数据库池的实现上包裹出一层UserPool来支持多用户（即每个用户一个连接池）。

* 所有的管理功能支持都将随机抽取一个user来进行操作（未来区分管理帐号与普通帐号）。

* 所有的数据操作都将使用连接kingshard的用户作为mysql的用户进行后端MYSQL操作

* 完全不破坏原有逻辑与结构，某些细节可能有轻微的改动


##### 配置说明

为了实现该特性，对配置文件做了不兼容的重构。
重构部分如下：

```yaml
addr : 0.0.0.0:9696
# kingshard用户列表
users:
-
 user: root
 password: root

-
 user: user_1
 password: pass_1

-
 user: user_2
 password: pass_2

-
 user: user_3
 password: pass_3

-
 user: user_4
 password: pass_4

log_level : debug
#log_sql: off
proxy_charset: utf8
#log_memory 当query的数据集超过该值时，则输出到日志。 0 表示关闭，单位为byte
log_memory: 0
#max_memory 当query的数据超过该值时，则报错（Mysql协议）。0 表示关闭，单位为byte
max_memory: 0

nodes :
-
    name : node1
    max_conns_limit : 5
    # 数据库用户列表,数据库用户列表
    # 需要与Kingshard的用户列表一一对应
    # 只需要用户名一一对应，密码可以不同
    users:
    -
      user: root
      password: root

    -
      user: user_1
      password: pass_1

    -
      user: user_2
      password: pass_2

    -
      user: user_3
      password: pass_3

    -
      user: user_4
      password: pass_4

    master : 10.211.55.3:3306

schema :
    db :
    nodes: [node1]
    default: node1
    shard:
    -
```


#### 多用户支持带来的潜在问题

1. 按照目前实现，一个user就对应一个独立的mysql池。user的数量将对性能有影响。如果该问题严重，未来版本中将重构连接池的实现，使用一个带用户纬度的连接池来解决该问题。
2. 真实的IdelConnection的数量将等于配置文件中的数量 * 用户数

#### 性能测试

暂无

