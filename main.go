package main

import (
	"fmt"

	"github.com/serenize/snaker"
	"test-ci/lib"
)

func main() {
	var s string

	if s, _ := hello(); s != "" {
		return
	}
	fmt.Println(s)
}

func hello() (string, error) {
	return snaker.CamelToSnake(lib.GetStr()), nil
}
