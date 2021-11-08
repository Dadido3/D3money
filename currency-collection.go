package d3money

import (
	"fmt"
)

// CurrencyCollection provides a common interface for different currency standards. E.g. ISO 4217 or Cryptocurrencies.
type CurrencyCollection interface {
	Name() string // Name returns the name of the currency collection.

	Currencies() []Currency                          // Currency returns the list of currencies that are contained in this collection.
	CurrencyByUniqueID(uniqueID int32) Currency      // CurrencyByUniqueID finds a currency by its unique ID (e.g. 42170978).
	CurrencyByUniqueCode(uniqueCode string) Currency // CurrencyByUniqueCode finds a currency by its unique code (e.g. "ISO4217-EUR").
	CurrencyByCode(code string) Currency             // CurrencyByCode finds a currency by its code (e.g. "EUR"). This may not yield a result, as the code is not unique across currency standards. Best is to use it only in combination with a collection of a single standard, like ISO4217Currencies.

	AddCurrencies(c ...Currency) error // AddCurrencies adds one or more currencies to this collection.
}

// currencyCollectionSet implements CurrencyCollection.
// It's basically a set of currencies.
type currencyCollectionSet struct {
	name string

	currencies   []Currency
	byUniqueID   map[int32]Currency
	byUniqueCode map[string]Currency
	byCode       map[string]Currency
}

// Make sure this type implements the CurrencyCollection interface.
var _ CurrencyCollection = (*currencyCollectionSet)(nil)

// NewCurrencyCollection takes one or more list of currencies and returns them as a collection.
func NewCurrencyCollection(name string, listsOfCurrencies ...[]Currency) (CurrencyCollection, error) {
	cc := &currencyCollectionSet{
		name:         name,
		byUniqueID:   map[int32]Currency{},
		byUniqueCode: map[string]Currency{},
		byCode:       map[string]Currency{},
	}

	for _, listOfCurrencies := range listsOfCurrencies {
		if err := cc.AddCurrencies(listOfCurrencies...); err != nil {
			return nil, fmt.Errorf("failed to add currency to collection: %w", err)
		}
	}

	return cc, nil
}

// MustNewCurrencyCollection takes one or more list of currencies and returns them as a collection.
// It will panic on any error.
func MustNewCurrencyCollection(name string, listsOfCurrencies ...[]Currency) CurrencyCollection {
	cc, err := NewCurrencyCollection(name, listsOfCurrencies...)
	if err != nil {
		panic(fmt.Sprintf("failed to create currency collection %q: %v", name, err))
	}

	return cc
}

// CombineCurrencyCollections takes one ore more currency collections, and returns them as one combined collection.
// TODO: Consider removal of CombineCurrencyCollections
func CombineCurrencyCollections(name string, listOfCollections ...CurrencyCollection) (CurrencyCollection, error) {
	var currencies []Currency

	for _, collection := range listOfCollections {
		currencies = append(currencies, collection.Currencies()...)
	}

	return NewCurrencyCollection(name, currencies)
}

// MustCombineCurrencyCollections takes one ore more currency collections, and returns them as one combined collection.
// It will panic on any error.
func MustCombineCurrencyCollections(name string, listOfCollections ...CurrencyCollection) CurrencyCollection {
	cc, err := CombineCurrencyCollections(name, listOfCollections...)
	if err != nil {
		panic(fmt.Sprintf("failed to create combined currency collection %q: %v", name, err))
	}

	return cc
}

// Name returns the name of the currency collection.
func (cc *currencyCollectionSet) Name() string {
	return cc.name
}

// Currency returns the list of currencies that are contained in this collection.
func (cc *currencyCollectionSet) Currencies() []Currency {
	return cc.currencies
}

// CurrencyByUniqueID finds a currency by its unique ID (e.g. 42170978).
func (cc *currencyCollectionSet) CurrencyByUniqueID(uniqueID int32) Currency {
	return cc.byUniqueID[uniqueID]
}

// CurrencyByUniqueCode finds a currency by its unique code (e.g. "ISO4217-EUR").
func (cc *currencyCollectionSet) CurrencyByUniqueCode(uniqueCode string) Currency {
	return cc.byUniqueCode[uniqueCode]
}

// CurrencyByCode finds a currency by its code (e.g. "EUR").
// This may not yield a result, as the code is not unique across currency standards.
// Best is to use it only in combination with a collection of a single standard, like ISO4217Currencies.
func (cc *currencyCollectionSet) CurrencyByCode(code string) Currency {
	if cc.byCode != nil {
		return cc.byCode[code]
	}
	return nil
}

// addCurrency adds a currency to a collection.
func (cc *currencyCollectionSet) addCurrency(c Currency) error {
	uniqueID, uniqueCode, Code := c.UniqueID(), c.UniqueCode(), c.Code()

	// Check if the currency already exists.
	// Prevent duplicate entries, but prevent ID or code collisions.
	currencyByUniqueID, foundByUniqueID := cc.byUniqueID[uniqueID]
	currencyByUniqueCode, foundByUniqueCode := cc.byUniqueCode[uniqueCode]

	if foundByUniqueID && foundByUniqueCode && currencyByUniqueID == c && currencyByUniqueCode == c {
		// The currency already exists, ignore it.
		return nil
	} else if foundByUniqueID && currencyByUniqueID != c {
		// There is another currency with the same unique ID.
		return fmt.Errorf("currency %q has the same unique ID %d as the already existing currency %q", c, uniqueID, currencyByUniqueID)
	} else if foundByUniqueCode && currencyByUniqueCode != c {
		// There is another currency with the same unique Code.
		return fmt.Errorf("currency %q has the same unique code %q as the already existing currency %q", c, uniqueCode, currencyByUniqueCode)
	}

	cc.currencies = append(cc.currencies, c)
	cc.byUniqueID[uniqueID] = c
	cc.byUniqueCode[uniqueCode] = c

	// Special case for currency codes, as currency codes are not unique across different currency standards.
	// This will remove the byCode map, if a single duplicate code entry is found.
	// As a result this will prevent users from searching in collections that can't be unique.
	if _, found := cc.byCode[Code]; found {
		cc.byCode = nil
	} else if cc.byCode != nil {
		cc.byCode[Code] = c
	}

	return nil
}

// AddCurrencies adds a currency to a collection.
func (cc *currencyCollectionSet) AddCurrencies(currencies ...Currency) error {
	for _, c := range currencies {
		if err := cc.addCurrency(c); err != nil {
			return err
		}
	}

	return nil
}
