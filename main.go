package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/simplemoney/codebank/infra/grpc/server"
	"github.com/simplemoney/codebank/infra/kafka"
	"github.com/simplemoney/codebank/infra/repository"
	"github.com/simplemoney/codebank/usecase"
)

func main() {
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer("host.docker.internal:9094")
	return producer
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	grpcServer.Serve()
}

// cc := domain.NewCreditCard()
// 	cc.Number = "1234"
// 	cc.Name = "test card"
// 	cc.ExpirationYear = 2023
// 	cc.ExpirationMonth = 8
// 	cc.CVV = 123
// 	cc.Limit = 1000
// 	cc.Balance = 0

// 	repo := repository.NewTransactionRepositoryDb(db)
// 	repo.CreateCreditCard(*cc)
