package d3money

import "fmt"

// Currency defines a currency and its properties.
type Currency interface {
	Name() string // Name returns the name of the currency.

	UniqueCode() string  // Code returns a string representing the currency. This representation is unique across different currency standards. Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC"
	Code() string        // Code returns a string representing the currency. This may be an ISO 4217 code, depending on the currency type. Examples: "USD", "AUD", "EUR", "BTC"
	Symbol() string      // Symbol returns a string containing the symbol of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This also doesn't follow any official standard. Examples: "US$", "AU$", "€", "₿"
	ShortSymbol() string // ShortSymbol returns a string containing the short symbol variant of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This needs additional context when used in text output, otherwise there is no way to differentiate between AUD and USD for example. This also doesn't follow any official standard. Examples: "$", "$", "€", "₿"

	DecimalPlaces() (int, bool) // DecimalPlaces returns the number of decimal places that represents the "Minor unit". If the resulting bool is false, there is no smallest unit.

	Standard() string // Standard returns an alphanumeric string that identifies the standard the currency is defined in. Examples: "ISO4217"
}

// MatchCurrencyByUniqueCode finds a currency by its unique code (e.g. "ISO4217-EUR").
func MatchCurrencyByUniqueCode(uniqueCode string, currencyStandards ...map[string]Currency) (Currency, error) {
	for _, currencyStandard := range currencyStandards {
		// TODO: Use the map, instead of iterating over all currencies
		for _, currency := range currencyStandard {
			if currency != nil {
				if currency.UniqueCode() == uniqueCode {
					return currency, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("couldn't find or match currency %q", uniqueCode)
}
