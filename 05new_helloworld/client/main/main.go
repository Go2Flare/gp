package main

import (
	"fmt"
	proxy "go_code/gp/05new_helloworld/client_proxy"
	_ "go_code/gp/05new_helloworld/handler"
	"log"
)

func main() {
	//conn, err := rpc.Dial("tcp", "localhost:8887")//原先go的协议是Gob
	//if err != nil{
	//	panic("连接失败")
	//}
	client := proxy.NewHelloServiceClient("tcp", "localhost:8887")
	var reply string
	//err = conn.Call(handler.HelloServiceName+".Hello", "flare", &reply)
	//if err != nil{
	//	panic("调用失败")
	//}
	err := client.Hello("flare", &reply)
	if err != nil {
		log.Fatal("err:",err)
	}
	fmt.Println(reply)
}
