package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Proto2Map(a any) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return nil, errors.New("Failed to marshal message: " + err.Error())
	}
	var m map[string]interface{}
	err = json.Unmarshal(jsonBytes, &m)
	if err != nil {
		return nil, errors.New("Failed to unmarshal message: " + err.Error())
	}

	return m, nil
}

func GenImageID() *uint32 {
	rand.NewSource(time.Now().UnixNano())
	num := uint32(rand.Intn(900000) + 100000)
	return &num
}
