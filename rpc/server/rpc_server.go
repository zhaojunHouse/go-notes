package server

// HelloService rpc
type HelloService struct {
}

// Hello rpc method
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
