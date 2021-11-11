// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"database/sql/driver"
	"encoding/binary"
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
			return fmt.Errorf("can't find currency with unique code %q", d.Currency)
		}
	}

	v.value, v.currency = d.Value, cur

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (v Value) MarshalBinary() ([]byte, error) {
	data1 := []byte{0, 0, 0, 0}
	if v.currency != nil {
		binary.BigEndian.PutUint32(data1, uint32(v.currency.UniqueID()))
	}

	data2, err := v.value.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return append(data1, data2...), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (v *Value) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", data, len(data))
	}

	if uniqueID := int32(binary.BigEndian.Uint32(data[:4])); uniqueID == 0 {
		v.currency = nil
	} else {
		cur := Currencies.ByUniqueID(uniqueID)
		if cur == nil {
			return fmt.Errorf("can't find currency with unique ID %d", uniqueID)
		}
		v.currency = cur
	}

	return v.value.UnmarshalBinary(data[4:])
}

// MarshalText implements the encoding.TextMarshaler interface.
func (v Value) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (v *Value) UnmarshalText(text []byte) error {
	val, cur, err := parse(string(text), Currencies, nil)
	if err != nil {
		return fmt.Errorf("failed to parse text %q: %w", string(text), err)
	}

	v.value, v.currency = val, cur
	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (v Value) GobEncode() ([]byte, error) {
	return v.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface.
func (v *Value) GobDecode(data []byte) error {
	return v.UnmarshalBinary(data)
}

// Value implements the valuer interface of databases.
func (v Value) Value() (driver.Value, error) {
	return v.String(), nil
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
