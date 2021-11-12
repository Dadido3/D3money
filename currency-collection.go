// Copyright (c) 2021 David Vogel
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package money

import (
	"fmt"
)

// CurrencyCollection provides a common interface for different currency standards. E.g. ISO 4217 or Cryptocurrencies.
type CurrencyCollection interface {
	Name() string // Name returns the name of the currency collection.

	All() []Currency                         // All returns the full list of currencies that are contained in this collection.
	ByUniqueID(uniqueID int32) Currency      // ByUniqueID finds a currency by its unique ID (e.g. 42170978).
	ByUniqueCode(uniqueCode string) Currency // ByUniqueCode finds a currency by its unique code (e.g. "ISO4217-EUR").
	ByCode(code string) Currency             // ByCode finds a currency by its code (e.g. "EUR"). This may not yield a result, as the code is not unique across currency standards. Best is to use it only in combination with a collection of a single standard, like ISO4217Currencies.

	Add(c ...Currency) error // Add adds one or more currencies to this collection.
}

// currencyCollectionSet implements CurrencyCollection.
// It's basically a set of currencies.
type currencyCollectionSet struct {
	name string

	all          []Currency
	hasCurrency  map[Currency]struct{}
	byUniqueID   map[int32]Currency
	byUniqueCode map[string]Currency
	byCode       map[string]Currency
}

// Make sure this type implements the CurrencyCollection interface.
var _ CurrencyCollection = (*currencyCollectionSet)(nil)

// NewCurrencyCollection takes one or more lists of currencies and returns them as a collection.
//
// singleStandard defines whether the collection contains currencies of just a single standard.
// If you combine multiple standards into a collection, set it to false.
func NewCurrencyCollection(name string, singleStandard bool, listsOfCurrencies ...[]Currency) (CurrencyCollection, error) {
	cc := &currencyCollectionSet{
		name:         name,
		hasCurrency:  map[Currency]struct{}{},
		byUniqueID:   map[int32]Currency{},
		byUniqueCode: map[string]Currency{},
		byCode:       nil,
	}

	if singleStandard {
		cc.byCode = map[string]Currency{}
	}

	for _, listOfCurrencies := range listsOfCurrencies {
		if err := cc.Add(listOfCurrencies...); err != nil {
			return nil, fmt.Errorf("failed to add currency to collection: %w", err)
		}
	}

	return cc, nil
}

// MustNewCurrencyCollection takes one or more list of currencies and returns them as a collection.
// It will panic on any error.
func MustNewCurrencyCollection(name string, enableCode bool, listsOfCurrencies ...[]Currency) CurrencyCollection {
	cc, err := NewCurrencyCollection(name, enableCode, listsOfCurrencies...)
	if err != nil {
		panic(fmt.Sprintf("failed to create currency collection %q: %v", name, err))
	}

	return cc
}

// Name returns the name of the currency collection.
func (cc *currencyCollectionSet) Name() string {
	return cc.name
}

// All returns the full list of currencies that are contained in this collection.
func (cc *currencyCollectionSet) All() []Currency {
	return cc.all
}

// ByUniqueID finds a currency by its unique ID (e.g. 42170978).
func (cc *currencyCollectionSet) ByUniqueID(uniqueID int32) Currency {
	return cc.byUniqueID[uniqueID]
}

// ByUniqueCode finds a currency by its unique code (e.g. "ISO4217-EUR").
func (cc *currencyCollectionSet) ByUniqueCode(uniqueCode string) Currency {
	return cc.byUniqueCode[uniqueCode]
}

// ByCode finds a currency by its code (e.g. "EUR").
// This may not yield a result, as the code is not unique across currency standards.
// Best is to use it only in combination with a collection of a single standard, like ISO4217Currencies.
func (cc *currencyCollectionSet) ByCode(code string) Currency {
	if cc.byCode != nil {
		return cc.byCode[code]
	}
	return nil
}

// add adds a currency to this collection.
func (cc *currencyCollectionSet) add(c Currency) error {
	uniqueID, uniqueCode, code := c.UniqueID(), c.UniqueCode(), c.Code()

	// Check if the currency already exists.
	// Ignore duplicate entries, but prevent collisions of unique IDs, unique codes or codes.
	currencyByUniqueID, foundByUniqueID := cc.byUniqueID[uniqueID]
	currencyByUniqueCode, foundByUniqueCode := cc.byUniqueCode[uniqueCode]
	var currencyByCode Currency
	var foundByCode bool
	if cc.byCode != nil {
		currencyByCode, foundByCode = cc.byCode[code]
	}

	if foundByUniqueID && currencyByUniqueID != c {
		// There is another currency with the same unique ID.
		return fmt.Errorf("currency %q has the same unique ID %d as the already existing currency %q", c, uniqueID, currencyByUniqueID)
	} else if foundByUniqueCode && currencyByUniqueCode != c {
		// There is another currency with the same unique code.
		return fmt.Errorf("currency with unique ID %d has the same unique code %q as the already existing currency with unique ID %d", uniqueID, uniqueCode, currencyByUniqueCode.UniqueID())
	} else if foundByCode && currencyByCode != c {
		// There is another currency with the same code.
		return fmt.Errorf("currency with unique ID %d has the same code %q as the already existing currency with unique ID %d", uniqueID, code, currencyByCode.UniqueID())
	}

	if _, found := cc.hasCurrency[c]; found {
		// Currency is already in the set. Ignore it.
		return nil
	}

	// TODO: If singleStandard is set for NewCurrencyCollection, don't allow multiple standards

	cc.all = append(cc.all, c)
	cc.hasCurrency[c] = struct{}{}
	cc.byUniqueID[uniqueID] = c
	cc.byUniqueCode[uniqueCode] = c
	if cc.byCode != nil {
		cc.byCode[code] = c
	}

	return nil
}

// Add adds one or more currencies to this collection.
func (cc *currencyCollectionSet) Add(currencies ...Currency) error {
	for _, c := range currencies {
		if err := cc.add(c); err != nil {
			return err
		}
	}

	return nil
}
