package initiator

import (
	"fmt"
	"users/internal/storage"
	mo "users/internal/storage/users"
)

type Persistence struct {
	users storage.UsersRep
}

func InitPersistence(uri, dbName, collectionName string) (Persistence, error) {
	store, err := mo.NewMong(uri, dbName, collectionName)
	if err != nil {
		return Persistence{}, fmt.Errorf("failed to initialize Mongo: %w", err)
	}
	return Persistence{users: store}, nil
}
