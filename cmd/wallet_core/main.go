package main

import (
	"database/sql"
	"fmt"
	"github.com/C4st3ll4n/wallet/internal/database"
	"github.com/C4st3ll4n/wallet/internal/event"
	"github.com/C4st3ll4n/wallet/internal/usecase/account"
	"github.com/C4st3ll4n/wallet/internal/usecase/client"
	"github.com/C4st3ll4n/wallet/internal/usecase/transaction"
	"github.com/C4st3ll4n/wallet/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf&parseTime=True&loc=True", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//dispatcher.Register("TransactionCreated", transacrionHandler)

	clientDb := database.NewClientDB(db)
	acountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUsecase := client.NewCreateClientUseCase(clientDb)
	createAccountUsecase := account.NewCreateAccountUsecase(acountDb, clientDb)
	createTransactionUsecase := transaction.NewCreateTransaction(transactionDb, acountDb, dispatcher, transactionCreatedEvent)
}
