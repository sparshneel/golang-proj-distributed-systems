package domain

import(

)

type Business struct{
	Id 				string
	Name            string
    BusinessAddress *Address
	Telephone       int64
}

type Address struct{
	Street 	string
	City 	string
	Postcode int64
}

type Parking struct{
	Location 	string
	Capacity 	int64
	ParkingType string
}
