package goutil

import (
	"encoding/json"
	"fmt"
)

func printJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
