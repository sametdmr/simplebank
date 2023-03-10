package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sametdmr/simplebank/utility"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utility.IsSupportedCurrency(currency)
	}
	return false
}
