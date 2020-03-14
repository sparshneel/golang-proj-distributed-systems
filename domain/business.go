package domain

import(

)

type Business struct{
	Id 				string
	Name            string
    BusinessAddress string
	Telephone       int64
}

type Parking struct{
	Location 	string
	Capacity 	int64
	ParkingType string
}
