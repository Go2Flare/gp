package main


import(
	"encoding/json"
	"fmt"
	HttpReq "github.com/kirinlabs/HttpRequest"
)
type repData struct{
	Data int `json:"data"`
}

//如同调用本地函数
func Add(a,b int){
	//传输协议http
	//http请求
	req := HttpReq.NewRequest()
	//http response
	res,_ := req.Get(fmt.Sprintf("http://localhost:8000/%s?a=%d&b=%d","add",a,b))
	//http请求体
	body,_ := res.Body()
	fmt.Println(string(body))
	repD := repData{}
	_ =json.Unmarshal(body, &repD)
	fmt.Println(repD.Data)
}
func main(){
	Add(3,5)
}
