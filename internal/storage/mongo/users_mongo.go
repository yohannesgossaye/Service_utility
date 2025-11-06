package mongo

import (
	"errors"
	"time"
	"users/internal/domain/dto"
	"users/internal/domain/models"

	"context"

	"go.mongodb.org/mongo-driver/bson"

	Act_gen "users/pkgs/utils/helper"

	"users/pkgs/auth"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUsersStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMong(uri, dbName, collectionName string) (*MongoUsersStorage, error) {
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
	return &MongoUsersStorage{
		client:     client,
		collection: db.Collection(collectionName),
	}, nil
}

func (m *MongoUsersStorage) CreateUser(ctx context.Context, user models.Users) (models.Users, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if user.Account_number == "" {
		user.Account_number = Act_gen.GenerateAccountNumber()
	}
	user.Is_active = false

	create, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		return models.Users{}, err
	}

	if oid, ok := create.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid.Hex()
	}
	return user, nil

}

type userDoc struct {
	ID             primitive.ObjectID `bson:"_id"`
	FullName       string             `bson:"fullname"`
	Email          string             `bson:"email"`
	Phone_number   string             `bson:"phone_number"`
	Account_number string             `bson:"account_number"`
	Balance        float64            `bson:"balance"`
	Is_active      bool               `bson:"is_active"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}

func (m *MongoUsersStorage) GetUser(ctx context.Context, id string) (models.Users, error) {
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Users{}, errors.New("invalid user id")
	}

	var doc userDoc
	err = m.collection.FindOne(ctx, bson.M{"_id": idparam}).Decode(doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Users{}, errors.New("user not found")
		}
		return models.Users{}, err
	}
	user := models.Users{
		ID:        doc.ID.Hex(),
		FullName:  doc.FullName,
		Email:     doc.Email,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
	}
	return user, nil
}

func (m *MongoUsersStorage) GetUsers(ctx context.Context) ([]models.Users, error) {
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var userDocs []userDoc
	if err = cursor.All(ctx, &userDocs); err != nil {
		return nil, err
	}

	users := make([]models.Users, len(userDocs))
	for i, doc := range userDocs {
		users[i] = models.Users{
			ID:        doc.ID.Hex(),
			FullName:  doc.FullName,
			Email:     doc.Email,
			CreatedAt: doc.CreatedAt,
			UpdatedAt: doc.UpdatedAt,
		}
	}

	return users, nil
}

func (m *MongoUsersStorage) UpdateUsers(ctx context.Context, id string, updateduser *models.Users) (models.Users, error) {

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Users{}, errors.New("invalid user id is provided")
	}
	updateduser.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"fullname":  updateduser.FullName,
			"Email":     updateduser.Email,
			"UpdatedAt": updateduser.UpdatedAt,
		},
	}

	res, err := m.collection.UpdateOne(ctx, bson.M{"_id": objid}, update)
	if err != nil {
		return models.Users{}, err
	}
	if res.MatchedCount == 0 {
		return models.Users{}, errors.New("user not found")
	}
	updateduser.ID = id
	return *updateduser, nil
}

func (m *MongoUsersStorage) DeleteUser(ctx context.Context, id string) (string, error) {
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", errors.New("invalid user id is provided")
	}

	res, err := m.collection.DeleteOne(ctx, bson.M{"_id": idparam})
	if err != nil {
		return "", errors.New("deleted is failed")
	}
	if res.DeletedCount == 0 {
		return "", errors.New("user not found")
	}

	return "deleted sucessfully", nil
}

func (m *MongoUsersStorage) EnableAccount(ctx context.Context, id string) (string, error) {
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", errors.New("invalid user id is provided")
	}

	res, err := m.collection.UpdateOne(
		ctx,
		bson.M{"_id": idparam},
		bson.M{"$set": bson.M{"is_active": true}},
	)
	if err != nil {
		return "", errors.New("enable failed")
	}

	if res.ModifiedCount == 0 {
		return "", errors.New("user not found or already enabled")
	}

	return "enabled successfully", nil
}

func (m *MongoUsersStorage) LoginUser(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var user models.Users
	err := m.collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}
	if user.Password != req.Password {
		return nil, errors.New("password is not correct")
	}
	// generate the token
	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token:   token,
		Message: "Login sucessfully ",
	}, nil
}
