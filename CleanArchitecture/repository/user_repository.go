package repository

import (
	"context"

	"github.com/naol86/learn-go/CleanArchitecture/domain"
	"github.com/naol86/learn-go/CleanArchitecture/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// Fetch implements domain.UserRepository.
func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection);
	opts := options.Find().SetProjection(bson.M{"password": 0})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []domain.User;
	err = cursor.All(c, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByEmail implements domain.UserRepository.
func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection);
	var user domain.User;
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

// GetByID implements domain.UserRepository.
func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection);
	var user domain.User;
	index, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = collection.FindOne(c, primitive.M{"_id": index}).Decode(&user)
	return user, err
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
