# util
Utility lib for go 

##go get

`go get github.com/mossila/goutil`

## Readlns 
usage 

```go
package main

import (
	"fmt"

	"github.com/mossila/goutil"
)

func main() {
	for s := range util.Readlns("README.md") {
		fmt.Println(s)
	}
}
```
