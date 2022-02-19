package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{

}

//远端处理逻辑函数
func (s *HelloService)Hello(request string, reply *string) error{
	//返回值通过修改reply的值
	*reply = "Hello, "+request
	return nil
}

func main(){
	//1.实例化一个server
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/httprpc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		var conn io.ReadWriteCloser = struct{//读写请求
			io.Writer
			io.ReadCloser
		}{
			Writer: w,
			ReadCloser: r.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))//conn使用json编码
		fmt.Fprintf(w,"r.Body:%v\n\nr.Header:\n%v\n\nr.URL:%v\n\nr.Write(w):%v\n\nr.Response:%v\n\nr.Form:%v\n\n",
			r.Body,r.Header,r.URL,r.Write(w),r.Response,r.Form)
	})
	http.ListenAndServe(":8887", nil)
}