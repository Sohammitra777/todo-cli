package utils

import (
	"encoding/json"
	"os"
)

func SaveToJson(filename string, v any) {
	data, err := json.MarshalIndent(v, "", "    ")
	PrintErr(err)

	err = os.WriteFile(filename, data, 0644)
	PrintErr(err)
}
