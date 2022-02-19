package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main(){
	//http://127.0.0.1:8000/add?a=1&b=2
	//http://127.0.0.1:8000?method=add&a=1&b=2
	//返回的格式： json{"data":3}
	//1. callID问题： r.URL.Path, 2.数据的传输协议，url参数传递协议
	//3. 网络传输协议TCP协议
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){
		_ = r.ParseForm()

		fmt.Println("path: ", r.URL.Path)
		a,_ := strconv.Atoi(r.Form["a"][0])
		b,_ := strconv.Atoi(r.Form["b"][0])

		w.Header().Set("Content-Type", "application/json")//请求头设置
		jD, _ := json.Marshal(map[string]int{
			"data":a+b,
		})
		_,_ = w.Write(jD)
	})
	http.ListenAndServe("localhost:8000", nil)
}
