package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Cristianoaf81GIT/codebank/domain"
	"github.com/Cristianoaf81GIT/codebank/infrastructure/repository"
	"github.com/Cristianoaf81GIT/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "cristiano"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 12
	cc.CVV = 123
	cc.Limit = 2000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "root", "codebank")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error connection to database")
	}
	return db
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}
