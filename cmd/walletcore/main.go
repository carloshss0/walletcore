package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/carloshss0/walletcore/internal/database"
	"github.com/carloshss0/walletcore/internal/event"
	createaccount "github.com/carloshss0/walletcore/internal/usecase/create_account"
	createclient "github.com/carloshss0/walletcore/internal/usecase/create_client"
	createtransaction "github.com/carloshss0/walletcore/internal/usecase/create_transaction"
	"github.com/carloshss0/walletcore/internal/web"
	"github.com/carloshss0/walletcore/internal/web/webserver"
	"github.com/carloshss0/walletcore/pkg/events"
	"github.com/carloshss0/walletcore/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3307", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)


	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func (tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func (tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})
	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransacationUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webServer := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransacationUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webServer.Start()

}
