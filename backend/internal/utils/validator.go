package utils

import "github.com/shopspring/decimal"

func ValidateQuantity(qty string) (decimal.Decimal, error) {
	value, err := decimal.NewFromString(qty)
	if err != nil {
		return decimal.Zero, err
	}

	if value.LessThanOrEqual(decimal.Zero) {
		return decimal.Zero, err
	}

	return value, nil
}
