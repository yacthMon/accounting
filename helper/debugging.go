package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func PrintJSON(result interface{}) {
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", output)
	}
}

func ConvertObjToStringReader[T any](obj T) io.Reader {
 result, _ := json.Marshal(&obj)
 return strings.NewReader(string(result))
}