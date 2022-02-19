package handler

//server，client调用函数名称问题
const HelloServiceName = "handler/HelloService"//路径作为名称

//server中函数的传输对象HelloService更改了
type NewHelloService struct{}
//我们关心的是这个结构体的名称还是这个结构体内的方法？别的包肯定关心方法

func (s *NewHelloService)Hello(request string, reply *string) error{
	*reply = "Hello, "+request
	return nil
}