# D3money ![testing state](https://github.com/Dadido3/D3money/actions/workflows/test.yml/badge.svg?branch=master)

A library to handle monetary values and currencies.

## Features

- Uses [shopspring/decimal](https://github.com/shopspring/decimal) for arbitrary-precision fixed-point decimal numbers
- Immutable by default for safety
- ISO 4217 currencies
- Extendable to more/custom currencies without risking overlaps of codes or IDs
- Tests to ensure uniqueness and validity of currencies
- Data-bindings for JSON, Binary, Text, Gob encodings
- Implements scanner and valuer interfaces for databases
- Implements `GormDBDataTypeInterface`

Planned:

- [ ] List of cryptocurrencies
- [ ] Formatting of values with specified locale

## Usage

To be able to use this library, first download it via

```shell
go get github.com/Dadido3/D3money
```

And then import it with

```go
import "github.com/Dadido3/D3money"
```

### Values

A simple way to create a monetary value is by using `FromString(...)`.

```go
value1, err := money.FromString("123.45 ISO4217-EUR") // Value with ISO4217 EUR as currency.
value2, err := money.FromString("123.45")             // Value without currency or unit.
```

It's also possible to create a value by using a currency object.

```go
value, err := money.FromStringAndCurrency("123.45", money.Currencies.ByUniqueCode("ISO4217-EUR"))
```

For non user-input strings, the `MustFrom...` variants can be used.
They will not return any error, but panic if something is wrong.

```go
value1 := money.MustFromString("123.45 ISO4217-EUR")
value2 := money.MustFromStringAndCurrency("123.45 ISO4217-EUR", money.Currencies.ByUniqueCode("ISO4217-EUR"))
value3 := money.MustFromString("123.45 FOO-BAR") // Will panic if FOO-BAR is not a registered currency.
```

### Currencies

Selecting from all available currencies.

```go
eur := money.Currencies.ByUniqueCode("ISO4217-EUR")
```

Selecting a currency from ISO 4217 by its code or unique code.

```go
usd := money.ISO4217Currencies.ByCode("USD")
eur := money.ISO4217Currencies.ByUniqueCode("ISO4217-EUR")
```

Assert currency standard.

```go
eur := money.Currencies.ByUniqueCode("ISO4217-EUR")
_, isISO4217 := eur.(money.ISO4217Currency) // isISO4217 will be true
```

### Custom currencies

To create custom currencies, you need to create a type that implements the `money.Currency` interface.
For an example, see [currency-iso4217.go](currency-iso4217.go).

You can check if your instances are valid by using.

```go
err := money.ValidateCurrency(customCurrency)
```

Afterwards you can register the currency by adding it to the library by using.

```go
err := money.Currencies.Add(customCurrency)      // Register single custom currency.
err := money.Currencies.Add(customCurrencies...) // Register list of custom currencies.
```
