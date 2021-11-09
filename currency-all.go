// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

// Currencies contains all the currencies that come with this library.
// This will be used for unmarshalling or for scanning from a database.
// If you want to add custom currencies, add them to this collection by using 'Currencies.Add(...)'.
var Currencies = MustNewCurrencyCollection("D3money-currencies", iso4217Currencies)
