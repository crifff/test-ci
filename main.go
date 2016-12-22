package main

import (
	"github.com/crifff/test-ci/lib"
	"github.com/k0kubun/pp"
	"github.com/serenize/snaker"
)

func main() {
	s, err := hello()
	if err != nil {
		return
	}
	pp.Print(s)
}

func hello() (string, error) {
	return snaker.CamelToSnake(lib.GetStr()), nil
}
