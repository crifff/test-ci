package main

import "fmt"

func main() {
	var s string

	if s, _ := hoge(); s != "" {
		return
	}
	fmt.Println(s)
}

func hoge() (string, error) {
	return "hello", nil
}
