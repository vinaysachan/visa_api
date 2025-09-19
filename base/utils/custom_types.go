package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type StringOrNumber string

func (s *StringOrNumber) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*s = StringOrNumber(str)
		return nil
	}

	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*s = StringOrNumber(fmt.Sprintf("%.0f", num)) // no decimal
		return nil
	}

	return fmt.Errorf("invalid type for mobile_number")
}

func (s StringOrNumber) ToUint64() (uint64, error) {
	return strconv.ParseUint(string(s), 10, 64)
}
