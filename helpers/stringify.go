package helpers

import (
	"encoding/json"
)

func Stringify(arg any) string {
	bytes, _ := json.MarshalIndent(arg, "", "\t")
	return string(bytes)
}
