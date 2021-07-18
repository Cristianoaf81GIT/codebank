package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Cristianoaf81GIT/codebank/infrastructure/grpc/server"
	"github.com/Cristianoaf81GIT/codebank/infrastructure/kafka"
	"github.com/Cristianoaf81GIT/codebank/infrastructure/repository"
	"github.com/Cristianoaf81GIT/codebank/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error connection to database")
	}
	return db
}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
	return producer
}

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("running gRPC server")
	grpcServer.Serve()
}

// cc := domain.NewCreditCard()
// cc.Number = "1234"
// cc.Name = "cristiano"
// cc.ExpirationYear = 2021
// cc.ExpirationMonth = 12
// cc.CVV = 123
// cc.Limit = 2000
// cc.Balance = 0

// repo := repository.NewTransactionRepositoryDb(db)
// err := repo.CreateCreditCard(*cc)
// if err != nil {
// 	fmt.Println(err)
// }
