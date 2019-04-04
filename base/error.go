package base

import "fmt"

// ValidationError abstrai um erro identificável em Hyperledger Fabric
type ValidationError string

// NewValidationError cria um novo erro de validação
func NewValidationError(message string, args ...interface{}) error {
	return ValidationError(fmt.Sprintf(message, args))
}

// Error retorna uma versão legível do erro
func (v ValidationError) Error() string {
	return string(v)
}
