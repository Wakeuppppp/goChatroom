# GO note
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.io,direct
#### GoLand 配置代理 File -> settings -> Go -> Go Modules -> environment
    GOPROXY=https://goproxy.cn,direct
# 接口
一个接口可以继承多个接口，但是在这多个接口中，不能有相同的方法
```go
type A interface{
    Test01()
    Test02()
}

type B interface{
    Test01()
    Test03()
}
// C继承了A和B,但是A和B中有相同的方法Test01(),编译器会报错
type C interface{
    A
    B
}
```
关于接口和继承的疑问：
    A实现了接口M的方法,B继承了A,那么B是否也继承了接口M,或者需要实现接口M
```go
type M interface {
	Mfunc()
}

type A struct {
}

func (A) Mfunc() {
	fmt.Println("A.Mfunc()")
}

type B struct {
	A
}

func (B) Mfunc() {
    fmt.Println("B.Mfunc()")
}

func main() {
	var b B
	var m M = b
	m.Mfunc()
}


// output：B.Mfunc()
// 若注释掉B实现接口M的方法,则
// output: A.Mfunc()
```
需求分析->设计阶段->实现阶段->测试阶段->实施阶段->维护阶段


类型断言





# Redis

## Go安装Redis

cd 到Go  PATH目录下，执行 “go get github.com/gomodule/redigo/redis ”

## 基本使用

默认有16个数据库 初始默认使用0号库 0~15

1. 添加key-val	[set]
2. 查看当前redis的所有key  [keys *]
3. 获取key对应的值  [get key]
4. 切换redis数据库  [select index]
5. 查看当前数据库key-val的数量  [dbsize]
6. 清空当前数据库的key-val和清空所有数据库的key-val  [flushdb  flushall]

## 五大类型

### String

是二进制安全的，除普通的字符串外，也可以存放图片等数据

字符串value最大是512M

set key val

get key

del key

setex key seconds val(seconds之后该key-val失效)

mset key val key val...(同时设置多对key-val)

mget key key...(同时获取多对key-val)

![](E:\GoProjects\src\png\微信截图_20220324105504.png)

### Hash

键值对的集合，特别适合存储对象

hset key field val

hget key field

hgetall key 会获取key所有的field-val

hdel key field 删除key的field-val

hmget key field field...获取多个val

hmset key field val field val...给key同时设置多个field-val

hexists key field 查看哈希表key中，给定域中field是否存在

![](E:\GoProjects\src\png\微信截图_20220324105757.png)

### List

列表是简单的字符串列表，按照插入顺序排序，添加时可以指定头部或尾部

本质是一个链表，元素是有序的，可以重复

![](E:\GoProjects\src\png\微信截图_20220324105813.png)

### Set

底层是HashTable数据结构，是string类型的**无序**集合，且元素的值不能重复

![](E:\GoProjects\src\png\微信截图_20220324105805.png)

### zset

## Redis连接池

