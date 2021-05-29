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

// Code returns a alphabetic code representing the currency.
// This may be a ISO 4217 code, depending on the currency type.
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

// Standard returns a string that identifies the standard that the currency is defined in.
func (c *ISO4217) Standard() string {
	return "ISO\u00A04217"
}

func (c *ISO4217) String() string {
	return c.Standard() + " " + c.alphabeticCode
}

// ISO4217Currencies contains the official and active ISO 4217 currencies as of August 29, 2018.
//
// Source: https://www.currency-iso.org/en/home/tables/table-a1.html
//
// This data has been modified in the following ways:
//  - Removed "ENTITY" column
//  - Removed duplicate entries (due to removal of the "ENTITY" column)
//  - Add "symbol" and "shortSymbol" columns that contain symbols which are NOT part of ISO 4217. Based on https://web.archive.org/web/20111129141202/http://fx.sauder.ubc.ca/currency_table.html and https://wikipedia.org
// TODO: Add more currency symbols, check symbols for correctness, add symbol plural variants
var ISO4217Currencies = map[string]*ISO4217{
	"AFN": {alphabeticCode: "AFN", numericCode: "971", symbol: "؋", shortSymbol: "", decimalPlaces: 2, name: "Afghani"},
	"EUR": {alphabeticCode: "EUR", numericCode: "978", symbol: "€", shortSymbol: "", decimalPlaces: 2, name: "Euro"},
	"ALL": {alphabeticCode: "ALL", numericCode: "008", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lek"},
	"DZD": {alphabeticCode: "DZD", numericCode: "012", symbol: "DA", shortSymbol: "", decimalPlaces: 2, name: "Algerian Dinar"},
	"USD": {alphabeticCode: "USD", numericCode: "840", symbol: "US$", shortSymbol: "$", decimalPlaces: 2, name: "US Dollar"},
	"AOA": {alphabeticCode: "AOA", numericCode: "973", symbol: "Kz", shortSymbol: "", decimalPlaces: 2, name: "Kwanza"},
	"XCD": {alphabeticCode: "XCD", numericCode: "951", symbol: "EC$", shortSymbol: "$", decimalPlaces: 2, name: "East Caribbean Dollar"},
	"ARS": {alphabeticCode: "ARS", numericCode: "032", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Argentine Peso"},
	"AMD": {alphabeticCode: "AMD", numericCode: "051", symbol: "֏", shortSymbol: "", decimalPlaces: 2, name: "Armenian Dram"},
	"AWG": {alphabeticCode: "AWG", numericCode: "533", symbol: "ƒ", shortSymbol: "", decimalPlaces: 2, name: "Aruban Florin"},
	"AUD": {alphabeticCode: "AUD", numericCode: "036", symbol: "AU$", shortSymbol: "$", decimalPlaces: 2, name: "Australian Dollar"},
	"AZN": {alphabeticCode: "AZN", numericCode: "944", symbol: "₼", shortSymbol: "", decimalPlaces: 2, name: "Azerbaijan Manat"},
	"BSD": {alphabeticCode: "BSD", numericCode: "044", symbol: "B$", shortSymbol: "$", decimalPlaces: 2, name: "Bahamian Dollar"},
	"BHD": {alphabeticCode: "BHD", numericCode: "048", symbol: "BD", shortSymbol: "", decimalPlaces: 3, name: "Bahraini Dinar"},
	"BDT": {alphabeticCode: "BDT", numericCode: "050", symbol: "৳", shortSymbol: "", decimalPlaces: 2, name: "Taka"},
	"BBD": {alphabeticCode: "BBD", numericCode: "052", symbol: "Bds$", shortSymbol: "$", decimalPlaces: 2, name: "Barbados Dollar"},
	"BYN": {alphabeticCode: "BYN", numericCode: "933", symbol: "Br", shortSymbol: "", decimalPlaces: 2, name: "Belarusian Ruble"},
	"BZD": {alphabeticCode: "BZD", numericCode: "084", symbol: "BZ$", shortSymbol: "$", decimalPlaces: 2, name: "Belize Dollar"},
	"XOF": {alphabeticCode: "XOF", numericCode: "952", symbol: "CFA", shortSymbol: "Fr", decimalPlaces: 0, name: "CFA Franc BCEAO"},
	"BMD": {alphabeticCode: "BMD", numericCode: "060", symbol: "BD$", shortSymbol: "$", decimalPlaces: 2, name: "Bermudian Dollar"},
	"INR": {alphabeticCode: "INR", numericCode: "356", symbol: "₹", shortSymbol: "", decimalPlaces: 2, name: "Indian Rupee"},
	"BTN": {alphabeticCode: "BTN", numericCode: "064", symbol: "Nu.", shortSymbol: "", decimalPlaces: 2, name: "Ngultrum"},
	"BOB": {alphabeticCode: "BOB", numericCode: "068", symbol: "Bs.", shortSymbol: "", decimalPlaces: 2, name: "Boliviano"},
	"BOV": {alphabeticCode: "BOV", numericCode: "984", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Mvdol"},
	"BAM": {alphabeticCode: "BAM", numericCode: "977", symbol: "KM", shortSymbol: "", decimalPlaces: 2, name: "Convertible Mark"},
	"BWP": {alphabeticCode: "BWP", numericCode: "072", symbol: "P", shortSymbol: "", decimalPlaces: 2, name: "Pula"},
	"NOK": {alphabeticCode: "NOK", numericCode: "578", symbol: "kr", shortSymbol: "", decimalPlaces: 2, name: "Norwegian Krone"},
	"BRL": {alphabeticCode: "BRL", numericCode: "986", symbol: "R$", shortSymbol: "", decimalPlaces: 2, name: "Brazilian Real"},
	"BND": {alphabeticCode: "BND", numericCode: "096", symbol: "B$", shortSymbol: "$", decimalPlaces: 2, name: "Brunei Dollar"},
	"BGN": {alphabeticCode: "BGN", numericCode: "975", symbol: "лв.", shortSymbol: "", decimalPlaces: 2, name: "Bulgarian Lev"},
	"BIF": {alphabeticCode: "BIF", numericCode: "108", symbol: "FBu", shortSymbol: "", decimalPlaces: 0, name: "Burundi Franc"},
	"CVE": {alphabeticCode: "CVE", numericCode: "132", symbol: "Esc", shortSymbol: "", decimalPlaces: 2, name: "Cabo Verde Escudo"},
	"KHR": {alphabeticCode: "KHR", numericCode: "116", symbol: "៛", shortSymbol: "", decimalPlaces: 2, name: "Riel"},
	"XAF": {alphabeticCode: "XAF", numericCode: "950", symbol: "CFA", shortSymbol: "Fr", decimalPlaces: 0, name: "CFA Franc BEAC"},
	"CAD": {alphabeticCode: "CAD", numericCode: "124", symbol: "CA$", shortSymbol: "$", decimalPlaces: 2, name: "Canadian Dollar"},
	"KYD": {alphabeticCode: "KYD", numericCode: "136", symbol: "KY$", shortSymbol: "", decimalPlaces: 2, name: "Cayman Islands Dollar"},
	"CLP": {alphabeticCode: "CLP", numericCode: "152", symbol: "CLP$", shortSymbol: "$", decimalPlaces: 0, name: "Chilean Peso"},
	"CLF": {alphabeticCode: "CLF", numericCode: "990", symbol: "", shortSymbol: "", decimalPlaces: 4, name: "Unidad de Fomento"},
	"CNY": {alphabeticCode: "CNY", numericCode: "156", symbol: "¥", shortSymbol: "", decimalPlaces: 2, name: "Yuan Renminbi"},
	"COP": {alphabeticCode: "COP", numericCode: "170", symbol: "Col$", shortSymbol: "$", decimalPlaces: 2, name: "Colombian Peso"},
	"COU": {alphabeticCode: "COU", numericCode: "970", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Unidad de Valor Real"},
	"KMF": {alphabeticCode: "KMF", numericCode: "174", symbol: "CF", shortSymbol: "Fr", decimalPlaces: 0, name: "Comorian Franc "},
	"CDF": {alphabeticCode: "CDF", numericCode: "976", symbol: "F", shortSymbol: "", decimalPlaces: 2, name: "Congolese Franc"},
	"NZD": {alphabeticCode: "NZD", numericCode: "554", symbol: "NZ$", shortSymbol: "$", decimalPlaces: 2, name: "New Zealand Dollar"},
	"CRC": {alphabeticCode: "CRC", numericCode: "188", symbol: "₡", shortSymbol: "", decimalPlaces: 2, name: "Costa Rican Colon"},
	"HRK": {alphabeticCode: "HRK", numericCode: "191", symbol: "kn", shortSymbol: "", decimalPlaces: 2, name: "Kuna"},
	"CUP": {alphabeticCode: "CUP", numericCode: "192", symbol: "₱", shortSymbol: "", decimalPlaces: 2, name: "Cuban Peso"},
	"CUC": {alphabeticCode: "CUC", numericCode: "931", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Peso Convertible"},
	"ANG": {alphabeticCode: "ANG", numericCode: "532", symbol: "NAƒ", shortSymbol: "", decimalPlaces: 2, name: "Netherlands Antillean Guilder"},
	"CZK": {alphabeticCode: "CZK", numericCode: "203", symbol: "Kč", shortSymbol: "", decimalPlaces: 2, name: "Czech Koruna"},
	"DKK": {alphabeticCode: "DKK", numericCode: "208", symbol: "Kr", shortSymbol: "", decimalPlaces: 2, name: "Danish Krone"},
	"DJF": {alphabeticCode: "DJF", numericCode: "262", symbol: "Fdj", shortSymbol: "", decimalPlaces: 0, name: "Djibouti Franc"},
	"DOP": {alphabeticCode: "DOP", numericCode: "214", symbol: "RD$", shortSymbol: "$", decimalPlaces: 2, name: "Dominican Peso"},
	"EGP": {alphabeticCode: "EGP", numericCode: "818", symbol: "E£", shortSymbol: "£", decimalPlaces: 2, name: "Egyptian Pound"},
	"SVC": {alphabeticCode: "SVC", numericCode: "222", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "El Salvador Colon"},
	"ERN": {alphabeticCode: "ERN", numericCode: "232", symbol: "Nkf", shortSymbol: "", decimalPlaces: 2, name: "Nakfa"},
	"SZL": {alphabeticCode: "SZL", numericCode: "748", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lilangeni"},
	"ETB": {alphabeticCode: "ETB", numericCode: "230", symbol: "Br", shortSymbol: "", decimalPlaces: 2, name: "Ethiopian Birr"},
	"FKP": {alphabeticCode: "FKP", numericCode: "238", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Falkland Islands Pound"},
	"FJD": {alphabeticCode: "FJD", numericCode: "242", symbol: "FJ$", shortSymbol: "$", decimalPlaces: 2, name: "Fiji Dollar"},
	"XPF": {alphabeticCode: "XPF", numericCode: "953", symbol: "₣", shortSymbol: "", decimalPlaces: 0, name: "CFP Franc"},
	"GMD": {alphabeticCode: "GMD", numericCode: "270", symbol: "D", shortSymbol: "", decimalPlaces: 2, name: "Dalasi"},
	"GEL": {alphabeticCode: "GEL", numericCode: "981", symbol: "₾", shortSymbol: "", decimalPlaces: 2, name: "Lari"},
	"GHS": {alphabeticCode: "GHS", numericCode: "936", symbol: "₵", shortSymbol: "", decimalPlaces: 2, name: "Ghana Cedi"},
	"GIP": {alphabeticCode: "GIP", numericCode: "292", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Gibraltar Pound"},
	"GTQ": {alphabeticCode: "GTQ", numericCode: "320", symbol: "Q", shortSymbol: "", decimalPlaces: 2, name: "Quetzal"},
	"GBP": {alphabeticCode: "GBP", numericCode: "826", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Pound Sterling"},
	"GNF": {alphabeticCode: "GNF", numericCode: "324", symbol: "FG", shortSymbol: "", decimalPlaces: 0, name: "Guinean Franc"},
	"GYD": {alphabeticCode: "GYD", numericCode: "328", symbol: "GY$", shortSymbol: "", decimalPlaces: 2, name: "Guyana Dollar"},
	"HTG": {alphabeticCode: "HTG", numericCode: "332", symbol: "G", shortSymbol: "", decimalPlaces: 2, name: "Gourde"},
	"HNL": {alphabeticCode: "HNL", numericCode: "340", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Lempira"},
	"HKD": {alphabeticCode: "HKD", numericCode: "344", symbol: "HK$", shortSymbol: "", decimalPlaces: 2, name: "Hong Kong Dollar"},
	"HUF": {alphabeticCode: "HUF", numericCode: "348", symbol: "Ft", shortSymbol: "", decimalPlaces: 2, name: "Forint"},
	"ISK": {alphabeticCode: "ISK", numericCode: "352", symbol: "kr", shortSymbol: "", decimalPlaces: 0, name: "Iceland Krona"},
	"IDR": {alphabeticCode: "IDR", numericCode: "360", symbol: "Rp", shortSymbol: "", decimalPlaces: 2, name: "Rupiah"},
	"XDR": {alphabeticCode: "XDR", numericCode: "960", symbol: "SDR", shortSymbol: "", decimalPlaces: -1, name: "SDR (Special Drawing Right)"},
	"IRR": {alphabeticCode: "IRR", numericCode: "364", symbol: "﷼", shortSymbol: "", decimalPlaces: 2, name: "Iranian Rial"},
	"IQD": {alphabeticCode: "IQD", numericCode: "368", symbol: "د.ع", shortSymbol: "", decimalPlaces: 3, name: "Iraqi Dinar"},
	"ILS": {alphabeticCode: "ILS", numericCode: "376", symbol: "₪", shortSymbol: "", decimalPlaces: 2, name: "New Israeli Sheqel"},
	"JMD": {alphabeticCode: "JMD", numericCode: "388", symbol: "J$", shortSymbol: "", decimalPlaces: 2, name: "Jamaican Dollar"},
	"JPY": {alphabeticCode: "JPY", numericCode: "392", symbol: "¥", shortSymbol: "", decimalPlaces: 0, name: "Yen"},
	"JOD": {alphabeticCode: "JOD", numericCode: "400", symbol: "د.أ", shortSymbol: "", decimalPlaces: 3, name: "Jordanian Dinar"},
	"KZT": {alphabeticCode: "KZT", numericCode: "398", symbol: "₸", shortSymbol: "", decimalPlaces: 2, name: "Tenge"},
	"KES": {alphabeticCode: "KES", numericCode: "404", symbol: "KSh", shortSymbol: "Sh", decimalPlaces: 2, name: "Kenyan Shilling"},
	"KPW": {alphabeticCode: "KPW", numericCode: "408", symbol: "₩", shortSymbol: "", decimalPlaces: 2, name: "North Korean Won"},
	"KRW": {alphabeticCode: "KRW", numericCode: "410", symbol: "₩", shortSymbol: "", decimalPlaces: 0, name: "Won"},
	"KWD": {alphabeticCode: "KWD", numericCode: "414", symbol: "KD", shortSymbol: "", decimalPlaces: 3, name: "Kuwaiti Dinar"},
	"KGS": {alphabeticCode: "KGS", numericCode: "417", symbol: "⃀", shortSymbol: "", decimalPlaces: 2, name: "Som"},
	"LAK": {alphabeticCode: "LAK", numericCode: "418", symbol: "₭", shortSymbol: "", decimalPlaces: 2, name: "Lao Kip"},
	"LBP": {alphabeticCode: "LBP", numericCode: "422", symbol: "ل.ل", shortSymbol: "", decimalPlaces: 2, name: "Lebanese Pound"},
	"LSL": {alphabeticCode: "LSL", numericCode: "426", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Loti"},
	"ZAR": {alphabeticCode: "ZAR", numericCode: "710", symbol: "R", shortSymbol: "", decimalPlaces: 2, name: "Rand"},
	"LRD": {alphabeticCode: "LRD", numericCode: "430", symbol: "LD$", shortSymbol: "$", decimalPlaces: 2, name: "Liberian Dollar"},
	"LYD": {alphabeticCode: "LYD", numericCode: "434", symbol: "ل.د", shortSymbol: "", decimalPlaces: 3, name: "Libyan Dinar"},
	"CHF": {alphabeticCode: "CHF", numericCode: "756", symbol: "Fr.", shortSymbol: "", decimalPlaces: 2, name: "Swiss Franc"},
	"MOP": {alphabeticCode: "MOP", numericCode: "446", symbol: "MOP$", shortSymbol: "", decimalPlaces: 2, name: "Pataca"},
	"MKD": {alphabeticCode: "MKD", numericCode: "807", symbol: "ден", shortSymbol: "", decimalPlaces: 2, name: "Denar"},
	"MGA": {alphabeticCode: "MGA", numericCode: "969", symbol: "Ar", shortSymbol: "", decimalPlaces: 2, name: "Malagasy Ariary"},
	"MWK": {alphabeticCode: "MWK", numericCode: "454", symbol: "MK", shortSymbol: "", decimalPlaces: 2, name: "Malawi Kwacha"},
	"MYR": {alphabeticCode: "MYR", numericCode: "458", symbol: "RM", shortSymbol: "", decimalPlaces: 2, name: "Malaysian Ringgit"},
	"MVR": {alphabeticCode: "MVR", numericCode: "462", symbol: "Rf", shortSymbol: "", decimalPlaces: 2, name: "Rufiyaa"},
	"MRU": {alphabeticCode: "MRU", numericCode: "929", symbol: "UM", shortSymbol: "", decimalPlaces: 2, name: "Ouguiya"},
	"MUR": {alphabeticCode: "MUR", numericCode: "480", symbol: "Rs", shortSymbol: "", decimalPlaces: 2, name: "Mauritius Rupee"},
	"XUA": {alphabeticCode: "XUA", numericCode: "965", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "ADB Unit of Account"},
	"MXN": {alphabeticCode: "MXN", numericCode: "484", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Mexican Peso"},
	"MXV": {alphabeticCode: "MXV", numericCode: "979", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Mexican Unidad de Inversion (UDI)"},
	"MDL": {alphabeticCode: "MDL", numericCode: "498", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Moldovan Leu"},
	"MNT": {alphabeticCode: "MNT", numericCode: "496", symbol: "₮", shortSymbol: "", decimalPlaces: 2, name: "Tugrik"},
	"MAD": {alphabeticCode: "MAD", numericCode: "504", symbol: "DH", shortSymbol: "", decimalPlaces: 2, name: "Moroccan Dirham"},
	"MZN": {alphabeticCode: "MZN", numericCode: "943", symbol: "MT", shortSymbol: "", decimalPlaces: 2, name: "Mozambique Metical"},
	"MMK": {alphabeticCode: "MMK", numericCode: "104", symbol: "K", shortSymbol: "", decimalPlaces: 2, name: "Kyat"},
	"NAD": {alphabeticCode: "NAD", numericCode: "516", symbol: "N$", shortSymbol: "", decimalPlaces: 2, name: "Namibia Dollar"},
	"NPR": {alphabeticCode: "NPR", numericCode: "524", symbol: "NRs", shortSymbol: "", decimalPlaces: 2, name: "Nepalese Rupee"},
	"NIO": {alphabeticCode: "NIO", numericCode: "558", symbol: "C$", shortSymbol: "", decimalPlaces: 2, name: "Cordoba Oro"},
	"NGN": {alphabeticCode: "NGN", numericCode: "566", symbol: "₦", shortSymbol: "", decimalPlaces: 2, name: "Naira"},
	"OMR": {alphabeticCode: "OMR", numericCode: "512", symbol: "ر.ع.", shortSymbol: "", decimalPlaces: 3, name: "Rial Omani"},
	"PKR": {alphabeticCode: "PKR", numericCode: "586", symbol: "Rs.", shortSymbol: "", decimalPlaces: 2, name: "Pakistan Rupee"},
	"PAB": {alphabeticCode: "PAB", numericCode: "590", symbol: "B./", shortSymbol: "", decimalPlaces: 2, name: "Balboa"},
	"PGK": {alphabeticCode: "PGK", numericCode: "598", symbol: "K", shortSymbol: "", decimalPlaces: 2, name: "Kina"},
	"PYG": {alphabeticCode: "PYG", numericCode: "600", symbol: "₲", shortSymbol: "", decimalPlaces: 0, name: "Guarani"},
	"PEN": {alphabeticCode: "PEN", numericCode: "604", symbol: "S/.", shortSymbol: "", decimalPlaces: 2, name: "Sol"},
	"PHP": {alphabeticCode: "PHP", numericCode: "608", symbol: "₱", shortSymbol: "", decimalPlaces: 2, name: "Philippine Peso"},
	"PLN": {alphabeticCode: "PLN", numericCode: "985", symbol: "zł", shortSymbol: "", decimalPlaces: 2, name: "Zloty"},
	"QAR": {alphabeticCode: "QAR", numericCode: "634", symbol: "QR", shortSymbol: "", decimalPlaces: 2, name: "Qatari Rial"},
	"RON": {alphabeticCode: "RON", numericCode: "946", symbol: "L", shortSymbol: "", decimalPlaces: 2, name: "Romanian Leu"},
	"RUB": {alphabeticCode: "RUB", numericCode: "643", symbol: "R", shortSymbol: "", decimalPlaces: 2, name: "Russian Ruble"},
	"RWF": {alphabeticCode: "RWF", numericCode: "646", symbol: "RF", shortSymbol: "", decimalPlaces: 0, name: "Rwanda Franc"},
	"SHP": {alphabeticCode: "SHP", numericCode: "654", symbol: "£", shortSymbol: "", decimalPlaces: 2, name: "Saint Helena Pound"},
	"WST": {alphabeticCode: "WST", numericCode: "882", symbol: "WS$", shortSymbol: "", decimalPlaces: 2, name: "Tala"},
	"STN": {alphabeticCode: "STN", numericCode: "930", symbol: "Db", shortSymbol: "", decimalPlaces: 2, name: "Dobra"},
	"SAR": {alphabeticCode: "SAR", numericCode: "682", symbol: "SR", shortSymbol: "", decimalPlaces: 2, name: "Saudi Riyal"},
	"RSD": {alphabeticCode: "RSD", numericCode: "941", symbol: "din.", shortSymbol: "", decimalPlaces: 2, name: "Serbian Dinar"},
	"SCR": {alphabeticCode: "SCR", numericCode: "690", symbol: "SR", shortSymbol: "", decimalPlaces: 2, name: "Seychelles Rupee"},
	"SLL": {alphabeticCode: "SLL", numericCode: "694", symbol: "Le", shortSymbol: "", decimalPlaces: 2, name: "Leone"},
	"SGD": {alphabeticCode: "SGD", numericCode: "702", symbol: "S$", shortSymbol: "", decimalPlaces: 2, name: "Singapore Dollar"},
	"XSU": {alphabeticCode: "XSU", numericCode: "994", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Sucre"},
	"SBD": {alphabeticCode: "SBD", numericCode: "090", symbol: "SI$", shortSymbol: "", decimalPlaces: 2, name: "Solomon Islands Dollar"},
	"SOS": {alphabeticCode: "SOS", numericCode: "706", symbol: "Sh.", shortSymbol: "", decimalPlaces: 2, name: "Somali Shilling"},
	"SSP": {alphabeticCode: "SSP", numericCode: "728", symbol: "SS£", shortSymbol: "", decimalPlaces: 2, name: "South Sudanese Pound"},
	"LKR": {alphabeticCode: "LKR", numericCode: "144", symbol: "Rs", shortSymbol: "", decimalPlaces: 2, name: "Sri Lanka Rupee"},
	"SDG": {alphabeticCode: "SDG", numericCode: "938", symbol: "£SD", shortSymbol: "", decimalPlaces: 2, name: "Sudanese Pound"},
	"SRD": {alphabeticCode: "SRD", numericCode: "968", symbol: "$", shortSymbol: "", decimalPlaces: 2, name: "Surinam Dollar"},
	"SEK": {alphabeticCode: "SEK", numericCode: "752", symbol: "kr", shortSymbol: "", decimalPlaces: 2, name: "Swedish Krona"},
	"CHE": {alphabeticCode: "CHE", numericCode: "947", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "WIR Euro"},
	"CHW": {alphabeticCode: "CHW", numericCode: "948", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "WIR Franc"},
	"SYP": {alphabeticCode: "SYP", numericCode: "760", symbol: "£S", shortSymbol: "", decimalPlaces: 2, name: "Syrian Pound"},
	"TWD": {alphabeticCode: "TWD", numericCode: "901", symbol: "NT$", shortSymbol: "$", decimalPlaces: 2, name: "New Taiwan Dollar"},
	"TJS": {alphabeticCode: "TJS", numericCode: "972", symbol: "SM", shortSymbol: "", decimalPlaces: 2, name: "Somoni"},
	"TZS": {alphabeticCode: "TZS", numericCode: "834", symbol: "TSh", shortSymbol: "", decimalPlaces: 2, name: "Tanzanian Shilling"},
	"THB": {alphabeticCode: "THB", numericCode: "764", symbol: "฿", shortSymbol: "", decimalPlaces: 2, name: "Baht"},
	"TOP": {alphabeticCode: "TOP", numericCode: "776", symbol: "T$", shortSymbol: "", decimalPlaces: 2, name: "Pa’anga"},
	"TTD": {alphabeticCode: "TTD", numericCode: "780", symbol: "TT$", shortSymbol: "", decimalPlaces: 2, name: "Trinidad and Tobago Dollar"},
	"TND": {alphabeticCode: "TND", numericCode: "788", symbol: "DT", shortSymbol: "", decimalPlaces: 3, name: "Tunisian Dinar"},
	"TRY": {alphabeticCode: "TRY", numericCode: "949", symbol: "YTL", shortSymbol: "", decimalPlaces: 2, name: "Turkish Lira"},
	"TMT": {alphabeticCode: "TMT", numericCode: "934", symbol: "m", shortSymbol: "", decimalPlaces: 2, name: "Turkmenistan New Manat"},
	"UGX": {alphabeticCode: "UGX", numericCode: "800", symbol: "USh", shortSymbol: "", decimalPlaces: 0, name: "Uganda Shilling"},
	"UAH": {alphabeticCode: "UAH", numericCode: "980", symbol: "₴", shortSymbol: "", decimalPlaces: 2, name: "Hryvnia"},
	"AED": {alphabeticCode: "AED", numericCode: "784", symbol: "د.إ", shortSymbol: "", decimalPlaces: 2, name: "UAE Dirham"},
	"USN": {alphabeticCode: "USN", numericCode: "997", symbol: "US$", shortSymbol: "$", decimalPlaces: 2, name: "US Dollar (Next day)"},
	"UYU": {alphabeticCode: "UYU", numericCode: "858", symbol: "$U", shortSymbol: "$", decimalPlaces: 2, name: "Peso Uruguayo"},
	"UYI": {alphabeticCode: "UYI", numericCode: "940", symbol: "", shortSymbol: "", decimalPlaces: 0, name: "Uruguay Peso en Unidades Indexadas (UI)"},
	"UYW": {alphabeticCode: "UYW", numericCode: "927", symbol: "", shortSymbol: "", decimalPlaces: 4, name: "Unidad Previsional"},
	"UZS": {alphabeticCode: "UZS", numericCode: "860", symbol: "сум", shortSymbol: "", decimalPlaces: 2, name: "Uzbekistan Sum"},
	"VUV": {alphabeticCode: "VUV", numericCode: "548", symbol: "VT", shortSymbol: "", decimalPlaces: 0, name: "Vatu"},
	"VES": {alphabeticCode: "VES", numericCode: "928", symbol: "Bs.S", shortSymbol: "", decimalPlaces: 2, name: "Bolívar Soberano"},
	"VND": {alphabeticCode: "VND", numericCode: "704", symbol: "₫", shortSymbol: "", decimalPlaces: 0, name: "Dong"},
	"YER": {alphabeticCode: "YER", numericCode: "886", symbol: "﷼", shortSymbol: "", decimalPlaces: 2, name: "Yemeni Rial"},
	"ZMW": {alphabeticCode: "ZMW", numericCode: "967", symbol: "ZK", shortSymbol: "", decimalPlaces: 2, name: "Zambian Kwacha"},
	"ZWL": {alphabeticCode: "ZWL", numericCode: "932", symbol: "", shortSymbol: "", decimalPlaces: 2, name: "Zimbabwe Dollar"},
	"XBA": {alphabeticCode: "XBA", numericCode: "955", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Composite Unit (EURCO)"},
	"XBB": {alphabeticCode: "XBB", numericCode: "956", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)"},
	"XBC": {alphabeticCode: "XBC", numericCode: "957", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)"},
	"XBD": {alphabeticCode: "XBD", numericCode: "958", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)"},
	"XTS": {alphabeticCode: "XTS", numericCode: "963", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Codes specifically reserved for testing purposes"},
	"XXX": {alphabeticCode: "XXX", numericCode: "999", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "The codes assigned for transactions where no currency is involved"},
	"XAU": {alphabeticCode: "XAU", numericCode: "959", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Gold"},
	"XPD": {alphabeticCode: "XPD", numericCode: "964", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Palladium"},
	"XPT": {alphabeticCode: "XPT", numericCode: "962", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Platinum"},
	"XAG": {alphabeticCode: "XAG", numericCode: "961", symbol: "", shortSymbol: "", decimalPlaces: -1, name: "Silver"},
}
