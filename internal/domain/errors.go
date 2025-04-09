package domain

import "errors"

var (
	// ErrAccountNotFound é retornada quando uma conta não é encontrada.
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicateAccount é retornada quando há uma tentativa de criar uma conta com API Key duplicada.
	ErrDuplicateAccount = errors.New("api key already exists")
	// ErrInvoiceNotFound é retornado quando uma fatura não é encontrada.
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess é retornado quando há tentativa de acesso não autorizado a um recurso.
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
