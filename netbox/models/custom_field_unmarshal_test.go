package models

import (
	"encoding/json"
	"testing"
)

func TestCustomFieldUnmarshal_validationBoundsAcceptFloatJSON(t *testing.T) {
	raw := []byte(`{
		"id": 1,
		"name": "cf_integer",
		"validation_minimum": 10.0,
		"validation_maximum": 1000.0
	}`)

	var cf CustomField
	if err := json.Unmarshal(raw, &cf); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if cf.ValidationMinimum == nil || *cf.ValidationMinimum != 10 {
		t.Fatalf("validation_minimum = %v, want 10", cf.ValidationMinimum)
	}
	if cf.ValidationMaximum == nil || *cf.ValidationMaximum != 1000 {
		t.Fatalf("validation_maximum = %v, want 1000", cf.ValidationMaximum)
	}
}

func TestCustomFieldUnmarshal_validationBoundsAcceptIntegerJSON(t *testing.T) {
	raw := []byte(`{
		"id": 1,
		"name": "cf_integer",
		"validation_minimum": 10,
		"validation_maximum": 1000
	}`)

	var cf CustomField
	if err := json.Unmarshal(raw, &cf); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if cf.ValidationMinimum == nil || *cf.ValidationMinimum != 10 {
		t.Fatalf("validation_minimum = %v, want 10", cf.ValidationMinimum)
	}
	if cf.ValidationMaximum == nil || *cf.ValidationMaximum != 1000 {
		t.Fatalf("validation_maximum = %v, want 1000", cf.ValidationMaximum)
	}
}

func TestCustomFieldUnmarshal_validationBoundsOmittedStayNil(t *testing.T) {
	raw := []byte(`{"id": 1, "name": "cf_text"}`)

	var cf CustomField
	if err := json.Unmarshal(raw, &cf); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if cf.ValidationMinimum != nil {
		t.Fatalf("validation_minimum = %v, want nil", *cf.ValidationMinimum)
	}
	if cf.ValidationMaximum != nil {
		t.Fatalf("validation_maximum = %v, want nil", *cf.ValidationMaximum)
	}
}
