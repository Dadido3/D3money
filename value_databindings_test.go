package d3money

import (
	"encoding/json"
	"testing"
)

func TestJSONMarshalling(t *testing.T) {
	val, err := NewFromStringWithCurrencies("-12345.67 ISO4217-EUR", nil, ISO4217Currencies)
	if err != nil {
		t.Errorf("NewFromStringWithCurrencies() failed: %v", err)
	}

	jsonBytes, err := json.Marshal(val)
	if err != nil {
		t.Errorf("json.Marshal() failed: %v", err)
	}

	correct := `{"Value":"-12345.67","Currency":"ISO4217-EUR"}`
	if string(jsonBytes) != correct {
		t.Errorf("json.Marshal() = %q: want %q", jsonBytes, correct)
	}

	val2 := NewWithCurrencies(ISO4217Currencies)

	if err := json.Unmarshal(jsonBytes, &val2); err != nil {
		t.Errorf("json.Unmarshal() failed: %v", err)
	}

	if !val.Equal(val2) {
		t.Errorf("val2 = %v, want %v", val2, val)
	}
}
