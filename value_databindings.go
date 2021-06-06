package d3money

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

// MarshalJSON returns the marshaled representation of the object.
func (v Value) MarshalJSON() ([]byte, error) {
	d := struct {
		Value    decimal.Decimal
		Currency string `json:",omitempty"`
	}{
		Value: v.value,
	}

	if v.currency != nil {
		d.Currency = v.currency.UniqueCode()
	}

	return json.Marshal(d)
}

// UnmarshalJSON fills the object with data matching the json representation.
// This will not be called if the JSON field of this value doesn't exist, therefore old data may persist after unmarshalling.
func (v *Value) UnmarshalJSON(data []byte) error {
	d := struct {
		Value    decimal.Decimal
		Currency string `json:",omitempty"`
	}{}

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	var cur Currency
	if d.Currency != "" {
		var err error
		cur, err = MatchCurrencyByUniqueCode(d.Currency, v.currencyStandards...)
		if err != nil {
			return fmt.Errorf("can't get currency: %w", err)
		}
	}

	v.value, v.currency = d.Value, cur

	return nil
}
