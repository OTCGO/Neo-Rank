package services

import (
	"Neo-Rank/db/mgo"
	"Neo-Rank/models"
)

const collection = "address"
const db = "testnet-node"

func Find(limit int, skip int) (ads []models.Address, err error) {
	session := mgo.GetSession()
	//执行完就关闭
	defer session.Close()
	err = session.DB(db).C(collection).Find(nil).Limit(limit).Skip(skip).All(&ads)
	return
}
