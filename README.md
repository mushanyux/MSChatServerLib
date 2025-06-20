# MSChatServerLib

一些通用的库

## common

### cache
对 pkg/cache 的封装

### constant
常量定义，包括群，频道，社区等定义

### msg
消息相关的常量定义，根据两个id生成虚拟id和还原

### page
页格式的封装

## config

### asynctask
主要是对machinery/v2的封装, 支持创建，注册，发送任务

### config
各种配置，比如缓存、文件、推送、系统账户等

### context
事件上下文封装，在线状态封装，消息监听相关

### elastic
获取elastic客户端

### msg
消息，好友处理，频道管理等类型和方法的定义

### msg_channel
更新频道

### msg_group
群相关操作，主要是会在群中产生消息的操作

### msg_rtc
通话相关

### seq
生成序列号，包括当前序号和最大序号。步长1000

### stream
开启流和结束流

### tracer
主要是对jaeger库的封装，以及空标签的定义

## model

### channel
响应数据的模型定义

### respond
一些接口的模型定义

## pkg

### cache

#### cache
缓存接口

### db
数据库连接

#### db
一些通用函数，比如自定义时间类型的序列化和反序列化，用于描述创建时间和修改时间

#### mysql
新建会话，数据库迁移，sql文件查找

#### redis
新建连接

#### sqlite
新建会话并迁移

### keylock
关键字锁

### log
日志管理，使用Uber开源的zap库

#### logger
多种级别的日志输出

#### options
日志配置

### redis
Redis相关的操作，比如 Set, Del, Hget 等。

### markdown
将markdown转换为html。

### msevent
事件定义

### mshook
grpc生成的代码

### mshttp
主要是对 gin 的封装，还包括一些认证、封装

### msrsa
用 rsa 密钥生成 md5 签名

### network
get, put, post 等操作

### pool
池，实现了工作者，任务队列，分发器

#### dispatcher
分发器内部使用 ans.Pool, 负载过高时输出提示

#### queue
任务队列，使用锁和条件变量实现 push, pop 等操作

#### worker
任务接口，工作者接口，开始、结束工作

### register
注册模块，路由定义

### util
字符串，时间，加密等基础工具

#### aes
aes加密。现代密码学似乎都在往公开的方向发展，大抵是因为目标从无法攻克转移到提升攻克难度。

aes的保障大概是因为域之类的东西，当然可以在程序员的视角，简单的看成运算符重载。重载后的异或+条件判断带来多种可能。

大致流程为初始变换、循环运算，最终轮。循环运算包含字节代换，行移位，轮密钥加。当然实际操作过程还要考虑填充等细节。

详细的移步博客（详细文章待施工）

#### base62
十进制转62进制

#### common
substr, MapToQueryParamSort, Sign

#### decimal
精确的小数计算。

四舍五入，银行家舍入等多种舍入。

序列化反序列化。

要看懂得复习一下数电的IEEE 754标准。

#### dh
Curve25519算法，生成密钥对、计算加密密钥。

#### hash
使用CRC32通过一个字符串生成一个32位的整数。

#### ip
查询服务器、客户端ip地址。根据ip查询地址（使用高德api）。

#### json
obj和json，json和map的转换。

#### md5
MD5加密，SHA1加密，使用sha1作为哈希的HMAC加密。

#### page
页结构体

#### qreflect
通过反射获取所有属性，并以下划线形式返回

#### sha
sha256加密

#### sign
为一个map生成签名。

#### string
字符串处理，包括：

驼峰转下划线，下划线转驼峰，获取uuid，随机生成

#### stringbuffer
实现Append，传入interface{}，写入到Buffer

#### time
时间格式化处理。

#### uuid
uuid相关。使用的是 Maxim Bublis 开源的 uuid 库。

#### yuan_cent
调用decimal实现元-分互转

### wait
等待-触发机制，可以注册id，向所有注册了的发数据

## server
启动server，http，https服务等

## testutil
测试