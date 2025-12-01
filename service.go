package main

type Service interface {
	SayHello(name string) string
}

type ServiceImplementation struct{}

func (ServiceImplementation) SayHello(name string) string {
	return "Hello, " + name + "!"
}
