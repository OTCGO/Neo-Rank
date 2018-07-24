package services

import (
	"Neo-Rank/db/mgo"
	"Neo-Rank/models"
)

type AssetService struct {
}

const collectionAsset = "asset"

// const db = "testnet-node"

func (assetService *AssetService) Find(limit int, skip int, condition map[string]interface{}) (ast []models.Asset, err error) {
	session := mgo.GetSession()
	//
	defer session.Close()
	err = session.DB(db).C(collectionAsset).Find(condition).Limit(limit).Skip(skip).All(&ast)
	return
}
