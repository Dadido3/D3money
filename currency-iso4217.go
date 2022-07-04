// Copyright (c) 2021-2022 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"github.com/shopspring/decimal"
	"golang.org/x/text/language"
)

// iso4217Name is the name of the currency standard.
const iso4217Name = "ISO4217"

// iso4217CodePrefix is the prefix that is used to generate the unique currency code.
const iso4217CodePrefix = "ISO4217-"

// iso4217UniqueIDOffset is the offset that is used in combination with the numeric code for the unique currency ID.
const iso4217UniqueIDOffset = 42170000

// ISO4217Currency defines a currency according to the ISO 4217 standard.
type ISO4217Currency struct {
	name string // English name of the currency.

	alphabeticCode       string // Official 3 alphabetic letter code.
	numericCode          int    // Official numeric code.
	narrowSymbol, symbol string // Additional representation forms. Official, but not standardized.

	smallestUnit decimal.Decimal // The value of the smallest unit for the given currency.
}

// Make sure this type implements the Currency interface.
var _ Currency = (*ISO4217Currency)(nil)

// Name returns the name of the currency.
// This is the english or native name.
func (c *ISO4217Currency) Name() string {
	return c.name
}

// Standard returns an alphanumeric string that identifies the standard the currency is defined in.
func (c *ISO4217Currency) Standard() string {
	return iso4217Name
}

// CodePrefix is the prefix that is used to generate the unique currency code.
func (c *ISO4217Currency) CodePrefix() string {
	return iso4217CodePrefix
}

// UniqueID returns the unique ID of the currency.
// This is specific to this library.
//
// All positive IDs are reserved for use in this library.
// If you add your own currencies use negative numbers to prevent collisions with the built in currencies in the future.
func (c *ISO4217Currency) UniqueID() int32 {
	return int32(c.numericCode) + iso4217UniqueIDOffset
}

// Code returns a string representing the currency.
// This representation is unique across different currency standards.
//
// Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC".
func (c *ISO4217Currency) UniqueCode() string {
	return c.CodePrefix() + c.alphabeticCode
}

// NumericCode returns the numeric code of the currency.
// This may be an ISO 4217 numeric code, depending on the currency type and is unique in a single currency standard.
//
// Examples: 840, 36, 978.
func (c *ISO4217Currency) NumericCode() int {
	return c.numericCode
}

// Code returns a string representing the currency.
// This is the official code defined by the standard, but it may not be unique across different standards.
// This may be an ISO 4217 code, depending on the currency type.
//
// Examples: "USD", "AUD", "EUR", "BTC".
func (c *ISO4217Currency) Code() string {
	return c.alphabeticCode
}

// Symbol returns a string containing the symbol of the currency.
// This may be ambiguous, and should only be used for formatting into a human readable format.
//
// Examples: "US$", "AU$", "€", "₿".
func (c *ISO4217Currency) Symbol(lang language.Tag) string {
	return c.symbol
}

// NarrowSymbol returns a string containing the narrow symbol variant of the currency.
// This may be ambiguous, and should only be used for formatting into a human readable format.
// This needs additional context when used in text output, as it doesn't differentiate between all the dollar currencies.
//
// Examples: "$", "$", "€", "₿".
func (c *ISO4217Currency) NarrowSymbol(lang language.Tag) string {
	if c.narrowSymbol != "" {
		return c.narrowSymbol
	}
	return c.symbol
}

// SmallestUnit returns the value of the fractional unit.
// This can be any value, even one that is larger than 1.
// A value of 0 means that there is no smallest unit.
func (c *ISO4217Currency) SmallestUnit() Value {
	return Value{c.smallestUnit, c}
}

func (c *ISO4217Currency) String() string {
	return c.CodePrefix() + c.alphabeticCode
}

// iso4217Currencies contains the official and active ISO 4217 currencies as of August 29, 2018.
//
// Source: https://www.currency-iso.org/en/home/tables/table-a1.html
//
// This data has been modified in the following ways:
//  - Removed "ENTITY" column
//  - Removed duplicate entries (due to removal of the "ENTITY" column)
//  - Add "symbol" and "narrowSymbol" columns that contain symbols which are NOT part of ISO 4217. Based on https://web.archive.org/web/20111129141202/http://fx.sauder.ubc.ca/currency_table.html and https://wikipedia.org
var iso4217Currencies = []Currency{
	&ISO4217Currency{alphabeticCode: "AFN", numericCode: 971, symbol: "؋", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Afghani"},
	&ISO4217Currency{alphabeticCode: "EUR", numericCode: 978, symbol: "€", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Euro"},
	&ISO4217Currency{alphabeticCode: "ALL", numericCode: 8, symbol: "L", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Lek"},
	&ISO4217Currency{alphabeticCode: "DZD", numericCode: 12, symbol: "DA", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Algerian Dinar"},
	&ISO4217Currency{alphabeticCode: "USD", numericCode: 840, symbol: "US$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "US Dollar"},
	&ISO4217Currency{alphabeticCode: "AOA", numericCode: 973, symbol: "Kz", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Kwanza"},
	&ISO4217Currency{alphabeticCode: "XCD", numericCode: 951, symbol: "EC$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "East Caribbean Dollar"},
	&ISO4217Currency{alphabeticCode: "ARS", numericCode: 32, symbol: "$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Argentine Peso"},
	&ISO4217Currency{alphabeticCode: "AMD", numericCode: 51, symbol: "֏", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Armenian Dram"},
	&ISO4217Currency{alphabeticCode: "AWG", numericCode: 533, symbol: "ƒ", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Aruban Florin"},
	&ISO4217Currency{alphabeticCode: "AUD", numericCode: 36, symbol: "AU$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Australian Dollar"},
	&ISO4217Currency{alphabeticCode: "AZN", numericCode: 944, symbol: "₼", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Azerbaijan Manat"},
	&ISO4217Currency{alphabeticCode: "BSD", numericCode: 44, symbol: "B$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Bahamian Dollar"},
	&ISO4217Currency{alphabeticCode: "BHD", numericCode: 48, symbol: "BD", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Bahraini Dinar"},
	&ISO4217Currency{alphabeticCode: "BDT", numericCode: 50, symbol: "৳", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Taka"},
	&ISO4217Currency{alphabeticCode: "BBD", numericCode: 52, symbol: "Bds$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Barbados Dollar"},
	&ISO4217Currency{alphabeticCode: "BYN", numericCode: 933, symbol: "Br", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Belarusian Ruble"},
	&ISO4217Currency{alphabeticCode: "BZD", numericCode: 84, symbol: "BZ$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Belize Dollar"},
	&ISO4217Currency{alphabeticCode: "XOF", numericCode: 952, symbol: "CFA", narrowSymbol: "Fr", smallestUnit: decimal.New(1, 0), name: "CFA Franc BCEAO"},
	&ISO4217Currency{alphabeticCode: "BMD", numericCode: 60, symbol: "BD$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Bermudian Dollar"},
	&ISO4217Currency{alphabeticCode: "INR", numericCode: 356, symbol: "₹", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Indian Rupee"},
	&ISO4217Currency{alphabeticCode: "BTN", numericCode: 64, symbol: "Nu.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Ngultrum"},
	&ISO4217Currency{alphabeticCode: "BOB", numericCode: 68, symbol: "Bs.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Boliviano"},
	&ISO4217Currency{alphabeticCode: "BOV", numericCode: 984, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Mvdol"},
	&ISO4217Currency{alphabeticCode: "BAM", numericCode: 977, symbol: "KM", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Convertible Mark"},
	&ISO4217Currency{alphabeticCode: "BWP", numericCode: 72, symbol: "P", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Pula"},
	&ISO4217Currency{alphabeticCode: "NOK", numericCode: 578, symbol: "kr", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Norwegian Krone"},
	&ISO4217Currency{alphabeticCode: "BRL", numericCode: 986, symbol: "R$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Brazilian Real"},
	&ISO4217Currency{alphabeticCode: "BND", numericCode: 96, symbol: "B$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Brunei Dollar"},
	&ISO4217Currency{alphabeticCode: "BGN", numericCode: 975, symbol: "лв.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Bulgarian Lev"},
	&ISO4217Currency{alphabeticCode: "BIF", numericCode: 108, symbol: "FBu", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Burundi Franc"},
	&ISO4217Currency{alphabeticCode: "CVE", numericCode: 132, symbol: "Esc", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Cabo Verde Escudo"},
	&ISO4217Currency{alphabeticCode: "KHR", numericCode: 116, symbol: "៛", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Riel"},
	&ISO4217Currency{alphabeticCode: "XAF", numericCode: 950, symbol: "CFA", narrowSymbol: "Fr", smallestUnit: decimal.New(1, 0), name: "CFA Franc BEAC"},
	&ISO4217Currency{alphabeticCode: "CAD", numericCode: 124, symbol: "CA$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Canadian Dollar"},
	&ISO4217Currency{alphabeticCode: "KYD", numericCode: 136, symbol: "KY$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Cayman Islands Dollar"},
	&ISO4217Currency{alphabeticCode: "CLP", numericCode: 152, symbol: "CLP$", narrowSymbol: "$", smallestUnit: decimal.New(1, 0), name: "Chilean Peso"},
	&ISO4217Currency{alphabeticCode: "CLF", numericCode: 990, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -4), name: "Unidad de Fomento"},
	&ISO4217Currency{alphabeticCode: "CNY", numericCode: 156, symbol: "¥", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Yuan Renminbi"},
	&ISO4217Currency{alphabeticCode: "COP", numericCode: 170, symbol: "Col$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Colombian Peso"},
	&ISO4217Currency{alphabeticCode: "COU", numericCode: 970, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Unidad de Valor Real"},
	&ISO4217Currency{alphabeticCode: "KMF", numericCode: 174, symbol: "CF", narrowSymbol: "Fr", smallestUnit: decimal.New(1, 0), name: "Comorian Franc "},
	&ISO4217Currency{alphabeticCode: "CDF", numericCode: 976, symbol: "F", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Congolese Franc"},
	&ISO4217Currency{alphabeticCode: "NZD", numericCode: 554, symbol: "NZ$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "New Zealand Dollar"},
	&ISO4217Currency{alphabeticCode: "CRC", numericCode: 188, symbol: "₡", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Costa Rican Colon"},
	&ISO4217Currency{alphabeticCode: "HRK", numericCode: 191, symbol: "kn", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Kuna"},
	&ISO4217Currency{alphabeticCode: "CUP", numericCode: 192, symbol: "₱", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Cuban Peso"},
	&ISO4217Currency{alphabeticCode: "CUC", numericCode: 931, symbol: "$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Peso Convertible"},
	&ISO4217Currency{alphabeticCode: "ANG", numericCode: 532, symbol: "NAƒ", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Netherlands Antillean Guilder"},
	&ISO4217Currency{alphabeticCode: "CZK", numericCode: 203, symbol: "Kč", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Czech Koruna"},
	&ISO4217Currency{alphabeticCode: "DKK", numericCode: 208, symbol: "Kr", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Danish Krone"},
	&ISO4217Currency{alphabeticCode: "DJF", numericCode: 262, symbol: "Fdj", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Djibouti Franc"},
	&ISO4217Currency{alphabeticCode: "DOP", numericCode: 214, symbol: "RD$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Dominican Peso"},
	&ISO4217Currency{alphabeticCode: "EGP", numericCode: 818, symbol: "E£", narrowSymbol: "£", smallestUnit: decimal.New(1, -2), name: "Egyptian Pound"},
	&ISO4217Currency{alphabeticCode: "SVC", numericCode: 222, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "El Salvador Colon"},
	&ISO4217Currency{alphabeticCode: "ERN", numericCode: 232, symbol: "Nkf", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Nakfa"},
	&ISO4217Currency{alphabeticCode: "SZL", numericCode: 748, symbol: "L", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Lilangeni"},
	&ISO4217Currency{alphabeticCode: "ETB", numericCode: 230, symbol: "Br", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Ethiopian Birr"},
	&ISO4217Currency{alphabeticCode: "FKP", numericCode: 238, symbol: "£", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Falkland Islands Pound"},
	&ISO4217Currency{alphabeticCode: "FJD", numericCode: 242, symbol: "FJ$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Fiji Dollar"},
	&ISO4217Currency{alphabeticCode: "XPF", numericCode: 953, symbol: "₣", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "CFP Franc"},
	&ISO4217Currency{alphabeticCode: "GMD", numericCode: 270, symbol: "D", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Dalasi"},
	&ISO4217Currency{alphabeticCode: "GEL", numericCode: 981, symbol: "₾", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Lari"},
	&ISO4217Currency{alphabeticCode: "GHS", numericCode: 936, symbol: "₵", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Ghana Cedi"},
	&ISO4217Currency{alphabeticCode: "GIP", numericCode: 292, symbol: "£", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Gibraltar Pound"},
	&ISO4217Currency{alphabeticCode: "GTQ", numericCode: 320, symbol: "Q", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Quetzal"},
	&ISO4217Currency{alphabeticCode: "GBP", numericCode: 826, symbol: "£", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Pound Sterling"},
	&ISO4217Currency{alphabeticCode: "GNF", numericCode: 324, symbol: "FG", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Guinean Franc"},
	&ISO4217Currency{alphabeticCode: "GYD", numericCode: 328, symbol: "GY$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Guyana Dollar"},
	&ISO4217Currency{alphabeticCode: "HTG", numericCode: 332, symbol: "G", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Gourde"},
	&ISO4217Currency{alphabeticCode: "HNL", numericCode: 340, symbol: "L", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Lempira"},
	&ISO4217Currency{alphabeticCode: "HKD", numericCode: 344, symbol: "HK$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Hong Kong Dollar"},
	&ISO4217Currency{alphabeticCode: "HUF", numericCode: 348, symbol: "Ft", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Forint"},
	&ISO4217Currency{alphabeticCode: "ISK", numericCode: 352, symbol: "kr", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Iceland Krona"},
	&ISO4217Currency{alphabeticCode: "IDR", numericCode: 360, symbol: "Rp", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Rupiah"},
	&ISO4217Currency{alphabeticCode: "XDR", numericCode: 960, symbol: "SDR", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "SDR (Special Drawing Right)"},
	&ISO4217Currency{alphabeticCode: "IRR", numericCode: 364, symbol: "﷼", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Iranian Rial"},
	&ISO4217Currency{alphabeticCode: "IQD", numericCode: 368, symbol: "د.ع", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Iraqi Dinar"},
	&ISO4217Currency{alphabeticCode: "ILS", numericCode: 376, symbol: "₪", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "New Israeli Sheqel"},
	&ISO4217Currency{alphabeticCode: "JMD", numericCode: 388, symbol: "J$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Jamaican Dollar"},
	&ISO4217Currency{alphabeticCode: "JPY", numericCode: 392, symbol: "¥", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Yen"},
	&ISO4217Currency{alphabeticCode: "JOD", numericCode: 400, symbol: "د.أ", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Jordanian Dinar"},
	&ISO4217Currency{alphabeticCode: "KZT", numericCode: 398, symbol: "₸", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Tenge"},
	&ISO4217Currency{alphabeticCode: "KES", numericCode: 404, symbol: "KSh", narrowSymbol: "Sh", smallestUnit: decimal.New(1, -2), name: "Kenyan Shilling"},
	&ISO4217Currency{alphabeticCode: "KPW", numericCode: 408, symbol: "₩", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "North Korean Won"},
	&ISO4217Currency{alphabeticCode: "KRW", numericCode: 410, symbol: "₩", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Won"},
	&ISO4217Currency{alphabeticCode: "KWD", numericCode: 414, symbol: "KD", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Kuwaiti Dinar"},
	&ISO4217Currency{alphabeticCode: "KGS", numericCode: 417, symbol: "⃀", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Som"},
	&ISO4217Currency{alphabeticCode: "LAK", numericCode: 418, symbol: "₭", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Lao Kip"},
	&ISO4217Currency{alphabeticCode: "LBP", numericCode: 422, symbol: "ل.ل", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Lebanese Pound"},
	&ISO4217Currency{alphabeticCode: "LSL", numericCode: 426, symbol: "L", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Loti"},
	&ISO4217Currency{alphabeticCode: "ZAR", numericCode: 710, symbol: "R", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Rand"},
	&ISO4217Currency{alphabeticCode: "LRD", numericCode: 430, symbol: "LD$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Liberian Dollar"},
	&ISO4217Currency{alphabeticCode: "LYD", numericCode: 434, symbol: "ل.د", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Libyan Dinar"},
	&ISO4217Currency{alphabeticCode: "CHF", numericCode: 756, symbol: "Fr.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Swiss Franc"},
	&ISO4217Currency{alphabeticCode: "MOP", numericCode: 446, symbol: "MOP$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Pataca"},
	&ISO4217Currency{alphabeticCode: "MKD", numericCode: 807, symbol: "ден", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Denar"},
	&ISO4217Currency{alphabeticCode: "MGA", numericCode: 969, symbol: "Ar", narrowSymbol: "", smallestUnit: decimal.New(2, -1), name: "Malagasy Ariary"},
	&ISO4217Currency{alphabeticCode: "MWK", numericCode: 454, symbol: "MK", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Malawi Kwacha"},
	&ISO4217Currency{alphabeticCode: "MYR", numericCode: 458, symbol: "RM", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Malaysian Ringgit"},
	&ISO4217Currency{alphabeticCode: "MVR", numericCode: 462, symbol: "Rf", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Rufiyaa"},
	&ISO4217Currency{alphabeticCode: "MRU", numericCode: 929, symbol: "UM", narrowSymbol: "", smallestUnit: decimal.New(2, -1), name: "Ouguiya"},
	&ISO4217Currency{alphabeticCode: "MUR", numericCode: 480, symbol: "Rs", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Mauritius Rupee"},
	&ISO4217Currency{alphabeticCode: "XUA", numericCode: 965, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "ADB Unit of Account"},
	&ISO4217Currency{alphabeticCode: "MXN", numericCode: 484, symbol: "$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Mexican Peso"},
	&ISO4217Currency{alphabeticCode: "MXV", numericCode: 979, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Mexican Unidad de Inversion (UDI)"},
	&ISO4217Currency{alphabeticCode: "MDL", numericCode: 498, symbol: "L", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Moldovan Leu"},
	&ISO4217Currency{alphabeticCode: "MNT", numericCode: 496, symbol: "₮", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Tugrik"},
	&ISO4217Currency{alphabeticCode: "MAD", numericCode: 504, symbol: "DH", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Moroccan Dirham"},
	&ISO4217Currency{alphabeticCode: "MZN", numericCode: 943, symbol: "MT", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Mozambique Metical"},
	&ISO4217Currency{alphabeticCode: "MMK", numericCode: 104, symbol: "K", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Kyat"},
	&ISO4217Currency{alphabeticCode: "NAD", numericCode: 516, symbol: "N$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Namibia Dollar"},
	&ISO4217Currency{alphabeticCode: "NPR", numericCode: 524, symbol: "NRs", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Nepalese Rupee"},
	&ISO4217Currency{alphabeticCode: "NIO", numericCode: 558, symbol: "C$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Cordoba Oro"},
	&ISO4217Currency{alphabeticCode: "NGN", numericCode: 566, symbol: "₦", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Naira"},
	&ISO4217Currency{alphabeticCode: "OMR", numericCode: 512, symbol: "ر.ع.", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Rial Omani"},
	&ISO4217Currency{alphabeticCode: "PKR", numericCode: 586, symbol: "Rs.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Pakistan Rupee"},
	&ISO4217Currency{alphabeticCode: "PAB", numericCode: 590, symbol: "B./", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Balboa"},
	&ISO4217Currency{alphabeticCode: "PGK", numericCode: 598, symbol: "K", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Kina"},
	&ISO4217Currency{alphabeticCode: "PYG", numericCode: 600, symbol: "₲", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Guarani"},
	&ISO4217Currency{alphabeticCode: "PEN", numericCode: 604, symbol: "S/.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Sol"},
	&ISO4217Currency{alphabeticCode: "PHP", numericCode: 608, symbol: "₱", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Philippine Peso"},
	&ISO4217Currency{alphabeticCode: "PLN", numericCode: 985, symbol: "zł", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Zloty"},
	&ISO4217Currency{alphabeticCode: "QAR", numericCode: 634, symbol: "QR", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Qatari Rial"},
	&ISO4217Currency{alphabeticCode: "RON", numericCode: 946, symbol: "L", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Romanian Leu"},
	&ISO4217Currency{alphabeticCode: "RUB", numericCode: 643, symbol: "R", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Russian Ruble"},
	&ISO4217Currency{alphabeticCode: "RWF", numericCode: 646, symbol: "RF", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Rwanda Franc"},
	&ISO4217Currency{alphabeticCode: "SHP", numericCode: 654, symbol: "£", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Saint Helena Pound"},
	&ISO4217Currency{alphabeticCode: "WST", numericCode: 882, symbol: "WS$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Tala"},
	&ISO4217Currency{alphabeticCode: "STN", numericCode: 930, symbol: "Db", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Dobra"},
	&ISO4217Currency{alphabeticCode: "SAR", numericCode: 682, symbol: "SR", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Saudi Riyal"},
	&ISO4217Currency{alphabeticCode: "RSD", numericCode: 941, symbol: "din.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Serbian Dinar"},
	&ISO4217Currency{alphabeticCode: "SCR", numericCode: 690, symbol: "SR", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Seychelles Rupee"},
	&ISO4217Currency{alphabeticCode: "SLL", numericCode: 694, symbol: "Le", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Leone"},
	&ISO4217Currency{alphabeticCode: "SGD", numericCode: 702, symbol: "S$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Singapore Dollar"},
	&ISO4217Currency{alphabeticCode: "XSU", numericCode: 994, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Sucre"},
	&ISO4217Currency{alphabeticCode: "SBD", numericCode: 90, symbol: "SI$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Solomon Islands Dollar"},
	&ISO4217Currency{alphabeticCode: "SOS", numericCode: 706, symbol: "Sh.", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Somali Shilling"},
	&ISO4217Currency{alphabeticCode: "SSP", numericCode: 728, symbol: "SS£", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "South Sudanese Pound"},
	&ISO4217Currency{alphabeticCode: "LKR", numericCode: 144, symbol: "Rs", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Sri Lanka Rupee"},
	&ISO4217Currency{alphabeticCode: "SDG", numericCode: 938, symbol: "£SD", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Sudanese Pound"},
	&ISO4217Currency{alphabeticCode: "SRD", numericCode: 968, symbol: "$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Surinam Dollar"},
	&ISO4217Currency{alphabeticCode: "SEK", numericCode: 752, symbol: "kr", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Swedish Krona"},
	&ISO4217Currency{alphabeticCode: "CHE", numericCode: 947, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "WIR Euro"},
	&ISO4217Currency{alphabeticCode: "CHW", numericCode: 948, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "WIR Franc"},
	&ISO4217Currency{alphabeticCode: "SYP", numericCode: 760, symbol: "£S", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Syrian Pound"},
	&ISO4217Currency{alphabeticCode: "TWD", numericCode: 901, symbol: "NT$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "New Taiwan Dollar"},
	&ISO4217Currency{alphabeticCode: "TJS", numericCode: 972, symbol: "SM", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Somoni"},
	&ISO4217Currency{alphabeticCode: "TZS", numericCode: 834, symbol: "TSh", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Tanzanian Shilling"},
	&ISO4217Currency{alphabeticCode: "THB", numericCode: 764, symbol: "฿", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Baht"},
	&ISO4217Currency{alphabeticCode: "TOP", numericCode: 776, symbol: "T$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Pa’anga"},
	&ISO4217Currency{alphabeticCode: "TTD", numericCode: 780, symbol: "TT$", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Trinidad and Tobago Dollar"},
	&ISO4217Currency{alphabeticCode: "TND", numericCode: 788, symbol: "DT", narrowSymbol: "", smallestUnit: decimal.New(1, -3), name: "Tunisian Dinar"},
	&ISO4217Currency{alphabeticCode: "TRY", numericCode: 949, symbol: "YTL", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Turkish Lira"},
	&ISO4217Currency{alphabeticCode: "TMT", numericCode: 934, symbol: "m", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Turkmenistan New Manat"},
	&ISO4217Currency{alphabeticCode: "UGX", numericCode: 800, symbol: "USh", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Uganda Shilling"},
	&ISO4217Currency{alphabeticCode: "UAH", numericCode: 980, symbol: "₴", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Hryvnia"},
	&ISO4217Currency{alphabeticCode: "AED", numericCode: 784, symbol: "د.إ", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "UAE Dirham"},
	&ISO4217Currency{alphabeticCode: "USN", numericCode: 997, symbol: "US$", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "US Dollar (Next day)"},
	&ISO4217Currency{alphabeticCode: "UYU", numericCode: 858, symbol: "$U", narrowSymbol: "$", smallestUnit: decimal.New(1, -2), name: "Peso Uruguayo"},
	&ISO4217Currency{alphabeticCode: "UYI", numericCode: 940, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Uruguay Peso en Unidades Indexadas (UI)"},
	&ISO4217Currency{alphabeticCode: "UYW", numericCode: 927, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -4), name: "Unidad Previsional"},
	&ISO4217Currency{alphabeticCode: "UZS", numericCode: 860, symbol: "сум", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Uzbekistan Sum"},
	&ISO4217Currency{alphabeticCode: "VUV", numericCode: 548, symbol: "VT", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Vatu"},
	&ISO4217Currency{alphabeticCode: "VES", numericCode: 928, symbol: "Bs.S", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Bolívar Soberano"},
	&ISO4217Currency{alphabeticCode: "VND", numericCode: 704, symbol: "₫", narrowSymbol: "", smallestUnit: decimal.New(1, 0), name: "Dong"},
	&ISO4217Currency{alphabeticCode: "YER", numericCode: 886, symbol: "﷼", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Yemeni Rial"},
	&ISO4217Currency{alphabeticCode: "ZMW", numericCode: 967, symbol: "ZK", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Zambian Kwacha"},
	&ISO4217Currency{alphabeticCode: "ZWL", numericCode: 932, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(1, -2), name: "Zimbabwe Dollar"},
	&ISO4217Currency{alphabeticCode: "XBA", numericCode: 955, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Bond Markets Unit European Composite Unit (EURCO)"},
	&ISO4217Currency{alphabeticCode: "XBB", numericCode: 956, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)"},
	&ISO4217Currency{alphabeticCode: "XBC", numericCode: 957, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)"},
	&ISO4217Currency{alphabeticCode: "XBD", numericCode: 958, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)"},
	&ISO4217Currency{alphabeticCode: "XTS", numericCode: 963, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Codes specifically reserved for testing purposes"},
	&ISO4217Currency{alphabeticCode: "XXX", numericCode: 999, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "The codes assigned for transactions where no currency is involved"},
	&ISO4217Currency{alphabeticCode: "XAU", numericCode: 959, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Gold"},
	&ISO4217Currency{alphabeticCode: "XPD", numericCode: 964, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Palladium"},
	&ISO4217Currency{alphabeticCode: "XPT", numericCode: 962, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Platinum"},
	&ISO4217Currency{alphabeticCode: "XAG", numericCode: 961, symbol: "", narrowSymbol: "", smallestUnit: decimal.New(0, 0), name: "Silver"},
}

// ISO4217Currencies currencies according to ISO 4217.
var ISO4217Currencies = MustNewCurrencyCollection(iso4217Name, iso4217Name, iso4217Currencies)
