package d3money

import (
	"database/sql"
	"encoding/json"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestJSONMarshalling(t *testing.T) {
	val := MustFromString("-12345.67 ISO4217-EUR")

	jsonBytes, err := json.Marshal(val)
	if err != nil {
		t.Errorf("json.Marshal() failed: %v", err)
	}

	correct := `{"Value":"-12345.67","Currency":"ISO4217-EUR"}`
	if string(jsonBytes) != correct {
		t.Errorf("json.Marshal() = %q: want %q", jsonBytes, correct)
	}

	var val2 Value
	if err := json.Unmarshal(jsonBytes, &val2); err != nil {
		t.Errorf("json.Unmarshal() failed: %v", err)
	}

	if !val.Equal(val2) {
		t.Errorf("val2 = %v, want %v", val2, val)
	}
}

func TestJSONMarshalling2(t *testing.T) {
	val := MustFromString("-12345.67")

	jsonBytes, err := json.Marshal(val)
	if err != nil {
		t.Errorf("json.Marshal() failed: %v", err)
	}

	correct := `{"Value":"-12345.67"}`
	if string(jsonBytes) != correct {
		t.Errorf("json.Marshal() = %q: want %q", jsonBytes, correct)
	}

	var val2 Value
	if err := json.Unmarshal(jsonBytes, &val2); err != nil {
		t.Errorf("json.Unmarshal() failed: %v", err)
	}

	if !val.Equal(val2) {
		t.Errorf("val2 = %v, want %v", val2, val)
	}
}

func TestSQLite(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Errorf("Failed to open database: %v", err)
	}
	defer db.Close()

	type customer struct {
		ID      int
		Balance Value
	}

	_, err = db.Exec("create table customers (id integer, balance text)")
	if err != nil {
		t.Errorf("Failed to create table: %v", err)
	}

	// Create test entry.
	val := MustFromString("-12345.67 ISO4217-EUR")
	e := &customer{
		ID:      1,
		Balance: val,
	}

	// Create test entry.
	val2 := MustFromString("-12345.67")
	e2 := &customer{
		ID:      2,
		Balance: val2,
	}

	// Add test entry into database.
	_, err = db.Exec("insert into customers (id, balance) values (?, ?);", e.ID, e.Balance)
	if err != nil {
		t.Errorf("Failed to insert entry: %v", err)
	}

	// Add test entry into database.
	_, err = db.Exec("insert into customers (id, balance) values (?, ?);", e2.ID, e2.Balance)
	if err != nil {
		t.Errorf("Failed to insert entry: %v", err)
	}

	// Read test entries from database.
	stmt, err := db.Prepare("select id, balance from customers where id = ?")
	if err != nil {
		t.Errorf("db.Prepare() failed: %v", err)
	}
	defer stmt.Close()

	eRead := new(customer)

	err = stmt.QueryRow(1).Scan(&eRead.ID, &eRead.Balance)
	if err != nil {
		t.Errorf("stmt.QueryRow() failed: %v", err)
	}

	if !e.Balance.Equal(eRead.Balance) {
		t.Errorf("eRead.Balance = %v, want %v", eRead.Balance, e.Balance)
	}

	eRead2 := new(customer)

	err = stmt.QueryRow(2).Scan(&eRead2.ID, &eRead2.Balance)
	if err != nil {
		t.Errorf("stmt.QueryRow() failed: %v", err)
	}

	if !e2.Balance.Equal(eRead2.Balance) {
		t.Errorf("eRead2.Balance = %v, want %v", eRead2.Balance, e2.Balance)
	}

}
