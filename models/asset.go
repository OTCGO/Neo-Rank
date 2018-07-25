package models

import "gopkg.in/mgo.v2/bson"

type Asset struct {
	Id_        bson.ObjectId `json:"_id" bson:"_id"`
	AssetId    string        `json:"assetId" bson:"assetId"`
	BlockIndex int32         `json:"blockIndex" bson:"blockIndex"`
	Type       string        `json:"type" bson:"type"`
}
