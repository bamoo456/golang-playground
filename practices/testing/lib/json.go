package lib

import (
	"errors"
)

func ToJSONArray(raw interface{}) ([]map[string]interface{}, error) {
	arr, ok := raw.([]interface{})

	if !ok {
		return []map[string]interface{}{}, errors.New("Parse error")
	}
	res := make([]map[string]interface{}, len(arr))
	for i, v := range arr {
		arr[i], ok = v.(map[string]interface{})
		if !ok {
			return res, errors.New("Parse error in array element ")
		}
	}

	return res, nil
}
