package d3money

// Currency defines a currency and its properties.
type Currency interface {
	Name() string         // Name returns the name of the currency.
	StandardName() string // StandardName returns an alphanumeric string that identifies the standard the currency is defined in.

	UniqueID() int32     // UniqueID returns the unique ID of the currency. This is specific to this library. All positive IDs are reserved for use in this library. If you add your own currencies use negative numbers to prevent collisions with the built in currencies in the future.
	UniqueCode() string  // Code returns a string representing the currency. This representation is unique across different currency standards. Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC"
	Code() string        // Code returns a string representing the currency. This may be an ISO 4217 code, depending on the currency type and is unique in a single currency standard. Examples: "USD", "AUD", "EUR", "BTC"
	Symbol() string      // Symbol returns a string containing the symbol of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This also doesn't follow any official standard. Examples: "US$", "AU$", "€", "₿"
	ShortSymbol() string // ShortSymbol returns a string containing the short symbol variant of the currency. This may be ambiguous, and should only be used for formatting into a human readable format. This needs additional context when used in text output, otherwise there is no way to differentiate between AUD and USD for example. This also doesn't follow any official standard. Examples: "$", "$", "€", "₿"

	DecimalPlaces() (int, bool) // DecimalPlaces returns the number of decimal places that represents the "Minor unit". If the resulting number is 0, this currency can't be divided any further. If the resulting bool is false, there is no smallest unit.

	// TODO: Add information if the currency is still official, still in use, ...
}
