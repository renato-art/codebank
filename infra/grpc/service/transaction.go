package service

import "github.com/simplemoney/codebank/usecase"

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}
