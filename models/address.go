package models

import "gopkg.in/mgo.v2/bson"

type Address struct {
	Id_        bson.ObjectId `json:"_id" bson:"_id"`
	Address    string        `json:"address" bson:"address"`
	BlockIndex int32         `json:"blockIndex" bson:"blockIndex"`
}
