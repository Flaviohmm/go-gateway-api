package service

import (
	"github.com/Flaviohmm/imersao22/go-gateway/internal/domain"
	"github.com/Flaviohmm/imersao22/go-gateway/internal/dto"
)

// AccountService é um serviço para gerenciar operações de conta.
type AccountService struct {
	repository domain.AccountRepository
}

// NewAccountService cria um novo serviço de conta com o repositório fornecido.
func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

// CreateAccount cria uma nova conta e valida duplicidade de API Key
// Retorna ErrDuplicateAPIKey se a API Key já existir.
func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	// Valida duplicidade de API Key
	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicateAccount
	}

	// Salva a conta no banco de dados
	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	// Converte a conta para o formato de saída
	output := dto.FromAccount(account)
	return &output, nil
}

// UpdateBalance atualiza o saldo de uma conta de forma thread-safe
// O amount pode ser positivo (crédito) ou negativo (débito)
func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

// FindByAPIKey busca uma conta pela API Key
func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

// FindByID busca uma conta pelo ID
func (s *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}
