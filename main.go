package main

import (
	"fmt"

	"github.com/hoshina85/test-ci/lib"
	"github.com/serenize/snaker"
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
