package domain

import "github.com/gocql/gocql"

type Business struct{
	Id 				gocql.UUID
	Name            string
    BusinessAddress string
	Telephone       int64
	City   string
	State string
	Pincode string
}

type BusinessRep struct {
	Id 				string `json:"id"`
	Name            string `json:"name"`
	BusinessAddress string `json:"business-address"`
	Telephone       int64  `json:"telephone"`
	City   string  `json:"city"`
	State string   `json:"state"`
	Pincode string `json:"pincode"`
}

type Parking struct{
	Location 	string
	Capacity 	int64
	ParkingType string
}
