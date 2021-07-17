package server

import "github.com/Cristianoaf81GIT/codebank/usecase"

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}
