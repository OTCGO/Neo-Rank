package services

import (
	"Neo-Rank/db/mgo"
	"Neo-Rank/models"
)

type AddressService struct {
}

const collectionAddress = "address"
const db = "testnet-node"

func (addressService *AddressService) Find(limit int, skip int) (ads []models.Address, err error) {
	session := mgo.GetSession()
	//执行完就关闭
	defer session.Close()
	err = session.DB(db).C(collectionAddress).Find(nil).Limit(limit).Skip(skip).All(&ads)
	return
}
