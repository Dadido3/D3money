package dbt

import (
	"math/rand"
	"testing"

	money "github.com/Dadido3/D3money"
	"github.com/shopspring/decimal"
)

func TestCreateQuery(t *testing.T) {
	// Delete tables.
	if err := gormDB.Migrator().DropTable(&Account{}); err != nil {
		t.Errorf("Failed to drop tables: %v", err)
	}

	// Create tables from structures.
	if err := gormDB.AutoMigrate(&Account{}); err != nil {
		t.Errorf("GORM auto migration failed: %v", err)
	}

	a1 := Account{
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), money.ISO4217Currencies.ByCode("EUR")),
	}
	if err := gormDB.Create(&a1).Error; err != nil {
		t.Errorf("Failed to create account 1: %v", err)
	}

	a2 := Account{
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), nil),
	}
	if err := gormDB.Create(&a2).Error; err != nil {
		t.Errorf("Failed to create account 2: %v", err)
	}

	var a1Read Account
	if err := gormDB.First(&a1Read, a1.ID).Error; err != nil {
		t.Errorf("Failed to query account 1: %v", err)
	}

	if equal, err := a1.Balance.Equal(a1Read.Balance); err != nil {
		t.Errorf("a1.Balance.Equal(a1Read.Balance) returned error: %v", err)
	} else if !equal {
		t.Errorf("Queried balance %v doesn't match written balance %v", a1Read.Balance, a1.Balance)
	}

	var a2Read Account
	if err := gormDB.First(&a2Read, a2.ID).Error; err != nil {
		t.Errorf("Failed to query account 2: %v", err)
	}

	if equal, err := a2.Balance.Equal(a2Read.Balance); err != nil {
		t.Errorf("a2.Balance.Equal(a2Read.Balance) returned error: %v", err)
	} else if !equal {
		t.Errorf("Queried balance %v doesn't match written balance %v", a2Read.Balance, a2.Balance)
	}
}
