package main

import (
	"fmt"

	"github.com/mossila/goutil"
)

func main() {
	for s := range goutil.Readlns("README.md") {
		fmt.Println(s)
	}
}
