package parking

import (
	"encoding/json"
	"io"
	"time"

	"github.com/tejaswayadav/CarParkingSystemMicroservice/models/vehicles"
)

// Parking is a struct for Parking spaces
type Parking struct {
	ParkingID        string                   `json:"parkingId"`
	Name             string                   `json:"name"`
	Capacity         int                      `json:"capacity"`
	HeavyVehicleSafe bool                     `json:"ifHeavyVehicleSafe"`
	Allotment        map[string]*vehicles.Car `json:"alltoment"`
	CreatedOn        time.Time                `json:"-"`
	UpdatedOn        time.Time                `json:"-"`
}

// FromJSON converts json to a Parking instance
func (p *Parking) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

// ToJSON converts Parking instance to JSON
func (p *Parking) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// AddCarToParking add a car to parking collection
func AddCarToParking() error {
	//collection := client.Database("CarParkingSystem").Collection("Parkings")
	return nil
}
