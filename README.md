# util
Utility lib for go 

##go get

`go get github.com/mossila/util`

## Readlns 
usage 

```go
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
```
