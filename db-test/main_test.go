package dbt

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"testing"

	_ "github.com/jackc/pgx"
	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	flagDBDriver         = flag.String("db-driver", "sqlite3", "The database driver to use")
	flagDBDataSourceName = flag.String("db-dsn", ":memory:", "The data source name")
)

// db is the normal database instance for the whole db-tests package.
var db *sql.DB

// gormDB is the GORM database instance for the whole db-tests package.
var gormDB *gorm.DB

// Set up some stuff before the tests run.
func TestMain(m *testing.M) {

	flag.Parse()
	driver, dsn := *flagDBDriver, *flagDBDataSourceName

	// Create normal DB instance.
	var err error
	if db, err = sql.Open(driver, dsn); err != nil {
		panic(fmt.Sprintf("sql.Open(%q, %q) failed: %v", driver, dsn, err))
	}
	defer db.Close()

	if db == nil {
		panic("Failed to open database instance")
	}

	// Create GORM DB instance.
	switch driver {
	case "sqlite3":
		var err error
		if gormDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{}); err != nil {
			panic(fmt.Sprintf("gorm.Open(sqlite.Open(%q)...) failed: %v", dsn, err))
		}

	case "pgx":
		var err error
		if gormDB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{}); err != nil {
			panic(fmt.Sprintf("gorm.Open(postgres.New(postgres.Config{DSN: %q})...) failed: %v", dsn, err))
		}

	default:
		panic(fmt.Sprintf("Unknown database driver %q", driver))
	}

	if gormDB == nil {
		panic("Failed to open GORM database instance")
	}

	// Run tests.
	os.Exit(m.Run())
}
