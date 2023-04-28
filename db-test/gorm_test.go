package dbt

import (
	"math/rand"
	"reflect"
	"testing"

	money "github.com/Dadido3/D3money"
	"github.com/shopspring/decimal"
)

func TestCreateQuery(t *testing.T) {
	// Delete tables.
	if err := gormDB.Migrator().DropTable(&TestAccount{}); err != nil {
		t.Fatalf("Failed to drop tables: %v", err)
	}

	// Create tables from structures.
	if err := gormDB.AutoMigrate(&TestAccount{}); err != nil {
		t.Fatalf("GORM auto migration failed: %v", err)
	}

	a1 := TestAccount{
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), money.ISO4217Currencies.ByCode("EUR")),
	}
	if err := gormDB.Create(&a1).Error; err != nil {
		t.Errorf("Failed to create account 1: %v", err)
	}

	a2 := TestAccount{
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), nil),
	}
	if err := gormDB.Create(&a2).Error; err != nil {
		t.Errorf("Failed to create account 2: %v", err)
	}

	var a1Read TestAccount
	if err := gormDB.First(&a1Read, a1.ID).Error; err != nil {
		t.Errorf("Failed to query account 1: %v", err)
	}

	if equal, err := a1.Balance.EqualDetailed(a1Read.Balance); err != nil {
		t.Errorf("a1.Balance.Equal(a1Read.Balance) returned error: %v", err)
	} else if !equal {
		t.Errorf("Queried balance %v doesn't match written balance %v", a1Read.Balance, a1.Balance)
	}

	var a2Read TestAccount
	if err := gormDB.First(&a2Read, a2.ID).Error; err != nil {
		t.Errorf("Failed to query account 2: %v", err)
	}

	if equal, err := a2.Balance.EqualDetailed(a2Read.Balance); err != nil {
		t.Errorf("a2.Balance.Equal(a2Read.Balance) returned error: %v", err)
	} else if !equal {
		t.Errorf("Queried balance %v doesn't match written balance %v", a2Read.Balance, a2.Balance)
	}
}

// TestCompositeType tests storing monetary values as composite type.
func TestCompositeType(t *testing.T) {
	// This works only with PostgreSQL.
	if *flagDBDriver != "pgx" {
		t.SkipNow()
	}

	// Delete tables.
	if err := gormDB.Migrator().DropTable(&TestAccountCompositeType{}); err != nil {
		t.Fatalf("Failed to drop tables: %v", err)
	}

	// Create custom composite type.
	if err := gormDB.Exec("DROP TYPE IF EXISTS d3money; CREATE TYPE d3money AS (amount DECIMAL, currency VARCHAR);").Error; err != nil {
		t.Fatalf("Failed to create d3money composite type: %v", err)
	}

	// Create tables from structures.
	if err := gormDB.AutoMigrate(&TestAccountCompositeType{}); err != nil {
		t.Fatalf("GORM auto migration failed: %v", err)
	}

	a1 := TestAccountCompositeType{
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), money.ISO4217Currencies.ByCode("EUR")),
	}
	if err := gormDB.Create(&a1).Error; err != nil {
		t.Errorf("Failed to create account 1: %v", err)
	}

	a2 := TestAccountCompositeType{
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), nil),
	}
	if err := gormDB.Create(&a2).Error; err != nil {
		t.Errorf("Failed to create account 2: %v", err)
	}

	var a1Read TestAccountCompositeType
	if err := gormDB.First(&a1Read, a1.ID).Error; err != nil {
		t.Errorf("Failed to query account 1: %v", err)
	}

	if equal, err := a1.Balance.EqualDetailed(a1Read.Balance); err != nil {
		t.Errorf("a1.Balance.Equal(a1Read.Balance) returned error: %v", err)
	} else if !equal {
		t.Errorf("Queried balance %v doesn't match written balance %v", a1Read.Balance, a1.Balance)
	}

	var a2Read TestAccountCompositeType
	if err := gormDB.First(&a2Read, a2.ID).Error; err != nil {
		t.Errorf("Failed to query account 2: %v", err)
	}

	if equal, err := a2.Balance.EqualDetailed(a2Read.Balance); err != nil {
		t.Errorf("a2.Balance.Equal(a2Read.Balance) returned error: %v", err)
	} else if !equal {
		t.Errorf("Queried balance %v doesn't match written balance %v", a2Read.Balance, a2.Balance)
	}

	// Test direct access to the fields of the composite type.
	type CompositeFields struct {
		Amount   decimal.Decimal
		Currency string
	}

	var a1ReadFields CompositeFields
	if err := gormDB.Raw("SELECT (balance).amount, (balance).currency FROM test_account_composite_types WHERE id = ?", 1).Scan(&a1ReadFields).Error; err != nil {
		t.Errorf("Failed to query account 1: %v", err)
	}

	a1ExpectedFields := CompositeFields{a1.Balance.Decimal(), a1.Balance.Currency().UniqueCode()}
	if !reflect.DeepEqual(a1ReadFields, a1ExpectedFields) {
		t.Errorf("Queried balance fields %+v don't match expected fields %+v", a1ReadFields, a1ExpectedFields)
	}

	var a2ReadFields CompositeFields
	if err := gormDB.Raw("SELECT (balance).amount, (balance).currency FROM test_account_composite_types WHERE id = ?", 2).Scan(&a2ReadFields).Error; err != nil {
		t.Errorf("Failed to query account 2: %v", err)
	}

	a2ExpectedFields := CompositeFields{a2.Balance.Decimal(), ""}
	if !reflect.DeepEqual(a2ReadFields, a2ExpectedFields) {
		t.Errorf("Queried balance fields %+v don't match expected fields %+v", a2ReadFields, a2ExpectedFields)
	}

}
