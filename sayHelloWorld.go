package testmodulego

import "fmt"

func SayHelloWorld() string {
	return "Hello World"
}

func JustHello(name string) string { 
	return fmt.Sprintf("Hello, %s", name)
}