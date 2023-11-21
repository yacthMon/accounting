package helper

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(result interface{}) {
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", output)
	}
}