package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tile struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Color string             `json:"color" bson:"color"`
}

type UpdateTiles struct {
	Added   []Tile               `json:"added" bson:"added"`
	Deleted []primitive.ObjectID `json:"deleted" bson:"deleted"`
	Changed []Tile               `json:"changed" bson:"changed"`
}
