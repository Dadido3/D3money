// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"
	"regexp"
)

// Currency defines a currency and its properties.
type Currency interface {
	Name() string     // Name returns the name of the currency. This is the english or native name.
	Standard() string // Standard returns an alphanumeric string that identifies the standard the currency is defined in.

	UniqueID() int32      // UniqueID returns the unique ID of the currency. This is specific to this library. All positive IDs are reserved for use in this library. If you add your own currencies use negative numbers to prevent collisions with the built in currencies in the future.
	UniqueCode() string   // Code returns a string representing the currency. This representation is unique across different currency standards. Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC".
	NumericCode() int     // NumericCode returns the numeric code of the currency. This may be an ISO 4217 numeric code, depending on the currency type and is unique in a single currency standard. Examples: 840, 36, 978.
	Code() string         // Code returns a string representing the currency. This may be an ISO 4217 code, depending on the currency type and is unique in a single currency standard. Examples: "USD", "AUD", "EUR", "BTC".
	Symbol() string       // Symbol returns a string containing the symbol of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This also doesn't follow any official standard. Examples: "US$", "AU$", "€", "₿"
	NarrowSymbol() string // NarrowSymbol returns a string containing the narrow symbol variant of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This needs additional context when used in text output, otherwise there is no way to differentiate between AUD and USD for example. This also doesn't follow any official standard. Examples: "$", "$", "€", "₿".

	DecimalPlaces() (decimalPlaces int, hasSmallestUnit bool) // DecimalPlaces returns the number of decimal places that represents the "Minor unit". If the resulting number is 0, this currency can't be divided any further. If the resulting bool is false and/or if the number of decimal places is -1, there is no smallest unit.
}

var regexFindNonAlphaNumeric = regexp.MustCompile("[^A-Z0-9]")

// ValidateCurrency checks if the given currency contains illegal characters and such things.
func ValidateCurrency(c Currency) error {
	standard, code, uniqueCode, symbol, narrowSymbol := c.Standard(), c.Code(), c.UniqueCode(), c.Symbol(), c.NarrowSymbol()

	// No currency should have an unique ID of 0.
	if c.UniqueID() == 0 {
		return &ErrorInvalidCurrency{"unique ID is 0. This value is reserved for \"no currency\""}
	}

	// The code should only contain alphanumeric characters.
	if regexFindNonAlphaNumeric.MatchString(code) {
		firstMatch := regexFindNonAlphaNumeric.FindString(code)
		return &ErrorInvalidCurrency{fmt.Sprintf("code contains illegal character(s) %q", firstMatch)}
	}

	// The standard name should only contain alphanumeric characters.
	if regexFindNonAlphaNumeric.MatchString(standard) {
		firstMatch := regexFindNonAlphaNumeric.FindString(standard)
		return &ErrorInvalidCurrency{fmt.Sprintf("standard name contains illegal character(s) %q", firstMatch)}
	}

	// Check if the unique code is in the following form: "Standard-Code".
	if fmt.Sprintf("%s-%s", standard, code) != uniqueCode {
		return &ErrorInvalidCurrency{fmt.Sprintf("unique code %q is not of the form \"Standard-Code\". Expected \"%s-%s\"", uniqueCode, standard, code)}
	}

	// Check for illegal result of DecimalPlaces().
	if decimalPlaces, hasSmallestUnit := c.DecimalPlaces(); hasSmallestUnit {
		// Has smallest unit.
		if decimalPlaces < 0 {
			return &ErrorInvalidCurrency{fmt.Sprintf("currency has smallest unit, but %d decimal places", decimalPlaces)}
		}
	} else {
		// Has no smallest unit.
		if decimalPlaces != -1 {
			return &ErrorInvalidCurrency{fmt.Sprintf("currency has no smallest unit, but %d decimal places. Expects %d decimal places", decimalPlaces, -1)}
		}
	}

	// Symbol() and NarrowSymbol() should both return equal strings, or both need to be non empty strings.
	// Allowed: "US$" and "$"
	// Allowed: "£" and "£"
	// Allowed: "" and ""
	// Illegal: "" and "$"
	// Illegal: "US$" and ""
	if narrowSymbol != symbol && (narrowSymbol == "" || symbol == "") {
		return &ErrorInvalidCurrency{fmt.Sprintf("symbol is %q and narrow symbol is %q. Both need to be the same, or both need to be non empty strings", symbol, narrowSymbol)}
	}

	return nil
}
