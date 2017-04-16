package tasks

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON handles unmarshaling null, a single string, or an
// array of strings to OptionalStringArray.
func (sa *OptionalStringArray) UnmarshalJSON(data []byte) error {
	switch data[0] {
	case 'n':
		if string(data) != "null" {
			return fmt.Errorf("unexpected data %v", data[1])
		}
		*sa = nil
	case '[':
		val := make([]string, 0)
		err := json.Unmarshal(data, &val)
		if err != nil {
			return err
		}
		*sa = OptionalStringArray(val)
	case '"':
		val := ""
		err := json.Unmarshal(data, &val)
		if err != nil {
			return err
		}
		*sa = OptionalStringArray([]string{val})
	default:
		return fmt.Errorf("unexpected data %v", data[0])
	}

	return nil
}
