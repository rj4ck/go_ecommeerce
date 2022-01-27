package models

type Customers struct {
	Email       string `bson:"email,omitempty"`
	LastName    string `bson:"lastName,omitempty"`
	FirstName   string `bson:"firstName,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`
}
