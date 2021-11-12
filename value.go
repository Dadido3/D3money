// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Value represents a monetary value in a specific currency.
type Value struct {
	value    decimal.Decimal
	currency Currency // Can be nil.
}

// FromString returns a value object from the given string.
// The string can contain a currency by its unique code.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// Examples:
//	FromString("-10000.123")             // Returns a currency-less value.
//	FromString("-10000.123 ISO4217-EUR") // Returns a value with the EUR currency defined by ISO 4217.
//	FromString("-10000.123 EUR")         // Returns an error, as the currency in the string can't be matched/found.
//	FromString("-10000.123 FOO-BAR")     // Result depends on whether the the custom currency "FOO-BAR" is registered. See `Currencies.Add(...)`.
func FromString(str string) (Value, error) {
	val, cur, err := parse(str, Currencies, nil)
	if err != nil {
		return Value{}, err
	}

	return Value{value: val, currency: cur}, nil
}

// MustFromString returns a value object from the given string.
// No currency matching will be done, and no currency will be given.
// If there is a currency code in the string, this function will return an error.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// In case of an error, this will panic.
//
// For examples, see FromString().
func MustFromString(str string) Value {
	v, err := FromString(str)
	if err != nil {
		panic(err)
	}

	return v
}

// FromStringAndCurrency returns a value object from the given string.
// The field cur can be used to define the currency.
// The string can contain a currency by its unique code, but it's checked whether it matches with the field cur.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// Examples:
//	FromStringAndCurrency("-10000.123", nil)                                         // Returns a currency-less value.
//	FromStringAndCurrency("-10000.123 ISO4217-EUR", nil)                             // Returns an error, as the currency differs from the one defined in field cur.
//	FromStringAndCurrency("-10000.123", ISO4217Currencies.ByCode("EUR"))             // Returns a value with EUR currency.
//	FromStringAndCurrency("-10000.123 ISO4217-EUR", ISO4217Currencies.ByCode("EUR")) // Returns a value with EUR currency.
//	FromStringAndCurrency("-10000.123 ISO4217-USD", ISO4217Currencies.ByCode("EUR")) // Returns an error, as the currency differs from the one defined in field cur.
//	FromStringAndCurrency("-10000.123 FOO-BAR", nil)                                 // Returns an error, as the currency differs from the one defined in field cur.
//	FromStringAndCurrency("-10000.123", FooBarCurrency)                              // Returns a value with custom currency.
//	FromStringAndCurrency("-10000.123 FOO-BAR", FooBarCurrency)                      // Returns a value with custom currency. This assumes that the unique code of that currency is "FOO-BAR".
func FromStringAndCurrency(str string, cur Currency) (Value, error) {
	val, newCur, err := parse(str, Currencies, cur)
	if err != nil {
		return Value{}, err
	}

	// The string doesn't contain a unique code, overwrite it with the user defined currency.
	if newCur == nil {
		newCur = cur
	}

	// Check if the parsed and the defined currencies match.
	// They will always match if the string doesn't contain a unique code.
	if newCur != cur {
		return Value{}, fmt.Errorf("the matched currency %q of the string doesn't match with the given one %q", newCur, cur)
	}

	return Value{value: val, currency: newCur}, nil
}

// MustFromStringAndCurrency returns a value object from the given string.
// The field cur can be used to define the currency.
// The string can contain a currency by its unique code, but it's checked whether it matches with the field cur.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// In case of an error, this will panic.
//
// For examples, see FromStringAndCurrency().
func MustFromStringAndCurrency(str string, cur Currency) Value {
	v, err := FromStringAndCurrency(str, cur)
	if err != nil {
		panic(fmt.Sprintf("failed to create currency: %v", err))
	}

	return v
}

// FromDecimal returns a value object from the given shopspring/decimal and currency.
func FromDecimal(val decimal.Decimal, cur Currency) Value {
	return Value{value: val, currency: cur}
}

// FromFloat64 returns a value object from the given float64 and currency.
func FromFloat64(f float64, cur Currency) Value {
	return Value{value: decimal.NewFromFloat(f), currency: cur}
}

// FromFloat32 returns a value object from the given float32 and currency.
func FromFloat32(f float32, cur Currency) Value {
	return Value{value: decimal.NewFromFloat32(f), currency: cur}
}

// FromInt64 returns a value object from the given int64 and currency.
func FromInt64(i int64, cur Currency) Value {
	return Value{value: decimal.NewFromInt(i), currency: cur}
}

// FromInt32 returns a value object from the given int32 and currency.
func FromInt32(i int32, cur Currency) Value {
	return Value{value: decimal.NewFromInt32(i), currency: cur}
}

// String returns the monetary value as a "Value UniqueCode" pair.
// This is locale independent.
func (v Value) String() string {
	if v.currency != nil {
		// Output "Value UniqueCode" pair.
		return v.value.String() + " " + v.currency.UniqueCode()
	}

	// If there is no currency output "Value" only.
	return v.value.String()
}

// parse takes a string and parses its value and currency delimited by a space or non breaking space ("Value UniqueCode").
// The additionalCurrency or the list of currencies is used for matching/checking the unique code.
// A match from additionalCurrency has a higher priority than a match from the list of currencies.
//
// This can parse the output of value.String() without loss of information.
//
// This will always return a currency if the input string contains a value and currency pair that is correctly delimited, in any other case it will return nil as currency!
func parse(str string, cc CurrencyCollection, additionalCurrency Currency) (decimal.Decimal, Currency, error) {
	var valStr, curStr string
	var matchedCurrency Currency

	// TODO: Check if cc contains a currency that violates the uniqueness constraint with the additionalCurrency.

	str = strings.ReplaceAll(str, "\u00A0", " ")

	// Parse expression.
	splitted := strings.Split(str, " ")
	switch len(splitted) {
	case 1:
		// String (probably) contains a number value.
		valStr = splitted[0]

	case 2:
		// String (probably) contains a number value + unique currency code.
		valStr, curStr = splitted[0], splitted[1]

		// Check if additionalCurrency matches with the unique code.
		if additionalCurrency != nil {
			if curStr == additionalCurrency.UniqueCode() {
				matchedCurrency = additionalCurrency
			}
		}

		// If there is no match, look in collection.
		if matchedCurrency == nil && cc != nil {
			matchedCurrency = cc.ByUniqueCode(curStr)
		}

		// If there is still no match, return error.
		if matchedCurrency == nil {
			return decimal.Zero, nil, &ErrorCantFindUniqueCode{curStr}
		}

	default:
		return decimal.Zero, nil, fmt.Errorf("input string %q contains too many spaces", str)
	}

	// Parse value string.
	value, err := decimal.NewFromString(valStr)
	if err != nil {
		return decimal.Zero, nil, err
	}

	return value, matchedCurrency, nil

}

// Currency returns the used currency.
func (v Value) Currency() Currency {
	return v.currency
}

// Format implements the fmt.Formatter interface.
/*func (v Value) Format(state fmt.State, verb rune) {
	switch verb {
	case 'f':
		if prec, ok := state.Precision(); ok {
			fmt.Fprintf(state, v.value.StringFixed(int32(prec)))
		} else {
			fmt.Fprintf(state, v.value.String())
		}
	}
}*/
