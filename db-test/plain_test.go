package dbt

import (
	"math/rand"
	"testing"

	money "github.com/Dadido3/D3money"
	"github.com/shopspring/decimal"
)

func TestPlain(t *testing.T) {
	// Drop tables.
	if _, err := db.Exec("DROP TABLE IF EXISTS test_accounts"); err != nil {
		t.Errorf("Failed to drop table: %v", err)
		return
	}

	// Create tables.
	if _, err := db.Exec("CREATE TABLE test_accounts (id INTEGER, balance VARCHAR(255))"); err != nil {
		t.Errorf("Failed to create table: %v", err)
		return
	}

	// Create test entry.
	a1 := &TestAccount{
		ID:      1,
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), money.ISO4217Currencies.ByCode("EUR")),
	}

	// Create test entry.
	a2 := &TestAccount{
		ID:      2,
		Balance: money.FromDecimal(decimal.New(int64(rand.Intn(100000000)), -4), nil),
	}

	var stmtString string
	if *flagDBDriver == "pgx" {
		stmtString = "INSERT INTO test_accounts (id, balance) VALUES ($1, $2);"
	} else {
		stmtString = "INSERT INTO test_accounts (id, balance) VALUES (?, ?);"
	}

	// Insert test entry 1 into database.
	if _, err := db.Exec(stmtString, a1.ID, a1.Balance); err != nil {
		t.Log()
		t.Errorf("Failed to insert entry 1: %v", err)
	}

	// Insert test entry 2 into database.
	if _, err := db.Exec(stmtString, a2.ID, a2.Balance); err != nil {
		t.Errorf("Failed to insert entry 2: %v", err)
	}

	if *flagDBDriver == "pgx" {
		stmtString = "SELECT id, balance FROM test_accounts WHERE id = $1"
	} else {
		stmtString = "SELECT id, balance FROM test_accounts WHERE id = ?"
	}

	// Read test entries from database.
	stmt, err := db.Prepare(stmtString)
	if err != nil {
		t.Errorf("db.Prepare() failed: %v", err)
		return
	}
	defer stmt.Close()

	a1Read := new(TestAccount)

	err = stmt.QueryRow(1).Scan(&a1Read.ID, &a1Read.Balance)
	if err != nil {
		t.Errorf("stmt.QueryRow(1) failed: %v", err)
	}

	if equal, _ := a1.Balance.Equal(a1Read.Balance); !equal {
		t.Errorf("a1Read.Balance = %v, want %v", a1Read.Balance, a1.Balance)
	}

	a2Read := new(TestAccount)

	err = stmt.QueryRow(2).Scan(&a2Read.ID, &a2Read.Balance)
	if err != nil {
		t.Errorf("stmt.QueryRow(2) failed: %v", err)
	}

	if equal, _ := a2.Balance.Equal(a2Read.Balance); !equal {
		t.Errorf("a2Read.Balance = %v, want %v", a2Read.Balance, a2.Balance)
	}

}
