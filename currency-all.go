package d3money

// Currencies contains all the currencies that come with this library.
// This will be used for unmarshalling or for scanning from a database.
// If you want to add custom currencies, add them to this collection by using 'Currencies.AddCurrencies()'.
var Currencies = MustNewCurrencyCollection("d3money-currencies", iso4217Currencies)
