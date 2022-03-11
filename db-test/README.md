# Database test package

This package is used to test the DB data bindings of [github.com/Dadido3/D3money](github.com/Dadido3/D3money).

For now the following database drivers are supported:

- `sqlite3`
- `pgx` (PostgreSQL)

## Run tests locally

Running the test via `go test` will use an in memory sqlite database for testing.
To run a test with another database driver use

``` shell
go test --db-driver pgx --db-dsn "host=localhost port=5432 user=test password=test dbname=test sslmode=disable"
```

This assumes that you have a local PostgreSQL database running at port 5432, with a user `test` that has access to a database named `test`.
