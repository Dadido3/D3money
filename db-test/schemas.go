package dbt

import (
	"time"

	money "github.com/Dadido3/D3money"
)

type TestAccount struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Balance money.Value
}

type TestAccountCompositeType struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Balance money.Value `gorm:"type:d3money"`
}
