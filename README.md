# D3money ![testing state](https://github.com/Dadido3/D3money/actions/workflows/test.yml/badge.svg?branch=master)

A library to handle monetary values and currencies.

This library is still in development.
While it works and most of its functionality is tested by unit tests, the API may have breaking changes until version 1.0.0 is reached.

## Features

- Uses [shopspring/decimal](https://github.com/shopspring/decimal) for arbitrary precision fixed-point decimal numbers.
- Values are immutable by default.
- ISO 4217 currencies.
- Extensible with custom currencies.
- Tests to ensure uniqueness and correctness of currencies (including user-defined ones).
- Useful mathematical operations, including a way to split a monetary value into n parts.
- Data bindings for JSON, binary, text, gob encodings.
- Implements scanner and valuer interfaces for databases.
- Implements `GormDBDataTypeInterface`.
- Supports postgresql composite types.

Planned:

- [ ] Include cryptocurrencies and tokens.
- [ ] Split locale and currency information.
- [ ] Formatting of values with specified locale. Support CLDR as good as possible, also `go generate` support to convert LDML based data into go code and structures.
- [ ] Migration field for currencies, e.g. to describe how custom currencies will map to official supported currencies.
- [ ] Generate currency data from the official ISO 4217 sources via `go generate`.

## What this is not

A high performance library to do number crunching with.
Although this library is not slow, it is not intended for processing large tables of monetary values.
The focus of this library is on correctness and ease of use.

## Usage

To be able to use this library, first download it via

```shell
go get github.com/Dadido3/D3money
```

And then import it with

```go
import money "github.com/Dadido3/D3money"
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

If you add custom currencies, make sure to only use negative unique IDs to prevent ID collisions with official currencies in the future.
