package goutil

import (
	"encoding/json"
	"fmt"
)

//PrintJSON pring as json result
func PrintJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
