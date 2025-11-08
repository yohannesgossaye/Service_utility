package bills

import (
	"context"
	"time"

	billmodels "users/internal/domain/bills/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTransactionsStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTransactionsMongo(uri, dbName, collectionName string) (*MongoTransactionsStorage, error) {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &MongoTransactionsStorage{
		client:     client,
		collection: db.Collection(collectionName),
	}, nil
}

func (m *MongoTransactionsStorage) InsertTxn(ctx context.Context, txn billmodels.Transaction) error {
	if txn.ID == "" {
		txn.ID = primitive.NewObjectID().Hex()
	}

	if txn.Timestamp.IsZero() {
		txn.Timestamp = time.Now()
	}

	_, err := m.collection.InsertOne(ctx, txn)
	return err
}
