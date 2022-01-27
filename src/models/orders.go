package models

type Orders struct {
	Product  string `bson:"product,omitempty"`
	Quantity string `bson:"quantity,omitempty"`
	Deliver  string `bson:"deliver,omitempty"`
}
