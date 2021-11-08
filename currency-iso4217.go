package d3money

// iso4217Name is the name of the currency standard.
const iso4217Name = "ISO4217"

// iso4217CodePrefix is the prefix that is used to generate the unique currency code.
const iso4217CodePrefix = "ISO4217-"

// iso4217UniqueIDOffset is the offset that is used in combination with the numeric code for the unique currency ID.
const iso4217UniqueIDOffset = 42170000

// ISO4217Currency defines a currency according to the ISO 4217 standard.
type ISO4217Currency struct {
	name string // English name of the currency.

	alphabeticCode      string // Official 3 alphabetic letter code.
	numericCode         int    // Official numeric code.
	shortSymbol, symbol string // Additional representation forms. Official, but not standardized.

	decimalPlaces int
}

// Make sure this type implements the Currency interface.
var _ Currency = (*ISO4217Currency)(nil)

// Name returns the name of the currency.
func (c *ISO4217Currency) Name() string {
	return c.name
}

// StandardName returns an alphanumeric string that identifies the standard the currency is defined in.
func (c *ISO4217Currency) StandardName() string {
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
// Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC"
func (c *ISO4217Currency) UniqueCode() string {
	return c.CodePrefix() + c.alphabeticCode
}

// Code returns a string representing the currency.
// This is the official code defined by the standard, but it may not be unique across different standards.
// This may be an ISO 4217 code, depending on the currency type.
//
// Examples: "USD", "AUD", "EUR", "BTC"
func (c *ISO4217Currency) Code() string {
	return c.alphabeticCode
}

// Symbol returns a string containing the symbol of the currency.
// This may be ambiguous, and should only be used for formatting into a human readable format.
// This also doesn't follow any official standard.
//
// Examples: "US$", "AU$", "€", "₿"
func (c *ISO4217Currency) Symbol() string {
	return c.symbol
}

// ShortSymbol returns a string containing the short symbol variant of the currency.
// This may be ambiguous, and should only be used for formatting into a human readable format.
// This needs additional context when used in text output, as it doesn't differentiate between all the dollar currencies.
// This also doesn't follow any official standard.
//
// Examples: "$", "$", "€", "₿"
func (c *ISO4217Currency) ShortSymbol() string {
	if c.shortSymbol != "" {
		return c.shortSymbol
	}
	return c.symbol
}

// DecimalPlaces returns the number of decimal places that represents the "Minor unit".
// If the resulting number is 0, this currency can't be divided any further.
// If the resulting bool is false, there is no smallest unit.
func (c *ISO4217Currency) DecimalPlaces() (int, bool) {
	if c.decimalPlaces != -1 {
		return 0, false
	}
	return c.decimalPlaces, false
}

func (c *ISO4217Currency) String() string {
	return c.CodePrefix() + c.alphabeticCode
}

// TODO: Make sure (by test) that the map key is equal to the alphabetic code of the currency
// TODO: Make sure (by test) that all currencies' unique codes are unique
// TODO: Make sure (by test) that all currencies' unique IDs are unique
// TODO: Make sure (by test) that no currency code/symbol contains illegal characters

// iso4217Currencies contains the official and active ISO 4217 currencies as of August 29, 2018.
//
// Source: https://www.currency-iso.org/en/home/tables/table-a1.html
//
// This data has been modified in the following ways:
//  - Removed "ENTITY" column
//  - Removed duplicate entries (due to removal of the "ENTITY" column)
//  - Add "symbol" and "shortSymbol" columns that contain symbols which are NOT part of ISO 4217. Based on https://web.archive.org/web/20111129141202/http://fx.sauder.ubc.ca/currency_table.html and https://wikipedia.org
// TODO: Add more currency symbols, check symbols for correctness, add symbol plural variants
var iso4217Currencies = []Currency{
	&ISO4217Currency{alphabeticCode: "AFN", numericCode: 971, symbol: "؋", shortSymbol: "", decimalPlaces: 2, name: "Afghani"},
	&ISO4217Currency{alphabeticCode: "EUR", numericCode: 978, symbol: "€", shortSymbol: "", decimalPlaces: 2, name: "Euro"},
	&ISO4217Currency{alphabeticCode: "ALL", numericCode: 8, symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lek"},
	&ISO4217Currency{alphabeticCode: "DZD", numericCode: 12, symbol: "DA", shortSymbol: "", decimalPlaces: 2, name: "Algerian Dinar"},
	&ISO4217Currency{alphabeticCode: "USD", numericCode: 840, symbol: "US$", shortSymbol: "$", decimalPlaces: 2, name: "US Dollar"},
	&ISO4217Currency{alphabeticCode: "AOA", numericCode: 973, symbol: "Kz", shortSymbol: "", decimalPlaces: 2, name: "Kwanza"},
	&ISO4217Currency{alphabeticCode: "XCD", numericCode: 951, symbol: "EC$", shortSymbol: "$", decimalPlaces: 2, name: "East Caribbean Dollar"},
	&ISO4217Currency{alphabeticCode: "ARS", numericCode: 32, symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Argentine Peso"},
	&ISO4217Currency{alphabeticCode: "AMD", numericCode: 51, symbol: "֏", shortSymbol: "", decimalPlaces: 2, name: "Armenian Dram"},
	&ISO4217Currency{alphabeticCode: "AWG", numericCode: 533, symbol: "ƒ", shortSymbol: "", decimalPlaces: 2, name: "Aruban Florin"},
	&ISO4217Currency{alphabeticCode: "AUD", numericCode: 36, symbol: "AU$", shortSymbol: "$", decimalPlaces: 2, name: "Australian Dollar"},
	&ISO4217Currency{alphabeticCode: "AZN", numericCode: 944, symbol: "₼", shortSymbol: "", decimalPlaces: 2, name: "Azerbaijan Manat"},
	&ISO4217Currency{alphabeticCode: "BSD", numericCode: 44, symbol: "B$", shortSymbol: "$", decimalPlaces: 2, name: "Bahamian Dollar"},
	&ISO4217Currency{alphabeticCode: "BHD", numericCode: 48, symbol: "BD", shortSymbol: "", decimalPlaces: 3, name: "Bahraini Dinar"},
	&ISO4217Currency{alphabeticCode: "BDT", numericCode: 50, symbol: "৳", shortSymbol: "", decimalPlaces: 2, name: "Taka"},
	&ISO4217Currency{alphabeticCode: "BBD", numericCode: 52, symbol: "Bds$", shortSymbol: "$", decimalPlaces: 2, name: "Barbados Dollar"},
	&ISO4217Currency{alphabeticCode: "BYN", numericCode: 933, symbol: "Br", shortSymbol: "", decimalPlaces: 2, name: "Belarusian Ruble"},
	&ISO4217Currency{alphabeticCode: "BZD", numericCode: 84, symbol: "BZ$", shortSymbol: "$", decimalPlaces: 2, name: "Belize Dollar"},
	&ISO4217Currency{alphabeticCode: "XOF", numericCode: 952, symbol: "CFA", shortSymbol: "Fr", decimalPlaces: 0, name: "CFA Franc BCEAO"},
	&ISO4217Currency{alphabeticCode: "BMD", numericCode: 60, symbol: "BD$", shortSymbol: "$", decimalPlaces: 2, name: "Bermudian Dollar"},
	&ISO4217Currency{alphabeticCode: "INR", numericCode: 356, symbol: "₹", shortSymbol: "", decimalPlaces: 2, name: "Indian Rupee"},
	&ISO4217Currency{alphabeticCode: "BTN", numericCode: 64, symbol: "Nu.", shortSymbol: "", decimalPlaces: 2, name: "Ngultrum"},
	&ISO4217Currency{alphabeticCode: "BOB", numericCode: 68, symbol: "Bs.", shortSymbol: "", decimalPlaces: 2, name: "Boliviano"},
	&ISO4217Currency{alphabeticCode: "BOV", numericCode: 984, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Mvdol"},
	&ISO4217Currency{alphabeticCode: "BAM", numericCode: 977, symbol: "KM", shortSymbol: "", decimalPlaces: 2, name: "Convertible Mark"},
	&ISO4217Currency{alphabeticCode: "BWP", numericCode: 72, symbol: "P", shortSymbol: "", decimalPlaces: 2, name: "Pula"},
	&ISO4217Currency{alphabeticCode: "NOK", numericCode: 578, symbol: "kr", shortSymbol: "", decimalPlaces: 2, name: "Norwegian Krone"},
	&ISO4217Currency{alphabeticCode: "BRL", numericCode: 986, symbol: "R$", shortSymbol: "", decimalPlaces: 2, name: "Brazilian Real"},
	&ISO4217Currency{alphabeticCode: "BND", numericCode: 96, symbol: "B$", shortSymbol: "$", decimalPlaces: 2, name: "Brunei Dollar"},
	&ISO4217Currency{alphabeticCode: "BGN", numericCode: 975, symbol: "лв.", shortSymbol: "", decimalPlaces: 2, name: "Bulgarian Lev"},
	&ISO4217Currency{alphabeticCode: "BIF", numericCode: 108, symbol: "FBu", shortSymbol: "", decimalPlaces: 0, name: "Burundi Franc"},
	&ISO4217Currency{alphabeticCode: "CVE", numericCode: 132, symbol: "Esc", shortSymbol: "", decimalPlaces: 2, name: "Cabo Verde Escudo"},
	&ISO4217Currency{alphabeticCode: "KHR", numericCode: 116, symbol: "៛", shortSymbol: "", decimalPlaces: 2, name: "Riel"},
	&ISO4217Currency{alphabeticCode: "XAF", numericCode: 950, symbol: "CFA", shortSymbol: "Fr", decimalPlaces: 0, name: "CFA Franc BEAC"},
	&ISO4217Currency{alphabeticCode: "CAD", numericCode: 124, symbol: "CA$", shortSymbol: "$", decimalPlaces: 2, name: "Canadian Dollar"},
	&ISO4217Currency{alphabeticCode: "KYD", numericCode: 136, symbol: "KY$", shortSymbol: "", decimalPlaces: 2, name: "Cayman Islands Dollar"},
	&ISO4217Currency{alphabeticCode: "CLP", numericCode: 152, symbol: "CLP$", shortSymbol: "$", decimalPlaces: 0, name: "Chilean Peso"},
	&ISO4217Currency{alphabeticCode: "CLF", numericCode: 990, symbol: "", shortSymbol: "", decimalPlaces: 4, name: "Unidad de Fomento"},
	&ISO4217Currency{alphabeticCode: "CNY", numericCode: 156, symbol: "¥", shortSymbol: "", decimalPlaces: 2, name: "Yuan Renminbi"},
	&ISO4217Currency{alphabeticCode: "COP", numericCode: 170, symbol: "Col$", shortSymbol: "$", decimalPlaces: 2, name: "Colombian Peso"},
	&ISO4217Currency{alphabeticCode: "COU", numericCode: 970, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Unidad de Valor Real"},
	&ISO4217Currency{alphabeticCode: "KMF", numericCode: 174, symbol: "CF", shortSymbol: "Fr", decimalPlaces: 0, name: "Comorian Franc "},
	&ISO4217Currency{alphabeticCode: "CDF", numericCode: 976, symbol: "F", shortSymbol: "", decimalPlaces: 2, name: "Congolese Franc"},
	&ISO4217Currency{alphabeticCode: "NZD", numericCode: 554, symbol: "NZ$", shortSymbol: "$", decimalPlaces: 2, name: "New Zealand Dollar"},
	&ISO4217Currency{alphabeticCode: "CRC", numericCode: 188, symbol: "₡", shortSymbol: "", decimalPlaces: 2, name: "Costa Rican Colon"},
	&ISO4217Currency{alphabeticCode: "HRK", numericCode: 191, symbol: "kn", shortSymbol: "", decimalPlaces: 2, name: "Kuna"},
	&ISO4217Currency{alphabeticCode: "CUP", numericCode: 192, symbol: "₱", shortSymbol: "", decimalPlaces: 2, name: "Cuban Peso"},
	&ISO4217Currency{alphabeticCode: "CUC", numericCode: 931, symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Peso Convertible"},
	&ISO4217Currency{alphabeticCode: "ANG", numericCode: 532, symbol: "NAƒ", shortSymbol: "", decimalPlaces: 2, name: "Netherlands Antillean Guilder"},
	&ISO4217Currency{alphabeticCode: "CZK", numericCode: 203, symbol: "Kč", shortSymbol: "", decimalPlaces: 2, name: "Czech Koruna"},
	&ISO4217Currency{alphabeticCode: "DKK", numericCode: 208, symbol: "Kr", shortSymbol: "", decimalPlaces: 2, name: "Danish Krone"},
	&ISO4217Currency{alphabeticCode: "DJF", numericCode: 262, symbol: "Fdj", shortSymbol: "", decimalPlaces: 0, name: "Djibouti Franc"},
	&ISO4217Currency{alphabeticCode: "DOP", numericCode: 214, symbol: "RD$", shortSymbol: "$", decimalPlaces: 2, name: "Dominican Peso"},
	&ISO4217Currency{alphabeticCode: "EGP", numericCode: 818, symbol: "E£", shortSymbol: "£", decimalPlaces: 2, name: "Egyptian Pound"},
	&ISO4217Currency{alphabeticCode: "SVC", numericCode: 222, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "El Salvador Colon"},
	&ISO4217Currency{alphabeticCode: "ERN", numericCode: 232, symbol: "Nkf", shortSymbol: "", decimalPlaces: 2, name: "Nakfa"},
	&ISO4217Currency{alphabeticCode: "SZL", numericCode: 748, symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lilangeni"},
	&ISO4217Currency{alphabeticCode: "ETB", numericCode: 230, symbol: "Br", shortSymbol: "", decimalPlaces: 2, name: "Ethiopian Birr"},
	&ISO4217Currency{alphabeticCode: "FKP", numericCode: 238, symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Falkland Islands Pound"},
	&ISO4217Currency{alphabeticCode: "FJD", numericCode: 242, symbol: "FJ$", shortSymbol: "$", decimalPlaces: 2, name: "Fiji Dollar"},
	&ISO4217Currency{alphabeticCode: "XPF", numericCode: 953, symbol: "₣", shortSymbol: "", decimalPlaces: 0, name: "CFP Franc"},
	&ISO4217Currency{alphabeticCode: "GMD", numericCode: 270, symbol: "D", shortSymbol: "", decimalPlaces: 2, name: "Dalasi"},
	&ISO4217Currency{alphabeticCode: "GEL", numericCode: 981, symbol: "₾", shortSymbol: "", decimalPlaces: 2, name: "Lari"},
	&ISO4217Currency{alphabeticCode: "GHS", numericCode: 936, symbol: "₵", shortSymbol: "", decimalPlaces: 2, name: "Ghana Cedi"},
	&ISO4217Currency{alphabeticCode: "GIP", numericCode: 292, symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Gibraltar Pound"},
	&ISO4217Currency{alphabeticCode: "GTQ", numericCode: 320, symbol: "Q", shortSymbol: "", decimalPlaces: 2, name: "Quetzal"},
	&ISO4217Currency{alphabeticCode: "GBP", numericCode: 826, symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Pound Sterling"},
	&ISO4217Currency{alphabeticCode: "GNF", numericCode: 324, symbol: "FG", shortSymbol: "", decimalPlaces: 0, name: "Guinean Franc"},
	&ISO4217Currency{alphabeticCode: "GYD", numericCode: 328, symbol: "GY$", shortSymbol: "", decimalPlaces: 2, name: "Guyana Dollar"},
	&ISO4217Currency{alphabeticCode: "HTG", numericCode: 332, symbol: "G", shortSymbol: "", decimalPlaces: 2, name: "Gourde"},
	&ISO4217Currency{alphabeticCode: "HNL", numericCode: 340, symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lempira"},
	&ISO4217Currency{alphabeticCode: "HKD", numericCode: 344, symbol: "HK$", shortSymbol: "", decimalPlaces: 2, name: "Hong Kong Dollar"},
	&ISO4217Currency{alphabeticCode: "HUF", numericCode: 348, symbol: "Ft", shortSymbol: "", decimalPlaces: 2, name: "Forint"},
	&ISO4217Currency{alphabeticCode: "ISK", numericCode: 352, symbol: "kr", shortSymbol: "", decimalPlaces: 0, name: "Iceland Krona"},
	&ISO4217Currency{alphabeticCode: "IDR", numericCode: 360, symbol: "Rp", shortSymbol: "", decimalPlaces: 2, name: "Rupiah"},
	&ISO4217Currency{alphabeticCode: "XDR", numericCode: 960, symbol: "SDR", shortSymbol: "", decimalPlaces: -1, name: "SDR (Special Drawing Right)"},
	&ISO4217Currency{alphabeticCode: "IRR", numericCode: 364, symbol: "﷼", shortSymbol: "", decimalPlaces: 2, name: "Iranian Rial"},
	&ISO4217Currency{alphabeticCode: "IQD", numericCode: 368, symbol: "د.ع", shortSymbol: "", decimalPlaces: 3, name: "Iraqi Dinar"},
	&ISO4217Currency{alphabeticCode: "ILS", numericCode: 376, symbol: "₪", shortSymbol: "", decimalPlaces: 2, name: "New Israeli Sheqel"},
	&ISO4217Currency{alphabeticCode: "JMD", numericCode: 388, symbol: "J$", shortSymbol: "", decimalPlaces: 2, name: "Jamaican Dollar"},
	&ISO4217Currency{alphabeticCode: "JPY", numericCode: 392, symbol: "¥", shortSymbol: "", decimalPlaces: 0, name: "Yen"},
	&ISO4217Currency{alphabeticCode: "JOD", numericCode: 400, symbol: "د.أ", shortSymbol: "", decimalPlaces: 3, name: "Jordanian Dinar"},
	&ISO4217Currency{alphabeticCode: "KZT", numericCode: 398, symbol: "₸", shortSymbol: "", decimalPlaces: 2, name: "Tenge"},
	&ISO4217Currency{alphabeticCode: "KES", numericCode: 404, symbol: "KSh", shortSymbol: "Sh", decimalPlaces: 2, name: "Kenyan Shilling"},
	&ISO4217Currency{alphabeticCode: "KPW", numericCode: 408, symbol: "₩", shortSymbol: "", decimalPlaces: 2, name: "North Korean Won"},
	&ISO4217Currency{alphabeticCode: "KRW", numericCode: 410, symbol: "₩", shortSymbol: "", decimalPlaces: 0, name: "Won"},
	&ISO4217Currency{alphabeticCode: "KWD", numericCode: 414, symbol: "KD", shortSymbol: "", decimalPlaces: 3, name: "Kuwaiti Dinar"},
	&ISO4217Currency{alphabeticCode: "KGS", numericCode: 417, symbol: "⃀", shortSymbol: "", decimalPlaces: 2, name: "Som"},
	&ISO4217Currency{alphabeticCode: "LAK", numericCode: 418, symbol: "₭", shortSymbol: "", decimalPlaces: 2, name: "Lao Kip"},
	&ISO4217Currency{alphabeticCode: "LBP", numericCode: 422, symbol: "ل.ل", shortSymbol: "", decimalPlaces: 2, name: "Lebanese Pound"},
	&ISO4217Currency{alphabeticCode: "LSL", numericCode: 426, symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Loti"},
	&ISO4217Currency{alphabeticCode: "ZAR", numericCode: 710, symbol: "R", shortSymbol: "", decimalPlaces: 2, name: "Rand"},
	&ISO4217Currency{alphabeticCode: "LRD", numericCode: 430, symbol: "LD$", shortSymbol: "$", decimalPlaces: 2, name: "Liberian Dollar"},
	&ISO4217Currency{alphabeticCode: "LYD", numericCode: 434, symbol: "ل.د", shortSymbol: "", decimalPlaces: 3, name: "Libyan Dinar"},
	&ISO4217Currency{alphabeticCode: "CHF", numericCode: 756, symbol: "Fr.", shortSymbol: "", decimalPlaces: 2, name: "Swiss Franc"},
	&ISO4217Currency{alphabeticCode: "MOP", numericCode: 446, symbol: "MOP$", shortSymbol: "", decimalPlaces: 2, name: "Pataca"},
	&ISO4217Currency{alphabeticCode: "MKD", numericCode: 807, symbol: "ден", shortSymbol: "", decimalPlaces: 2, name: "Denar"},
	&ISO4217Currency{alphabeticCode: "MGA", numericCode: 969, symbol: "Ar", shortSymbol: "", decimalPlaces: 2, name: "Malagasy Ariary"},
	&ISO4217Currency{alphabeticCode: "MWK", numericCode: 454, symbol: "MK", shortSymbol: "", decimalPlaces: 2, name: "Malawi Kwacha"},
	&ISO4217Currency{alphabeticCode: "MYR", numericCode: 458, symbol: "RM", shortSymbol: "", decimalPlaces: 2, name: "Malaysian Ringgit"},
	&ISO4217Currency{alphabeticCode: "MVR", numericCode: 462, symbol: "Rf", shortSymbol: "", decimalPlaces: 2, name: "Rufiyaa"},
	&ISO4217Currency{alphabeticCode: "MRU", numericCode: 929, symbol: "UM", shortSymbol: "", decimalPlaces: 2, name: "Ouguiya"},
	&ISO4217Currency{alphabeticCode: "MUR", numericCode: 480, symbol: "Rs", shortSymbol: "", decimalPlaces: 2, name: "Mauritius Rupee"},
	&ISO4217Currency{alphabeticCode: "XUA", numericCode: 965, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "ADB Unit of Account"},
	&ISO4217Currency{alphabeticCode: "MXN", numericCode: 484, symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Mexican Peso"},
	&ISO4217Currency{alphabeticCode: "MXV", numericCode: 979, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Mexican Unidad de Inversion (UDI)"},
	&ISO4217Currency{alphabeticCode: "MDL", numericCode: 498, symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Moldovan Leu"},
	&ISO4217Currency{alphabeticCode: "MNT", numericCode: 496, symbol: "₮", shortSymbol: "", decimalPlaces: 2, name: "Tugrik"},
	&ISO4217Currency{alphabeticCode: "MAD", numericCode: 504, symbol: "DH", shortSymbol: "", decimalPlaces: 2, name: "Moroccan Dirham"},
	&ISO4217Currency{alphabeticCode: "MZN", numericCode: 943, symbol: "MT", shortSymbol: "", decimalPlaces: 2, name: "Mozambique Metical"},
	&ISO4217Currency{alphabeticCode: "MMK", numericCode: 104, symbol: "K", shortSymbol: "", decimalPlaces: 2, name: "Kyat"},
	&ISO4217Currency{alphabeticCode: "NAD", numericCode: 516, symbol: "N$", shortSymbol: "", decimalPlaces: 2, name: "Namibia Dollar"},
	&ISO4217Currency{alphabeticCode: "NPR", numericCode: 524, symbol: "NRs", shortSymbol: "", decimalPlaces: 2, name: "Nepalese Rupee"},
	&ISO4217Currency{alphabeticCode: "NIO", numericCode: 558, symbol: "C$", shortSymbol: "", decimalPlaces: 2, name: "Cordoba Oro"},
	&ISO4217Currency{alphabeticCode: "NGN", numericCode: 566, symbol: "₦", shortSymbol: "", decimalPlaces: 2, name: "Naira"},
	&ISO4217Currency{alphabeticCode: "OMR", numericCode: 512, symbol: "ر.ع.", shortSymbol: "", decimalPlaces: 3, name: "Rial Omani"},
	&ISO4217Currency{alphabeticCode: "PKR", numericCode: 586, symbol: "Rs.", shortSymbol: "", decimalPlaces: 2, name: "Pakistan Rupee"},
	&ISO4217Currency{alphabeticCode: "PAB", numericCode: 590, symbol: "B./", shortSymbol: "", decimalPlaces: 2, name: "Balboa"},
	&ISO4217Currency{alphabeticCode: "PGK", numericCode: 598, symbol: "K", shortSymbol: "", decimalPlaces: 2, name: "Kina"},
	&ISO4217Currency{alphabeticCode: "PYG", numericCode: 600, symbol: "₲", shortSymbol: "", decimalPlaces: 0, name: "Guarani"},
	&ISO4217Currency{alphabeticCode: "PEN", numericCode: 604, symbol: "S/.", shortSymbol: "", decimalPlaces: 2, name: "Sol"},
	&ISO4217Currency{alphabeticCode: "PHP", numericCode: 608, symbol: "₱", shortSymbol: "", decimalPlaces: 2, name: "Philippine Peso"},
	&ISO4217Currency{alphabeticCode: "PLN", numericCode: 985, symbol: "zł", shortSymbol: "", decimalPlaces: 2, name: "Zloty"},
	&ISO4217Currency{alphabeticCode: "QAR", numericCode: 634, symbol: "QR", shortSymbol: "", decimalPlaces: 2, name: "Qatari Rial"},
	&ISO4217Currency{alphabeticCode: "RON", numericCode: 946, symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Romanian Leu"},
	&ISO4217Currency{alphabeticCode: "RUB", numericCode: 643, symbol: "R", shortSymbol: "", decimalPlaces: 2, name: "Russian Ruble"},
	&ISO4217Currency{alphabeticCode: "RWF", numericCode: 646, symbol: "RF", shortSymbol: "", decimalPlaces: 0, name: "Rwanda Franc"},
	&ISO4217Currency{alphabeticCode: "SHP", numericCode: 654, symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Saint Helena Pound"},
	&ISO4217Currency{alphabeticCode: "WST", numericCode: 882, symbol: "WS$", shortSymbol: "", decimalPlaces: 2, name: "Tala"},
	&ISO4217Currency{alphabeticCode: "STN", numericCode: 930, symbol: "Db", shortSymbol: "", decimalPlaces: 2, name: "Dobra"},
	&ISO4217Currency{alphabeticCode: "SAR", numericCode: 682, symbol: "SR", shortSymbol: "", decimalPlaces: 2, name: "Saudi Riyal"},
	&ISO4217Currency{alphabeticCode: "RSD", numericCode: 941, symbol: "din.", shortSymbol: "", decimalPlaces: 2, name: "Serbian Dinar"},
	&ISO4217Currency{alphabeticCode: "SCR", numericCode: 690, symbol: "SR", shortSymbol: "", decimalPlaces: 2, name: "Seychelles Rupee"},
	&ISO4217Currency{alphabeticCode: "SLL", numericCode: 694, symbol: "Le", shortSymbol: "", decimalPlaces: 2, name: "Leone"},
	&ISO4217Currency{alphabeticCode: "SGD", numericCode: 702, symbol: "S$", shortSymbol: "", decimalPlaces: 2, name: "Singapore Dollar"},
	&ISO4217Currency{alphabeticCode: "XSU", numericCode: 994, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Sucre"},
	&ISO4217Currency{alphabeticCode: "SBD", numericCode: 90, symbol: "SI$", shortSymbol: "", decimalPlaces: 2, name: "Solomon Islands Dollar"},
	&ISO4217Currency{alphabeticCode: "SOS", numericCode: 706, symbol: "Sh.", shortSymbol: "", decimalPlaces: 2, name: "Somali Shilling"},
	&ISO4217Currency{alphabeticCode: "SSP", numericCode: 728, symbol: "SS£", shortSymbol: "", decimalPlaces: 2, name: "South Sudanese Pound"},
	&ISO4217Currency{alphabeticCode: "LKR", numericCode: 144, symbol: "Rs", shortSymbol: "", decimalPlaces: 2, name: "Sri Lanka Rupee"},
	&ISO4217Currency{alphabeticCode: "SDG", numericCode: 938, symbol: "£SD", shortSymbol: "", decimalPlaces: 2, name: "Sudanese Pound"},
	&ISO4217Currency{alphabeticCode: "SRD", numericCode: 968, symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Surinam Dollar"},
	&ISO4217Currency{alphabeticCode: "SEK", numericCode: 752, symbol: "kr", shortSymbol: "", decimalPlaces: 2, name: "Swedish Krona"},
	&ISO4217Currency{alphabeticCode: "CHE", numericCode: 947, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "WIR Euro"},
	&ISO4217Currency{alphabeticCode: "CHW", numericCode: 948, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "WIR Franc"},
	&ISO4217Currency{alphabeticCode: "SYP", numericCode: 760, symbol: "£S", shortSymbol: "", decimalPlaces: 2, name: "Syrian Pound"},
	&ISO4217Currency{alphabeticCode: "TWD", numericCode: 901, symbol: "NT$", shortSymbol: "$", decimalPlaces: 2, name: "New Taiwan Dollar"},
	&ISO4217Currency{alphabeticCode: "TJS", numericCode: 972, symbol: "SM", shortSymbol: "", decimalPlaces: 2, name: "Somoni"},
	&ISO4217Currency{alphabeticCode: "TZS", numericCode: 834, symbol: "TSh", shortSymbol: "", decimalPlaces: 2, name: "Tanzanian Shilling"},
	&ISO4217Currency{alphabeticCode: "THB", numericCode: 764, symbol: "฿", shortSymbol: "", decimalPlaces: 2, name: "Baht"},
	&ISO4217Currency{alphabeticCode: "TOP", numericCode: 776, symbol: "T$", shortSymbol: "", decimalPlaces: 2, name: "Pa’anga"},
	&ISO4217Currency{alphabeticCode: "TTD", numericCode: 780, symbol: "TT$", shortSymbol: "", decimalPlaces: 2, name: "Trinidad and Tobago Dollar"},
	&ISO4217Currency{alphabeticCode: "TND", numericCode: 788, symbol: "DT", shortSymbol: "", decimalPlaces: 3, name: "Tunisian Dinar"},
	&ISO4217Currency{alphabeticCode: "TRY", numericCode: 949, symbol: "YTL", shortSymbol: "", decimalPlaces: 2, name: "Turkish Lira"},
	&ISO4217Currency{alphabeticCode: "TMT", numericCode: 934, symbol: "m", shortSymbol: "", decimalPlaces: 2, name: "Turkmenistan New Manat"},
	&ISO4217Currency{alphabeticCode: "UGX", numericCode: 800, symbol: "USh", shortSymbol: "", decimalPlaces: 0, name: "Uganda Shilling"},
	&ISO4217Currency{alphabeticCode: "UAH", numericCode: 980, symbol: "₴", shortSymbol: "", decimalPlaces: 2, name: "Hryvnia"},
	&ISO4217Currency{alphabeticCode: "AED", numericCode: 784, symbol: "د.إ", shortSymbol: "", decimalPlaces: 2, name: "UAE Dirham"},
	&ISO4217Currency{alphabeticCode: "USN", numericCode: 997, symbol: "US$", shortSymbol: "$", decimalPlaces: 2, name: "US Dollar (Next day)"},
	&ISO4217Currency{alphabeticCode: "UYU", numericCode: 858, symbol: "$U", shortSymbol: "$", decimalPlaces: 2, name: "Peso Uruguayo"},
	&ISO4217Currency{alphabeticCode: "UYI", numericCode: 940, symbol: "", shortSymbol: "", decimalPlaces: 0, name: "Uruguay Peso en Unidades Indexadas (UI)"},
	&ISO4217Currency{alphabeticCode: "UYW", numericCode: 927, symbol: "", shortSymbol: "", decimalPlaces: 4, name: "Unidad Previsional"},
	&ISO4217Currency{alphabeticCode: "UZS", numericCode: 860, symbol: "сум", shortSymbol: "", decimalPlaces: 2, name: "Uzbekistan Sum"},
	&ISO4217Currency{alphabeticCode: "VUV", numericCode: 548, symbol: "VT", shortSymbol: "", decimalPlaces: 0, name: "Vatu"},
	&ISO4217Currency{alphabeticCode: "VES", numericCode: 928, symbol: "Bs.S", shortSymbol: "", decimalPlaces: 2, name: "Bolívar Soberano"},
	&ISO4217Currency{alphabeticCode: "VND", numericCode: 704, symbol: "₫", shortSymbol: "", decimalPlaces: 0, name: "Dong"},
	&ISO4217Currency{alphabeticCode: "YER", numericCode: 886, symbol: "﷼", shortSymbol: "", decimalPlaces: 2, name: "Yemeni Rial"},
	&ISO4217Currency{alphabeticCode: "ZMW", numericCode: 967, symbol: "ZK", shortSymbol: "", decimalPlaces: 2, name: "Zambian Kwacha"},
	&ISO4217Currency{alphabeticCode: "ZWL", numericCode: 932, symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Zimbabwe Dollar"},
	&ISO4217Currency{alphabeticCode: "XBA", numericCode: 955, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Composite Unit (EURCO)"},
	&ISO4217Currency{alphabeticCode: "XBB", numericCode: 956, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)"},
	&ISO4217Currency{alphabeticCode: "XBC", numericCode: 957, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)"},
	&ISO4217Currency{alphabeticCode: "XBD", numericCode: 958, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)"},
	&ISO4217Currency{alphabeticCode: "XTS", numericCode: 963, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Codes specifically reserved for testing purposes"},
	&ISO4217Currency{alphabeticCode: "XXX", numericCode: 999, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "The codes assigned for transactions where no currency is involved"},
	&ISO4217Currency{alphabeticCode: "XAU", numericCode: 959, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Gold"},
	&ISO4217Currency{alphabeticCode: "XPD", numericCode: 964, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Palladium"},
	&ISO4217Currency{alphabeticCode: "XPT", numericCode: 962, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Platinum"},
	&ISO4217Currency{alphabeticCode: "XAG", numericCode: 961, symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Silver"},
}

// ISO4217Currencies currencies according to ISO 4217.
var ISO4217Currencies = MustNewCurrencyCollection(iso4217Name, iso4217Currencies)
