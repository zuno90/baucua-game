package utils

import "encoding/json"

func MarshalBinary(v any) ([]byte, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
