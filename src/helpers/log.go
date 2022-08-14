package helpers

import (
	"encoding/json"
	"fmt"
)

func Log(object interface{}) {
	json, _ := json.MarshalIndent(object, "", "\t")
	fmt.Println(string(json))
}
