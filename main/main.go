package main

import (
	"fmt"

	"github.com/mossila/util"
)

func main() {
	for s := range util.Readlns("README.md") {
		fmt.Println(s)
	}
}
