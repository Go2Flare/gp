package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	helloworld "go_code/gp/06proto_helloworld/proto/hello"
)

type Hello struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Courses []string `json:"courses"`
}

func main(){
	testStruct := Hello{Name:"flare",Age:55,Courses:[]string{"Math","English"}}
	jsonRsp, _ := json.Marshal(testStruct)
	fmt.Println(string(jsonRsp),len(jsonRsp))
	req := helloworld.HelloRequest{
		Name:"flare",
		Age: 55,
		Courses: []string{"Math","English"},
	}
	//proto能序列化的对象都是定义在proto文件中
	rsp,_ := proto.Marshal(&req)//编码原理varint
	fmt.Println(string(rsp), len(rsp))

	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp,&newReq)//将rsp信息反序列化
	fmt.Println(newReq.Name,newReq.Age,newReq.Courses)
}
