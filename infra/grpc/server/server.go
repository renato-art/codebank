package server

import (
	"log"
	"net"

	"github.com/simplemoney/codebank/infra/grpc/pb"
	"github.com/simplemoney/codebank/infra/grpc/service"
	"github.com/simplemoney/codebank/usecase"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tpc", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("could not listen tpc port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)
}
