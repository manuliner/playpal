package models

type Item struct {
	ID    string  `bson:"_id,omitempty" json:"id"`
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}
