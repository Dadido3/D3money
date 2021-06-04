package d3money

// ISO4217 defines a currency according to the ISO 4217 standard.
type ISO4217 struct {
	name string // English name of the currency.

	alphabeticCode      string // Official 3 alphabetic letter code.
	numericCode         string // Official numeric code as string.
	shortSymbol, symbol string // Additional representation forms. Official, but not standardized.

	decimalPlaces int
}

// Make sure if this type implements the Currency interface.
var _ Currency = (*ISO4217)(nil)

// Name returns the name of the currency.
func (c *ISO4217) Name() string {
	return c.name
}

// Code returns a string representing the currency.
// This representation is unique across different currency standards.
//
// Examples: "ISO4217-USD", "ISO4217-AUD", "ISO4217-EUR", "CRYPTO-BTC"
func (c *ISO4217) UniqueCode() string {
	return c.Standard() + "-" + c.alphabeticCode
}

// Code returns a string representing the currency.
// This is the official code defined by the standard, but it may not be unique across different standards.
// This may be an ISO 4217 code, depending on the currency type.
//
// Examples: "USD", "AUD", "EUR", "BTC"
func (c *ISO4217) Code() string {
	return c.alphabeticCode
}

// Symbol returns a string containing the symbol of the currency.
// This may be ambiguous, and should only be used for formatting into a human readable format.
// This also doesn't follow any official standard.
//
// Examples: "US$", "AU$", "€", "₿"
func (c *ISO4217) Symbol() string {
	return c.symbol
}

// ShortSymbol returns a string containing the short symbol variant of the currency.
// This may be ambiguous, and should only be used for formatting into a human readable format.
// This needs additional context when used in text output, as it doesn't differentiate between all the dollar currencies.
// This also doesn't follow any official standard.
//
// Examples: "$", "$", "€", "₿"
func (c *ISO4217) ShortSymbol() string {
	if c.shortSymbol != "" {
		return c.shortSymbol
	}
	return c.symbol
}

// DecimalPlaces returns the number of decimal places that represents the "Minor unit".
// If the resulting number is 0, this currency can't be divided any further.
// If the resulting bool is false, there is no smallest unit.
func (c *ISO4217) DecimalPlaces() (int, bool) {
	if c.decimalPlaces != -1 {
		return 0, false
	}
	return c.decimalPlaces, false
}

// Standard returns an alphanumeric string that identifies the standard the currency is defined in.
func (c *ISO4217) Standard() string {
	return "ISO4217"
}

func (c *ISO4217) String() string {
	return c.Standard() + "-" + c.alphabeticCode
}

// TODO: Make sure (by test) that the map key is equal to the alphabetic code of the currency
// TODO: Make sure (by test) that all currencies unique codes are unique
// TODO: Make sure (by test) that no currency code/symbol contains illegal characters

// ISO4217Currencies contains the official and active ISO 4217 currencies as of August 29, 2018.
//
// Source: https://www.currency-iso.org/en/home/tables/table-a1.html
//
// This data has been modified in the following ways:
//  - Removed "ENTITY" column
//  - Removed duplicate entries (due to removal of the "ENTITY" column)
//  - Add "symbol" and "shortSymbol" columns that contain symbols which are NOT part of ISO 4217. Based on https://web.archive.org/web/20111129141202/http://fx.sauder.ubc.ca/currency_table.html and https://wikipedia.org
// TODO: Add more currency symbols, check symbols for correctness, add symbol plural variants
var ISO4217Currencies = map[string]Currency{
	"AFN": &ISO4217{alphabeticCode: "AFN", numericCode: "971", symbol: "؋", shortSymbol: "", decimalPlaces: 2, name: "Afghani"},
	"EUR": &ISO4217{alphabeticCode: "EUR", numericCode: "978", symbol: "€", shortSymbol: "", decimalPlaces: 2, name: "Euro"},
	"ALL": &ISO4217{alphabeticCode: "ALL", numericCode: "008", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lek"},
	"DZD": &ISO4217{alphabeticCode: "DZD", numericCode: "012", symbol: "DA", shortSymbol: "", decimalPlaces: 2, name: "Algerian Dinar"},
	"USD": &ISO4217{alphabeticCode: "USD", numericCode: "840", symbol: "US$", shortSymbol: "$", decimalPlaces: 2, name: "US Dollar"},
	"AOA": &ISO4217{alphabeticCode: "AOA", numericCode: "973", symbol: "Kz", shortSymbol: "", decimalPlaces: 2, name: "Kwanza"},
	"XCD": &ISO4217{alphabeticCode: "XCD", numericCode: "951", symbol: "EC$", shortSymbol: "$", decimalPlaces: 2, name: "East Caribbean Dollar"},
	"ARS": &ISO4217{alphabeticCode: "ARS", numericCode: "032", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Argentine Peso"},
	"AMD": &ISO4217{alphabeticCode: "AMD", numericCode: "051", symbol: "֏", shortSymbol: "", decimalPlaces: 2, name: "Armenian Dram"},
	"AWG": &ISO4217{alphabeticCode: "AWG", numericCode: "533", symbol: "ƒ", shortSymbol: "", decimalPlaces: 2, name: "Aruban Florin"},
	"AUD": &ISO4217{alphabeticCode: "AUD", numericCode: "036", symbol: "AU$", shortSymbol: "$", decimalPlaces: 2, name: "Australian Dollar"},
	"AZN": &ISO4217{alphabeticCode: "AZN", numericCode: "944", symbol: "₼", shortSymbol: "", decimalPlaces: 2, name: "Azerbaijan Manat"},
	"BSD": &ISO4217{alphabeticCode: "BSD", numericCode: "044", symbol: "B$", shortSymbol: "$", decimalPlaces: 2, name: "Bahamian Dollar"},
	"BHD": &ISO4217{alphabeticCode: "BHD", numericCode: "048", symbol: "BD", shortSymbol: "", decimalPlaces: 3, name: "Bahraini Dinar"},
	"BDT": &ISO4217{alphabeticCode: "BDT", numericCode: "050", symbol: "৳", shortSymbol: "", decimalPlaces: 2, name: "Taka"},
	"BBD": &ISO4217{alphabeticCode: "BBD", numericCode: "052", symbol: "Bds$", shortSymbol: "$", decimalPlaces: 2, name: "Barbados Dollar"},
	"BYN": &ISO4217{alphabeticCode: "BYN", numericCode: "933", symbol: "Br", shortSymbol: "", decimalPlaces: 2, name: "Belarusian Ruble"},
	"BZD": &ISO4217{alphabeticCode: "BZD", numericCode: "084", symbol: "BZ$", shortSymbol: "$", decimalPlaces: 2, name: "Belize Dollar"},
	"XOF": &ISO4217{alphabeticCode: "XOF", numericCode: "952", symbol: "CFA", shortSymbol: "Fr", decimalPlaces: 0, name: "CFA Franc BCEAO"},
	"BMD": &ISO4217{alphabeticCode: "BMD", numericCode: "060", symbol: "BD$", shortSymbol: "$", decimalPlaces: 2, name: "Bermudian Dollar"},
	"INR": &ISO4217{alphabeticCode: "INR", numericCode: "356", symbol: "₹", shortSymbol: "", decimalPlaces: 2, name: "Indian Rupee"},
	"BTN": &ISO4217{alphabeticCode: "BTN", numericCode: "064", symbol: "Nu.", shortSymbol: "", decimalPlaces: 2, name: "Ngultrum"},
	"BOB": &ISO4217{alphabeticCode: "BOB", numericCode: "068", symbol: "Bs.", shortSymbol: "", decimalPlaces: 2, name: "Boliviano"},
	"BOV": &ISO4217{alphabeticCode: "BOV", numericCode: "984", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Mvdol"},
	"BAM": &ISO4217{alphabeticCode: "BAM", numericCode: "977", symbol: "KM", shortSymbol: "", decimalPlaces: 2, name: "Convertible Mark"},
	"BWP": &ISO4217{alphabeticCode: "BWP", numericCode: "072", symbol: "P", shortSymbol: "", decimalPlaces: 2, name: "Pula"},
	"NOK": &ISO4217{alphabeticCode: "NOK", numericCode: "578", symbol: "kr", shortSymbol: "", decimalPlaces: 2, name: "Norwegian Krone"},
	"BRL": &ISO4217{alphabeticCode: "BRL", numericCode: "986", symbol: "R$", shortSymbol: "", decimalPlaces: 2, name: "Brazilian Real"},
	"BND": &ISO4217{alphabeticCode: "BND", numericCode: "096", symbol: "B$", shortSymbol: "$", decimalPlaces: 2, name: "Brunei Dollar"},
	"BGN": &ISO4217{alphabeticCode: "BGN", numericCode: "975", symbol: "лв.", shortSymbol: "", decimalPlaces: 2, name: "Bulgarian Lev"},
	"BIF": &ISO4217{alphabeticCode: "BIF", numericCode: "108", symbol: "FBu", shortSymbol: "", decimalPlaces: 0, name: "Burundi Franc"},
	"CVE": &ISO4217{alphabeticCode: "CVE", numericCode: "132", symbol: "Esc", shortSymbol: "", decimalPlaces: 2, name: "Cabo Verde Escudo"},
	"KHR": &ISO4217{alphabeticCode: "KHR", numericCode: "116", symbol: "៛", shortSymbol: "", decimalPlaces: 2, name: "Riel"},
	"XAF": &ISO4217{alphabeticCode: "XAF", numericCode: "950", symbol: "CFA", shortSymbol: "Fr", decimalPlaces: 0, name: "CFA Franc BEAC"},
	"CAD": &ISO4217{alphabeticCode: "CAD", numericCode: "124", symbol: "CA$", shortSymbol: "$", decimalPlaces: 2, name: "Canadian Dollar"},
	"KYD": &ISO4217{alphabeticCode: "KYD", numericCode: "136", symbol: "KY$", shortSymbol: "", decimalPlaces: 2, name: "Cayman Islands Dollar"},
	"CLP": &ISO4217{alphabeticCode: "CLP", numericCode: "152", symbol: "CLP$", shortSymbol: "$", decimalPlaces: 0, name: "Chilean Peso"},
	"CLF": &ISO4217{alphabeticCode: "CLF", numericCode: "990", symbol: "", shortSymbol: "", decimalPlaces: 4, name: "Unidad de Fomento"},
	"CNY": &ISO4217{alphabeticCode: "CNY", numericCode: "156", symbol: "¥", shortSymbol: "", decimalPlaces: 2, name: "Yuan Renminbi"},
	"COP": &ISO4217{alphabeticCode: "COP", numericCode: "170", symbol: "Col$", shortSymbol: "$", decimalPlaces: 2, name: "Colombian Peso"},
	"COU": &ISO4217{alphabeticCode: "COU", numericCode: "970", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Unidad de Valor Real"},
	"KMF": &ISO4217{alphabeticCode: "KMF", numericCode: "174", symbol: "CF", shortSymbol: "Fr", decimalPlaces: 0, name: "Comorian Franc "},
	"CDF": &ISO4217{alphabeticCode: "CDF", numericCode: "976", symbol: "F", shortSymbol: "", decimalPlaces: 2, name: "Congolese Franc"},
	"NZD": &ISO4217{alphabeticCode: "NZD", numericCode: "554", symbol: "NZ$", shortSymbol: "$", decimalPlaces: 2, name: "New Zealand Dollar"},
	"CRC": &ISO4217{alphabeticCode: "CRC", numericCode: "188", symbol: "₡", shortSymbol: "", decimalPlaces: 2, name: "Costa Rican Colon"},
	"HRK": &ISO4217{alphabeticCode: "HRK", numericCode: "191", symbol: "kn", shortSymbol: "", decimalPlaces: 2, name: "Kuna"},
	"CUP": &ISO4217{alphabeticCode: "CUP", numericCode: "192", symbol: "₱", shortSymbol: "", decimalPlaces: 2, name: "Cuban Peso"},
	"CUC": &ISO4217{alphabeticCode: "CUC", numericCode: "931", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Peso Convertible"},
	"ANG": &ISO4217{alphabeticCode: "ANG", numericCode: "532", symbol: "NAƒ", shortSymbol: "", decimalPlaces: 2, name: "Netherlands Antillean Guilder"},
	"CZK": &ISO4217{alphabeticCode: "CZK", numericCode: "203", symbol: "Kč", shortSymbol: "", decimalPlaces: 2, name: "Czech Koruna"},
	"DKK": &ISO4217{alphabeticCode: "DKK", numericCode: "208", symbol: "Kr", shortSymbol: "", decimalPlaces: 2, name: "Danish Krone"},
	"DJF": &ISO4217{alphabeticCode: "DJF", numericCode: "262", symbol: "Fdj", shortSymbol: "", decimalPlaces: 0, name: "Djibouti Franc"},
	"DOP": &ISO4217{alphabeticCode: "DOP", numericCode: "214", symbol: "RD$", shortSymbol: "$", decimalPlaces: 2, name: "Dominican Peso"},
	"EGP": &ISO4217{alphabeticCode: "EGP", numericCode: "818", symbol: "E£", shortSymbol: "£", decimalPlaces: 2, name: "Egyptian Pound"},
	"SVC": &ISO4217{alphabeticCode: "SVC", numericCode: "222", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "El Salvador Colon"},
	"ERN": &ISO4217{alphabeticCode: "ERN", numericCode: "232", symbol: "Nkf", shortSymbol: "", decimalPlaces: 2, name: "Nakfa"},
	"SZL": &ISO4217{alphabeticCode: "SZL", numericCode: "748", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lilangeni"},
	"ETB": &ISO4217{alphabeticCode: "ETB", numericCode: "230", symbol: "Br", shortSymbol: "", decimalPlaces: 2, name: "Ethiopian Birr"},
	"FKP": &ISO4217{alphabeticCode: "FKP", numericCode: "238", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Falkland Islands Pound"},
	"FJD": &ISO4217{alphabeticCode: "FJD", numericCode: "242", symbol: "FJ$", shortSymbol: "$", decimalPlaces: 2, name: "Fiji Dollar"},
	"XPF": &ISO4217{alphabeticCode: "XPF", numericCode: "953", symbol: "₣", shortSymbol: "", decimalPlaces: 0, name: "CFP Franc"},
	"GMD": &ISO4217{alphabeticCode: "GMD", numericCode: "270", symbol: "D", shortSymbol: "", decimalPlaces: 2, name: "Dalasi"},
	"GEL": &ISO4217{alphabeticCode: "GEL", numericCode: "981", symbol: "₾", shortSymbol: "", decimalPlaces: 2, name: "Lari"},
	"GHS": &ISO4217{alphabeticCode: "GHS", numericCode: "936", symbol: "₵", shortSymbol: "", decimalPlaces: 2, name: "Ghana Cedi"},
	"GIP": &ISO4217{alphabeticCode: "GIP", numericCode: "292", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Gibraltar Pound"},
	"GTQ": &ISO4217{alphabeticCode: "GTQ", numericCode: "320", symbol: "Q", shortSymbol: "", decimalPlaces: 2, name: "Quetzal"},
	"GBP": &ISO4217{alphabeticCode: "GBP", numericCode: "826", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Pound Sterling"},
	"GNF": &ISO4217{alphabeticCode: "GNF", numericCode: "324", symbol: "FG", shortSymbol: "", decimalPlaces: 0, name: "Guinean Franc"},
	"GYD": &ISO4217{alphabeticCode: "GYD", numericCode: "328", symbol: "GY$", shortSymbol: "", decimalPlaces: 2, name: "Guyana Dollar"},
	"HTG": &ISO4217{alphabeticCode: "HTG", numericCode: "332", symbol: "G", shortSymbol: "", decimalPlaces: 2, name: "Gourde"},
	"HNL": &ISO4217{alphabeticCode: "HNL", numericCode: "340", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lempira"},
	"HKD": &ISO4217{alphabeticCode: "HKD", numericCode: "344", symbol: "HK$", shortSymbol: "", decimalPlaces: 2, name: "Hong Kong Dollar"},
	"HUF": &ISO4217{alphabeticCode: "HUF", numericCode: "348", symbol: "Ft", shortSymbol: "", decimalPlaces: 2, name: "Forint"},
	"ISK": &ISO4217{alphabeticCode: "ISK", numericCode: "352", symbol: "kr", shortSymbol: "", decimalPlaces: 0, name: "Iceland Krona"},
	"IDR": &ISO4217{alphabeticCode: "IDR", numericCode: "360", symbol: "Rp", shortSymbol: "", decimalPlaces: 2, name: "Rupiah"},
	"XDR": &ISO4217{alphabeticCode: "XDR", numericCode: "960", symbol: "SDR", shortSymbol: "", decimalPlaces: -1, name: "SDR (Special Drawing Right)"},
	"IRR": &ISO4217{alphabeticCode: "IRR", numericCode: "364", symbol: "﷼", shortSymbol: "", decimalPlaces: 2, name: "Iranian Rial"},
	"IQD": &ISO4217{alphabeticCode: "IQD", numericCode: "368", symbol: "د.ع", shortSymbol: "", decimalPlaces: 3, name: "Iraqi Dinar"},
	"ILS": &ISO4217{alphabeticCode: "ILS", numericCode: "376", symbol: "₪", shortSymbol: "", decimalPlaces: 2, name: "New Israeli Sheqel"},
	"JMD": &ISO4217{alphabeticCode: "JMD", numericCode: "388", symbol: "J$", shortSymbol: "", decimalPlaces: 2, name: "Jamaican Dollar"},
	"JPY": &ISO4217{alphabeticCode: "JPY", numericCode: "392", symbol: "¥", shortSymbol: "", decimalPlaces: 0, name: "Yen"},
	"JOD": &ISO4217{alphabeticCode: "JOD", numericCode: "400", symbol: "د.أ", shortSymbol: "", decimalPlaces: 3, name: "Jordanian Dinar"},
	"KZT": &ISO4217{alphabeticCode: "KZT", numericCode: "398", symbol: "₸", shortSymbol: "", decimalPlaces: 2, name: "Tenge"},
	"KES": &ISO4217{alphabeticCode: "KES", numericCode: "404", symbol: "KSh", shortSymbol: "Sh", decimalPlaces: 2, name: "Kenyan Shilling"},
	"KPW": &ISO4217{alphabeticCode: "KPW", numericCode: "408", symbol: "₩", shortSymbol: "", decimalPlaces: 2, name: "North Korean Won"},
	"KRW": &ISO4217{alphabeticCode: "KRW", numericCode: "410", symbol: "₩", shortSymbol: "", decimalPlaces: 0, name: "Won"},
	"KWD": &ISO4217{alphabeticCode: "KWD", numericCode: "414", symbol: "KD", shortSymbol: "", decimalPlaces: 3, name: "Kuwaiti Dinar"},
	"KGS": &ISO4217{alphabeticCode: "KGS", numericCode: "417", symbol: "⃀", shortSymbol: "", decimalPlaces: 2, name: "Som"},
	"LAK": &ISO4217{alphabeticCode: "LAK", numericCode: "418", symbol: "₭", shortSymbol: "", decimalPlaces: 2, name: "Lao Kip"},
	"LBP": &ISO4217{alphabeticCode: "LBP", numericCode: "422", symbol: "ل.ل", shortSymbol: "", decimalPlaces: 2, name: "Lebanese Pound"},
	"LSL": &ISO4217{alphabeticCode: "LSL", numericCode: "426", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Loti"},
	"ZAR": &ISO4217{alphabeticCode: "ZAR", numericCode: "710", symbol: "R", shortSymbol: "", decimalPlaces: 2, name: "Rand"},
	"LRD": &ISO4217{alphabeticCode: "LRD", numericCode: "430", symbol: "LD$", shortSymbol: "$", decimalPlaces: 2, name: "Liberian Dollar"},
	"LYD": &ISO4217{alphabeticCode: "LYD", numericCode: "434", symbol: "ل.د", shortSymbol: "", decimalPlaces: 3, name: "Libyan Dinar"},
	"CHF": &ISO4217{alphabeticCode: "CHF", numericCode: "756", symbol: "Fr.", shortSymbol: "", decimalPlaces: 2, name: "Swiss Franc"},
	"MOP": &ISO4217{alphabeticCode: "MOP", numericCode: "446", symbol: "MOP$", shortSymbol: "", decimalPlaces: 2, name: "Pataca"},
	"MKD": &ISO4217{alphabeticCode: "MKD", numericCode: "807", symbol: "ден", shortSymbol: "", decimalPlaces: 2, name: "Denar"},
	"MGA": &ISO4217{alphabeticCode: "MGA", numericCode: "969", symbol: "Ar", shortSymbol: "", decimalPlaces: 2, name: "Malagasy Ariary"},
	"MWK": &ISO4217{alphabeticCode: "MWK", numericCode: "454", symbol: "MK", shortSymbol: "", decimalPlaces: 2, name: "Malawi Kwacha"},
	"MYR": &ISO4217{alphabeticCode: "MYR", numericCode: "458", symbol: "RM", shortSymbol: "", decimalPlaces: 2, name: "Malaysian Ringgit"},
	"MVR": &ISO4217{alphabeticCode: "MVR", numericCode: "462", symbol: "Rf", shortSymbol: "", decimalPlaces: 2, name: "Rufiyaa"},
	"MRU": &ISO4217{alphabeticCode: "MRU", numericCode: "929", symbol: "UM", shortSymbol: "", decimalPlaces: 2, name: "Ouguiya"},
	"MUR": &ISO4217{alphabeticCode: "MUR", numericCode: "480", symbol: "Rs", shortSymbol: "", decimalPlaces: 2, name: "Mauritius Rupee"},
	"XUA": &ISO4217{alphabeticCode: "XUA", numericCode: "965", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "ADB Unit of Account"},
	"MXN": &ISO4217{alphabeticCode: "MXN", numericCode: "484", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Mexican Peso"},
	"MXV": &ISO4217{alphabeticCode: "MXV", numericCode: "979", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Mexican Unidad de Inversion (UDI)"},
	"MDL": &ISO4217{alphabeticCode: "MDL", numericCode: "498", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Moldovan Leu"},
	"MNT": &ISO4217{alphabeticCode: "MNT", numericCode: "496", symbol: "₮", shortSymbol: "", decimalPlaces: 2, name: "Tugrik"},
	"MAD": &ISO4217{alphabeticCode: "MAD", numericCode: "504", symbol: "DH", shortSymbol: "", decimalPlaces: 2, name: "Moroccan Dirham"},
	"MZN": &ISO4217{alphabeticCode: "MZN", numericCode: "943", symbol: "MT", shortSymbol: "", decimalPlaces: 2, name: "Mozambique Metical"},
	"MMK": &ISO4217{alphabeticCode: "MMK", numericCode: "104", symbol: "K", shortSymbol: "", decimalPlaces: 2, name: "Kyat"},
	"NAD": &ISO4217{alphabeticCode: "NAD", numericCode: "516", symbol: "N$", shortSymbol: "", decimalPlaces: 2, name: "Namibia Dollar"},
	"NPR": &ISO4217{alphabeticCode: "NPR", numericCode: "524", symbol: "NRs", shortSymbol: "", decimalPlaces: 2, name: "Nepalese Rupee"},
	"NIO": &ISO4217{alphabeticCode: "NIO", numericCode: "558", symbol: "C$", shortSymbol: "", decimalPlaces: 2, name: "Cordoba Oro"},
	"NGN": &ISO4217{alphabeticCode: "NGN", numericCode: "566", symbol: "₦", shortSymbol: "", decimalPlaces: 2, name: "Naira"},
	"OMR": &ISO4217{alphabeticCode: "OMR", numericCode: "512", symbol: "ر.ع.", shortSymbol: "", decimalPlaces: 3, name: "Rial Omani"},
	"PKR": &ISO4217{alphabeticCode: "PKR", numericCode: "586", symbol: "Rs.", shortSymbol: "", decimalPlaces: 2, name: "Pakistan Rupee"},
	"PAB": &ISO4217{alphabeticCode: "PAB", numericCode: "590", symbol: "B./", shortSymbol: "", decimalPlaces: 2, name: "Balboa"},
	"PGK": &ISO4217{alphabeticCode: "PGK", numericCode: "598", symbol: "K", shortSymbol: "", decimalPlaces: 2, name: "Kina"},
	"PYG": &ISO4217{alphabeticCode: "PYG", numericCode: "600", symbol: "₲", shortSymbol: "", decimalPlaces: 0, name: "Guarani"},
	"PEN": &ISO4217{alphabeticCode: "PEN", numericCode: "604", symbol: "S/.", shortSymbol: "", decimalPlaces: 2, name: "Sol"},
	"PHP": &ISO4217{alphabeticCode: "PHP", numericCode: "608", symbol: "₱", shortSymbol: "", decimalPlaces: 2, name: "Philippine Peso"},
	"PLN": &ISO4217{alphabeticCode: "PLN", numericCode: "985", symbol: "zł", shortSymbol: "", decimalPlaces: 2, name: "Zloty"},
	"QAR": &ISO4217{alphabeticCode: "QAR", numericCode: "634", symbol: "QR", shortSymbol: "", decimalPlaces: 2, name: "Qatari Rial"},
	"RON": &ISO4217{alphabeticCode: "RON", numericCode: "946", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Romanian Leu"},
	"RUB": &ISO4217{alphabeticCode: "RUB", numericCode: "643", symbol: "R", shortSymbol: "", decimalPlaces: 2, name: "Russian Ruble"},
	"RWF": &ISO4217{alphabeticCode: "RWF", numericCode: "646", symbol: "RF", shortSymbol: "", decimalPlaces: 0, name: "Rwanda Franc"},
	"SHP": &ISO4217{alphabeticCode: "SHP", numericCode: "654", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Saint Helena Pound"},
	"WST": &ISO4217{alphabeticCode: "WST", numericCode: "882", symbol: "WS$", shortSymbol: "", decimalPlaces: 2, name: "Tala"},
	"STN": &ISO4217{alphabeticCode: "STN", numericCode: "930", symbol: "Db", shortSymbol: "", decimalPlaces: 2, name: "Dobra"},
	"SAR": &ISO4217{alphabeticCode: "SAR", numericCode: "682", symbol: "SR", shortSymbol: "", decimalPlaces: 2, name: "Saudi Riyal"},
	"RSD": &ISO4217{alphabeticCode: "RSD", numericCode: "941", symbol: "din.", shortSymbol: "", decimalPlaces: 2, name: "Serbian Dinar"},
	"SCR": &ISO4217{alphabeticCode: "SCR", numericCode: "690", symbol: "SR", shortSymbol: "", decimalPlaces: 2, name: "Seychelles Rupee"},
	"SLL": &ISO4217{alphabeticCode: "SLL", numericCode: "694", symbol: "Le", shortSymbol: "", decimalPlaces: 2, name: "Leone"},
	"SGD": &ISO4217{alphabeticCode: "SGD", numericCode: "702", symbol: "S$", shortSymbol: "", decimalPlaces: 2, name: "Singapore Dollar"},
	"XSU": &ISO4217{alphabeticCode: "XSU", numericCode: "994", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Sucre"},
	"SBD": &ISO4217{alphabeticCode: "SBD", numericCode: "090", symbol: "SI$", shortSymbol: "", decimalPlaces: 2, name: "Solomon Islands Dollar"},
	"SOS": &ISO4217{alphabeticCode: "SOS", numericCode: "706", symbol: "Sh.", shortSymbol: "", decimalPlaces: 2, name: "Somali Shilling"},
	"SSP": &ISO4217{alphabeticCode: "SSP", numericCode: "728", symbol: "SS£", shortSymbol: "", decimalPlaces: 2, name: "South Sudanese Pound"},
	"LKR": &ISO4217{alphabeticCode: "LKR", numericCode: "144", symbol: "Rs", shortSymbol: "", decimalPlaces: 2, name: "Sri Lanka Rupee"},
	"SDG": &ISO4217{alphabeticCode: "SDG", numericCode: "938", symbol: "£SD", shortSymbol: "", decimalPlaces: 2, name: "Sudanese Pound"},
	"SRD": &ISO4217{alphabeticCode: "SRD", numericCode: "968", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Surinam Dollar"},
	"SEK": &ISO4217{alphabeticCode: "SEK", numericCode: "752", symbol: "kr", shortSymbol: "", decimalPlaces: 2, name: "Swedish Krona"},
	"CHE": &ISO4217{alphabeticCode: "CHE", numericCode: "947", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "WIR Euro"},
	"CHW": &ISO4217{alphabeticCode: "CHW", numericCode: "948", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "WIR Franc"},
	"SYP": &ISO4217{alphabeticCode: "SYP", numericCode: "760", symbol: "£S", shortSymbol: "", decimalPlaces: 2, name: "Syrian Pound"},
	"TWD": &ISO4217{alphabeticCode: "TWD", numericCode: "901", symbol: "NT$", shortSymbol: "$", decimalPlaces: 2, name: "New Taiwan Dollar"},
	"TJS": &ISO4217{alphabeticCode: "TJS", numericCode: "972", symbol: "SM", shortSymbol: "", decimalPlaces: 2, name: "Somoni"},
	"TZS": &ISO4217{alphabeticCode: "TZS", numericCode: "834", symbol: "TSh", shortSymbol: "", decimalPlaces: 2, name: "Tanzanian Shilling"},
	"THB": &ISO4217{alphabeticCode: "THB", numericCode: "764", symbol: "฿", shortSymbol: "", decimalPlaces: 2, name: "Baht"},
	"TOP": &ISO4217{alphabeticCode: "TOP", numericCode: "776", symbol: "T$", shortSymbol: "", decimalPlaces: 2, name: "Pa’anga"},
	"TTD": &ISO4217{alphabeticCode: "TTD", numericCode: "780", symbol: "TT$", shortSymbol: "", decimalPlaces: 2, name: "Trinidad and Tobago Dollar"},
	"TND": &ISO4217{alphabeticCode: "TND", numericCode: "788", symbol: "DT", shortSymbol: "", decimalPlaces: 3, name: "Tunisian Dinar"},
	"TRY": &ISO4217{alphabeticCode: "TRY", numericCode: "949", symbol: "YTL", shortSymbol: "", decimalPlaces: 2, name: "Turkish Lira"},
	"TMT": &ISO4217{alphabeticCode: "TMT", numericCode: "934", symbol: "m", shortSymbol: "", decimalPlaces: 2, name: "Turkmenistan New Manat"},
	"UGX": &ISO4217{alphabeticCode: "UGX", numericCode: "800", symbol: "USh", shortSymbol: "", decimalPlaces: 0, name: "Uganda Shilling"},
	"UAH": &ISO4217{alphabeticCode: "UAH", numericCode: "980", symbol: "₴", shortSymbol: "", decimalPlaces: 2, name: "Hryvnia"},
	"AED": &ISO4217{alphabeticCode: "AED", numericCode: "784", symbol: "د.إ", shortSymbol: "", decimalPlaces: 2, name: "UAE Dirham"},
	"USN": &ISO4217{alphabeticCode: "USN", numericCode: "997", symbol: "US$", shortSymbol: "$", decimalPlaces: 2, name: "US Dollar (Next day)"},
	"UYU": &ISO4217{alphabeticCode: "UYU", numericCode: "858", symbol: "$U", shortSymbol: "$", decimalPlaces: 2, name: "Peso Uruguayo"},
	"UYI": &ISO4217{alphabeticCode: "UYI", numericCode: "940", symbol: "", shortSymbol: "", decimalPlaces: 0, name: "Uruguay Peso en Unidades Indexadas (UI)"},
	"UYW": &ISO4217{alphabeticCode: "UYW", numericCode: "927", symbol: "", shortSymbol: "", decimalPlaces: 4, name: "Unidad Previsional"},
	"UZS": &ISO4217{alphabeticCode: "UZS", numericCode: "860", symbol: "сум", shortSymbol: "", decimalPlaces: 2, name: "Uzbekistan Sum"},
	"VUV": &ISO4217{alphabeticCode: "VUV", numericCode: "548", symbol: "VT", shortSymbol: "", decimalPlaces: 0, name: "Vatu"},
	"VES": &ISO4217{alphabeticCode: "VES", numericCode: "928", symbol: "Bs.S", shortSymbol: "", decimalPlaces: 2, name: "Bolívar Soberano"},
	"VND": &ISO4217{alphabeticCode: "VND", numericCode: "704", symbol: "₫", shortSymbol: "", decimalPlaces: 0, name: "Dong"},
	"YER": &ISO4217{alphabeticCode: "YER", numericCode: "886", symbol: "﷼", shortSymbol: "", decimalPlaces: 2, name: "Yemeni Rial"},
	"ZMW": &ISO4217{alphabeticCode: "ZMW", numericCode: "967", symbol: "ZK", shortSymbol: "", decimalPlaces: 2, name: "Zambian Kwacha"},
	"ZWL": &ISO4217{alphabeticCode: "ZWL", numericCode: "932", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Zimbabwe Dollar"},
	"XBA": &ISO4217{alphabeticCode: "XBA", numericCode: "955", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Composite Unit (EURCO)"},
	"XBB": &ISO4217{alphabeticCode: "XBB", numericCode: "956", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)"},
	"XBC": &ISO4217{alphabeticCode: "XBC", numericCode: "957", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)"},
	"XBD": &ISO4217{alphabeticCode: "XBD", numericCode: "958", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)"},
	"XTS": &ISO4217{alphabeticCode: "XTS", numericCode: "963", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Codes specifically reserved for testing purposes"},
	"XXX": &ISO4217{alphabeticCode: "XXX", numericCode: "999", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "The codes assigned for transactions where no currency is involved"},
	"XAU": &ISO4217{alphabeticCode: "XAU", numericCode: "959", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Gold"},
	"XPD": &ISO4217{alphabeticCode: "XPD", numericCode: "964", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Palladium"},
	"XPT": &ISO4217{alphabeticCode: "XPT", numericCode: "962", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Platinum"},
	"XAG": &ISO4217{alphabeticCode: "XAG", numericCode: "961", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Silver"},
}
