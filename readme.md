## go-rabbitmq

封装go 操作rabbitmq
- 支持confirm机制 
- 基于延迟插件的实现延迟队列生产函数
- 支持连接池，针对高并发请求连接数不够的情况



### 安装

```shell
go get github.com/cr-mao/go-rabbitmq@v1.1.0
```



### demo
见test目录下代码 



### 测试
```test
# 生产者测试
go test -v test/pub_test.go
# 消费者测试
go test -v test/sub_test.go

# 连接池测试
go test -v test/pool_test.go
```


### 其他说明
```text
lib.go 所有的操作都封装在里头在里头做扩展你想的要的功能即可
pub.go 生产函数
sub.go 消费函数
```


