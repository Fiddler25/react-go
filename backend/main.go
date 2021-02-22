package main

import "fmt"

func GetHello(s string) string {
	return "Hello " + s + "!"
}

func main() {
	s := GetHello("World")
	fmt.Println(s)
}
