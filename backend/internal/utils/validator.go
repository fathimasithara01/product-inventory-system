package utils

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ValidateUUID checks if a string is a valid UUID
func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// ValidateQuantity parses a string to decimal and checks if it is > 0
func ValidateQuantity(qty string) (decimal.Decimal, error) {
	value, err := decimal.NewFromString(qty)
	if err != nil {
		return decimal.Zero, errors.New("invalid quantity format")
	}

	if value.LessThanOrEqual(decimal.Zero) {
		return decimal.Zero, errors.New("quantity must be greater than zero")
	}

	return value, nil
}
