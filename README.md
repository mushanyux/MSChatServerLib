# MSChatServerLib

一些通用的库

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

### log
日志管理，使用Uber开源的zap库

#### logger
多种级别的日志输出

#### options
日志配置

### redis
Redis相关的操作，比如 Set, Del, Hget 等。

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