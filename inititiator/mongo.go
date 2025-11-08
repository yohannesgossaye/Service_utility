package initiator

import (
	"fmt"
	"users/internal/storage"
	billsStorage "users/internal/storage/bills"
	mo "users/internal/storage/users"
)

type Persistence struct {
	users        storage.UsersRep
	transactions storage.Billstxn
}

func InitPersistence(uri, dbName, usersColl, transactionsColl string) (Persistence, error) {
	usersStore, err := mo.NewMong(uri, dbName, usersColl)
	if err != nil {
		return Persistence{}, fmt.Errorf("failed to initialize users Mongo: %w", err)
	}

	transactionsStore, err := billsStorage.NewTransactionsMongo(uri, dbName, transactionsColl)
	if err != nil {
		return Persistence{}, fmt.Errorf("failed to initialize transactions Mongo: %w", err)
	}

	return Persistence{
		users:        usersStore,
		transactions: transactionsStore,
	}, nil
}
