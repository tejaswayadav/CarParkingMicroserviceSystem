package databases

import "github.com/tejaswayadav/CarParkingSystemMicroservice/models"

// CarParkSystemDatabase is any database that can store cars
type CarParkSystemDatabase interface {
	Find(identifier, collectionName string) (*models.CarParkingSystemModel, error)
	Add(model *models.CarParkingSystemModel) error
	Remove(identifier, collectionName string) error
}
