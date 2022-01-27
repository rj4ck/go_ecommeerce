package models

type Users struct {
	Role        string `bson:"role,omitempty"`
	Email       string `bson:"email,omitempty"`
	LastName    string `bson:"lastName,omitempty"`
	FirstName   string `bson:"firstName,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`
}
