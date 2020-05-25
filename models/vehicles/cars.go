package vehicles

import (
	"encoding/json"
	"io"
	"time"
)

// Car is struct for Cars
type Car struct {
	CarID        string    `json:"carId"`
	Name         string    `json:"name"`
	Manufacturer string    `json:"manufacturer"`
	LicensePlate string    `json:"licensePlate"`
	CreatedOn    time.Time `json:"-"`
	UpdatedOn    time.Time `json:"-"`
}

// CarList is a list of Cars
type CarList []*Car

// FromJSON converts json to a Parking instance
func (p *Car) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

// ToJSON converts Parking instance to JSON
func (p *Car) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}
