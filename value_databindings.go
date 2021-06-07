package d3money

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

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

// Value implements the valuer interface of databases.
func (v Value) Value() (driver.Value, error) {
	if v.currency != nil {
		// Currency defined, output "UniqueCode Value" pair. Use a normal space for database storage.
		return fmt.Sprintf("%s %s", v.currency.UniqueCode(), v.Decimal()), nil
	}

	// No currency defined, only output the value.
	return v.Decimal().String(), nil
}

// Scan fills the object with data matching the given value from the database.
func (v *Value) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("incompatible type %T, expected %T", value, str)
	}

	splitted := strings.Split(str, " ") // "UniqueCode Value". Use a normal space for database storage.
	switch len(splitted) {
	case 1:
		// String (probably) consists only of the value.
		var err error
		v.value, err = decimal.NewFromString(splitted[0])
		if err != nil {
			return fmt.Errorf("malformed monetary value: %w", err)
		}
		v.currency = nil

	case 2:
		// String (probably) consists of a value + unique code pair.
		var err error
		v.value, err = decimal.NewFromString(splitted[1])
		if err != nil {
			return fmt.Errorf("malformed monetary value: %w", err)
		}
		// Match currency.
		v.currency, err = MatchCurrencyByUniqueCode(splitted[0], v.currencyStandards...)
		if err != nil {
			return fmt.Errorf("can't get currency: %w", err)
		}

	default:
		return fmt.Errorf("malformed string %q: invalid amount of spaces", str)
	}

	return nil
}
