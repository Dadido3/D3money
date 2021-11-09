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
	Name() string         // Name returns the name of the currency. This is the english or native name.
	StandardName() string // StandardName returns an alphanumeric string that identifies the standard the currency is defined in.

	UniqueID() int32     // UniqueID returns the unique ID of the currency. This is specific to this library. All positive IDs are reserved for use in this library. If you add your own currencies use negative numbers to prevent collisions with the built in currencies in the future.
	UniqueCode() string  // Code returns a string representing the currency. This representation is unique across different currency standards. Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC"
	Code() string        // Code returns a string representing the currency. This may be an ISO 4217 code, depending on the currency type and is unique in a single currency standard. Examples: "USD", "AUD", "EUR", "BTC"
	Symbol() string      // Symbol returns a string containing the symbol of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This also doesn't follow any official standard. Examples: "US$", "AU$", "€", "₿"
	ShortSymbol() string // ShortSymbol returns a string containing the short symbol variant of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This needs additional context when used in text output, otherwise there is no way to differentiate between AUD and USD for example. This also doesn't follow any official standard. Examples: "$", "$", "€", "₿"

	DecimalPlaces() (decimalPlaces int, hasSmallestUnit bool) // DecimalPlaces returns the number of decimal places that represents the "Minor unit". If the resulting number is 0, this currency can't be divided any further. If the resulting bool is false and/or if the number of decimal places is -1, there is no smallest unit.

	// TODO: Add information if the currency is still official, still in use, ...
}

var regexFindNonAlphaNumeric = regexp.MustCompile("[^a-zA-Z0-9]")

// ValidateCurrency checks if the given currency contains illegal characters and such things.
func ValidateCurrency(c Currency) error {
	standardName, code, uniqueCode, symbol, shortSymbol := c.StandardName(), c.Code(), c.UniqueCode(), c.Symbol(), c.ShortSymbol()

	// The code should only contain alpha numeric characters.
	if regexFindNonAlphaNumeric.MatchString(code) {
		firstMatch := regexFindNonAlphaNumeric.FindString(code)
		return fmt.Errorf("result of Code() contains illegal character(s) %q", firstMatch)
	}

	// The standard name should only contain alpha numeric characters.
	if regexFindNonAlphaNumeric.MatchString(standardName) {
		firstMatch := regexFindNonAlphaNumeric.FindString(standardName)
		return fmt.Errorf("result of StandardName() contains illegal character(s) %q", firstMatch)
	}

	// Check if the unique code is in the following form: "StandardName-Code".
	if fmt.Sprintf("%s-%s", standardName, code) != uniqueCode {
		return fmt.Errorf("unique name %q is not of the form \"StandardName-Code\". Expected \"%s-%s\"", uniqueCode, standardName, code)
	}

	// Check for illegal result of DecimalPlaces().
	if decimalPlaces, hasSmallestUnit := c.DecimalPlaces(); hasSmallestUnit {
		// Has smallest unit.
		if decimalPlaces < 0 {
			return fmt.Errorf("currency has smallest unit, but %d decimal places", decimalPlaces)
		}
	} else {
		// Has no smallest unit.
		if decimalPlaces != -1 {
			return fmt.Errorf("currency has no smallest unit, but %d decimal places. Expects %d decimal places", decimalPlaces, -1)
		}
	}

	// Symbol() and ShortSymbol() should both return equal strings, or both need to be non empty strings.
	// Allowed: "US$" and "$"
	// Allowed: "£" and "£"
	// Allowed: "" and ""
	// Illegal: "" and "$"
	// Illegal: "US$" and ""
	if shortSymbol != symbol && (shortSymbol == "" || symbol == "") {
		return fmt.Errorf("symbol is %q and short symbol is %q. Both need to be the same, or both need to be non empty strings", symbol, shortSymbol)
	}

	return nil
}
