package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// UnmarshalJSON accepts NetBox 4.4.1+ validation bounds encoded as JSON
// floats (10.0) as well as integers (10). encoding/json cannot unmarshal
// a number with a decimal point into *int64.
func (m *CustomField) UnmarshalJSON(data []byte) error {
	type alias CustomField
	aux := &struct {
		*alias
		ValidationMaximum *json.Number `json:"validation_maximum"`
		ValidationMinimum *json.Number `json:"validation_minimum"`
	}{
		alias: (*alias)(m),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	vmax, err := int64PtrFromJSONNumber(aux.ValidationMaximum)
	if err != nil {
		return fmt.Errorf("validation_maximum: %w", err)
	}
	vmin, err := int64PtrFromJSONNumber(aux.ValidationMinimum)
	if err != nil {
		return fmt.Errorf("validation_minimum: %w", err)
	}
	m.ValidationMaximum = vmax
	m.ValidationMinimum = vmin
	return nil
}

func int64PtrFromJSONNumber(n *json.Number) (*int64, error) {
	if n == nil || string(*n) == "" {
		return nil, nil
	}
	i, err := n.Int64()
	if err != nil {
		f, ferr := strconv.ParseFloat(string(*n), 64)
		if ferr != nil {
			return nil, err
		}
		i = int64(f)
	}
	return &i, nil
}
