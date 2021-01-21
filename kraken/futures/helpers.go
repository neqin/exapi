package futures

import (
	"encoding/json"
)

func Marshal(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		return "{marshal error}"
	}
	return string(res)
}
