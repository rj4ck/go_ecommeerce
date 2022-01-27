package models

type Categories struct {
	Name  string `bson:"name,omitempty"`
	Index int    `bson:"index,omitempty"`
}
