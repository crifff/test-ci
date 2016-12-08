package main

import (
	"fmt"

	"github.com/crifff/test-ci/lib"
	"github.com/serenize/snaker"
)

func main() {
	s, err := hello()
	if err != nil {
		return
	}
	fmt.Println(s)
}

func hello() (string, error) {
	return snaker.CamelToSnake(lib.GetStr()), nil
}
