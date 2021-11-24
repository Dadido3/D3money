package dbt

import (
	"time"

	money "github.com/Dadido3/D3money"
)

type Account struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Balance money.Value
}
