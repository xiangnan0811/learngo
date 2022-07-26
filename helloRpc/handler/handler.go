package handler

const HelloServiceName = "handler/HelloService"

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
