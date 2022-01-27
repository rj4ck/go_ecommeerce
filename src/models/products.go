package models

type Products struct {
	Name        string `bson:"name,omitempty"`
	Price       int    `bson:"price,omitempty"`
	Image       string `bson:"image,omitempty"`
	IsActive    bool   `bson:"isActive,omitempty"`
	Category    string `bson:"category,omitempty"`
	Description string `bson:"description,omitempty"`
}
