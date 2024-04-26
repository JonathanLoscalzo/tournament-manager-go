package domain

import "time"

type Person struct {
	FirstName   string
	LastName    string
	Birthday    time.Time
	Nationality string

	Address   Address
	Cellphone string
	Email     string

	EmergencyContacts []EmergencyContact
}

type Address struct {
	Street       string
	Number       string
	Apartment    string
	Intersection string
	PostalCode   string
	City         string
	State        string
	Country      string
}

type EmergencyContact struct {
	FullName  string
	Cellphone string
	Email     string
	Address   string
}
