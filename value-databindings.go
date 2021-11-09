// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
		if cur = Currencies.ByUniqueCode(d.Currency); cur == nil {
			return fmt.Errorf("can't find unique currency code %q", d.Currency)
		}
	}

	v.value, v.currency = d.Value, cur

	return nil
}

// Value implements the valuer interface of databases.
func (v Value) Value() (driver.Value, error) {
	if v.currency != nil {
		// Currency defined, output "Value UniqueCode" pair. Use a normal space for database storage.
		return fmt.Sprintf("%s %s", v.Decimal(), v.currency.UniqueCode()), nil
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

	val, newCur, err := parse(str, Currencies, nil)
	if err != nil {
		return fmt.Errorf("failed to parse string %q: %w", str, err)
	}

	v.value, v.currency = val, newCur

	return nil
}

// GormDBDataType returns the datatype that a database should use for the field.
func (v Value) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// Use field.Tag, field.TagSettings gets field's tags.
	// Checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options.

	// Return database type based on driver name.
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "VARCHAR(255)"
	case "postgres":
		return "VARCHAR"
	}
	return ""
}
