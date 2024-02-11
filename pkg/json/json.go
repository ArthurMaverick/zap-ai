package json

import (
	"encoding/json"
)

// Response is the structure of a JSON response
type Response struct {
	StatusCode int         `json:"status_code"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// Stringify converts a payload to a JSON string
func Stringify(payload interface{}) ([]byte, error) {
	response, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Parse converts a JSON string to a payload
func Parse(payload []byte) (Response, error) {
	var jsonResponse Response
	err := json.Unmarshal(payload, &jsonResponse)
	if err != nil {
		return Response{}, err
	}
	return jsonResponse, nil
}
