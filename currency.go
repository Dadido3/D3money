package d3money

// Currency defines a currency and its properties.
type Currency interface {
	Name() string // Name returns the name of the currency.

	Code() string        // Code returns a alphabetic code representing the currency. This may be a ISO 4217 code, depending on the currency type. Examples: "USD", "AUD", "EUR", "BTC"
	Symbol() string      // Symbol returns a string containing the symbol of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This also doesn't follow any official standard. Examples: "US$", "AU$", "€", "₿"
	ShortSymbol() string // ShortSymbol returns a string containing the short symbol variant of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This needs additional context when used in text output, otherwise there is no way to differentiate between AUD and USD for example. This also doesn't follow any official standard. Examples: "$", "$", "€", "₿"

	DecimalPlaces() (int, bool) // DecimalPlaces returns the number of decimal places that represents the "Minor unit". If the resulting bool is false, there is no smallest unit.

	Standard() string // Standard returns a string that identifies the standard that the currency is defined in. Examples: "ISO 4217"
}
