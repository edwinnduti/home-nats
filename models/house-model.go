package models

type House struct {
	Id                   string  `json:"id"`
	HouseNumber          uint64  `json:"houseNumber"`
	Name                 string  `json:"name"`
	Location             string  `json:"location"`
	Cost                 uint64  `json:"cost"`
	YearOfConstruction   uint32  `json:"yearOfConstruction"`
	Category             string  `json:"category"`
	Area                 float32 `json:"area"`
	Perimeter            float32 `json:"perimeter"`
	NumberOfFloors       int     `json:"numberOfFloors"`
	NumberOfBedrooms     int     `json:"numberOfBedrooms"`
	ConstructionMaterial string  `json:"constructionMaterial"`
	RoofingType          string  `json:"roofingType"`
	FenceType            string  `json:"fenceType"`
	ParkingLot           bool    `json:"parkingLot"`
	SourceOfWaterSupply  string  `json:"sourceOfWaterSupply"`
	HasWifi              bool    `json:"hasWifi"`
}
