package mongodb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/tejaswayadav/CarParkingSystemMicroservice/databases"
	"github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"
	"github.com/tejaswayadav/CarParkingSystemMicroservice/models"
	"github.com/tejaswayadav/CarParkingSystemMicroservice/models/vehicles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoStorage struct {
	dbClient *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	loggers.GetLogger().InfoLogger.Println("Establishing a new MongoDB connection with ", mongoURL)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	loggers.GetLogger().InfoLogger.Println("Trying to ping the MongoDB client")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	loggers.GetLogger().InfoLogger.Println("Ping to MongoDB successful")
	return client, nil
}

// NewMongoStorage creates a new mongoStorage instance, adds a specified database to it and returns it
func NewMongoStorage(mongoURL, mongoDBName string, mongoTimeout int) (databases.CarParkSystemDatabase, error) {
	storage := &mongoStorage{
		database: mongoDBName,
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	mongoClient, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "databases.NewMongoStorage")
	}
	storage.dbClient = mongoClient
	return storage, nil
}

func (m *mongoStorage) Find(identifier, collectionName string) (*models.CarParkingSystemModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	model := &vehicles.Car{}
	collection := m.dbClient.Database(m.database).Collection(collectionName)
	filter := bson.M{"carId": model.CarID}
	err := collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		if err == models.ErrorModelNotFound {
			return nil, errors.Wrap(err, "databases.mongostorage.find")
		}
	}

	return nil, nil
}

func (m *mongoStorage) Add(model *models.CarParkingSystemModel) error {
	return nil
}

func (m *mongoStorage) Remove(identifier, collectionName string) error {
	return nil
}
