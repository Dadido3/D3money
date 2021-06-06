package d3money

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Value represents a monetary value in a specific currency.
type Value struct {
	value decimal.Decimal

	// Currency can be nil.
	currency Currency

	// List(s) of possible currencies to chose from.
	// This is used when unmarshalling from some data structure into a d3money.Value object.
	// Note: I don't want to use a global list of currencies for unmarshalling, as it may cause problems and is less customizable.
	currencyStandards []map[string]Currency
}

// NewWithCurrencies returns a value object with the given currencyStandards (Maps of available currencies) set.
// This should be used before unmarshalling, otherwise currencies can't be matched by their codes.
func NewWithCurrencies(currencyStandards ...map[string]Currency) Value {
	return Value{
		currencyStandards: currencyStandards,
	}
}

// NewFromString returns a value object from the given string.
// No currency matching will be done, and no currency will be given.
// If there is a currency code in the string, this function will return an error.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// Examples:
//	NewFromString("-10000.123")             // Returns a currency-less value.
//	NewFromString("-10000.123 ISO4217-EUR") // Returns an error, as the currency in the string can't be matched/found.
func NewFromString(str string) (Value, error) {
	val, _, err := parse(str, nil)
	if err != nil {
		return Value{}, err
	}

	return Value{value: val}, nil
}

// NewFromStringWithCurrency returns a value object from the given string.
// The field cur can be used to override the currency.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// Examples:
//	NewFromStringWithCurrency("-10000.123", nil)                                  // Returns a currency-less value.
//	NewFromStringWithCurrency("-10000.123 ISO4217-EUR", nil)                      // Returns an error, as the currency in the string can't be matched/found.
//	NewFromStringWithCurrency("-10000.123", ISO4217Currencies["EUR"])             // Returns a value with EUR currency.
//	NewFromStringWithCurrency("-10000.123 ISO4217-EUR", ISO4217Currencies["EUR"]) // Returns a value with EUR currency.
//	NewFromStringWithCurrency("-10000.123 ISO4217-USD", ISO4217Currencies["EUR"]) // Returns an error, as the currency in the string can't be matched/found.
func NewFromStringWithCurrency(str string, cur Currency) (Value, error) {
	// Add currency override to match list, so that it can be matched and checked.
	var currencies map[string]Currency
	if cur != nil {
		currencies = map[string]Currency{cur.Code(): cur}
	}

	// Parse string.
	val, newCur, err := parse(str, currencies)
	if err != nil {
		return Value{}, err
	}

	// Override currency, if the string doesn't contain any currency code.
	if newCur == nil {
		newCur = cur
	}

	return Value{value: val, currency: newCur}, nil
}

// NewFromStringWithCurrencies returns a value object from the given string.
// The field cur can be used to override the currency.
// The field currencies is used to define currencies for matching.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// Examples:
//	NewFromStringWithCurrencies("-10000.123", nil, nil)                                                // Returns a currency-less value.
//	NewFromStringWithCurrencies("-10000.123", nil, ISO4217Currencies)                                  // Returns a currency-less value.
//	NewFromStringWithCurrencies("-10000.123 ISO4217-EUR", nil, ISO4217Currencies)                      // Returns a value with EUR currency.
//	NewFromStringWithCurrencies("-10000.123", ISO4217Currencies["EUR"], ISO4217Currencies)             // Returns a value with EUR currency.
//	NewFromStringWithCurrencies("-10000.123 ISO4217-EUR", ISO4217Currencies["EUR"], ISO4217Currencies) // Returns a value with EUR currency.
//	NewFromStringWithCurrencies("-10000.123", ISO4217Currencies["EUR"], nil)                           // Returns a value with EUR currency.
//	NewFromStringWithCurrencies("-10000.123 ISO4217-EUR", ISO4217Currencies["EUR"], nil)               // Returns a value with EUR currency.
//	NewFromStringWithCurrencies("-10000.123 ISO4217-EUR", nil, nil)                                    // Returns an error, as the currency in the string can't be matched/found.
//	NewFromStringWithCurrencies("-10000.123 ISO4217-USD", ISO4217Currencies["EUR"], ISO4217Currencies) // Returns an error, as the currencies don't match.
func NewFromStringWithCurrencies(str string, cur Currency, currencyStandards ...map[string]Currency) (Value, error) {
	// Add currency override to match list, so that it can be matched and checked.
	if cur != nil {
		currencies := map[string]Currency{cur.Code(): cur}
		currencyStandards = append([]map[string]Currency{currencies}, currencyStandards...)
	}

	// Parse string.
	val, newCur, err := parse(str, currencyStandards...)
	if err != nil {
		return Value{}, err
	}

	if newCur == nil {
		newCur = cur
	}

	// Check currency.
	if cur != nil && newCur != nil && cur != newCur {
		return Value{}, fmt.Errorf("parsed %q and given currency %q don't match", newCur, cur)
	}

	return Value{value: val, currency: newCur}, nil
}

// RequireFromString returns a value object from the given string.
// No currency matching will be done, and no currency will be given.
// If there is a currency code in the string, this function will return an error.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// In case of an error, this will panic.
//
// For examples, see NewFromString().
func RequireFromString(str string) Value {
	v, err := NewFromString(str)
	if err != nil {
		panic(err)
	}

	return v
}

// RequireFromStringWithCurrency returns a value object from the given string.
// The field cur can be used to override the currency.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// In case of an error, this will panic.
//
// For examples, see NewFromStringWithCurrency().
func RequireFromStringWithCurrency(str string, cur Currency) Value {
	v, err := NewFromStringWithCurrency(str, cur)
	if err != nil {
		panic(err)
	}

	return v
}

// RequireFromStringWithCurrencies returns a value object from the given string.
// The field cur can be used to override the currency.
// The field currencies is used to define currencies for matching.
// This will not use any locale specific formatting, and is not suited for input from humans without any preprocessing.
//
// In case of an error, this will panic.
//
// For examples, see NewFromStringWithCurrencies().
func RequireFromStringWithCurrencies(str string, cur Currency, currencyStandards ...map[string]Currency) Value {
	v, err := NewFromStringWithCurrencies(str, cur, currencyStandards...)
	if err != nil {
		panic(err)
	}

	return v
}

// parse takes a string and parses its value and currency.
// The list of currencyStandards is used for matching.
func parse(str string, currencyStandards ...map[string]Currency) (decimal.Decimal, Currency, error) {
	var valStr, curStr string
	var matchedCurrency Currency

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

		// Match currency.
		var err error
		matchedCurrency, err = MatchCurrencyByUniqueCode(curStr, currencyStandards...)
		if err != nil {
			return decimal.Zero, nil, err
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

// Decimal returns the value as a shopspring/decimal number.
func (v Value) Decimal() decimal.Decimal {
	return v.value
}

// Currency returns the used currency.
func (v Value) Currency() Currency {
	return v.currency
}

// Equal returns if a monetary value is equal to another.
// No currency conversion is done, so if the currency differs this function will return false in any case.
func (v Value) Equal(comp Value) bool {
	if v.currency != comp.currency {
		return false
	}
	return v.value.Equal(comp.value)
}

// String returns the monetary value as string, while using a period as decimal separator.
// The unique currency code is placed as suffix with a non-breaking space in between.
func (v Value) String() string {
	if v.currency != nil {
		return fmt.Sprintf("%s\u00A0%s", v.value.String(), v.currency.UniqueCode())
	}

	// Fallback if there is no currency.
	return v.value.String()
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
