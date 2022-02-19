package rpc

import (
	"fmt"
)

type ss struct{
	name string
	address string
	ot aa
}
type aa struct{
	a1 string
	a2 string
}
type PrintResult struct {
	Info string
	Err error
}

func RPCPrintln(ss1 ss) {
	/*
	客户端：
		1. 建立连接 tcp/http
		2. 将ss对象序列化为json - 序列化
		3．发送json字符串–调用成功后实际上你接收到的是一个二进制的数据
		4. 等待服务器发送结果
		5. 将服务器返回的数据解析成PrintResult对象 - 反序列化
	服务端：
		1. 监听网络端口 80
		2. 读取数据 - 二进制的json数据
		3. 对数据进行反序列ss对象 - 反序列化
		4. 开始处理业务逻辑
		5. 将处理的结束PrintResult序列化成json二进制数据 - 序列化
		6. 将数据返回
	*/
}



func main(){
	/*
	电商系统逻辑：减库存，库存服务是一个独立的系统，reduce
	必须会牵扯到网络，做成web服务（gin,beego,net/httpserver)
	*/

	//1.这个函数的调用参数如何传递-json/xml/protobuf/msgpack(数据编码协议)，但json性能不高
	//网络有两个段，客户端：将数据传输到gin
	//gin-服务端：负责解析数据
	//fmt.Println(Add(1,3))
	fmt.Println(
		ss{
			"yoyo",
			"beijing",
			aa{
				"my name is a1",
				"this is a2",
			},
		})
}