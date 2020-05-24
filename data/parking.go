package data

type Parking struct {
	Name             string `json:"name"`
	Capacity         int    `json:"capacity"`
	HeavyVehicleSafe bool   `json:"ifHeavyVehicleSafe"`
}
