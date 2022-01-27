package models

type Banners struct {
	Image    string `bson:"image,omitempty"`
	IsActive bool   `bson:"isActive,omitempty"`
}
