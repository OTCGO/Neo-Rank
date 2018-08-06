package services

import (
	"Neo-Rank/config"
	"Neo-Rank/db/mgo"
	"Neo-Rank/models"
)

type AddressService struct {
}

const collectionAddress = "address"

var db = config.GetConfig().Mongodb.Database

func (addressService *AddressService) Find(limit int, skip int, condition map[string]interface{}) (ads []models.Address, err error) {
	session := mgo.GetSession()
	//执行完就关闭
	defer session.Close()
	err = session.DB(db).C(collectionAddress).Find(condition).Limit(limit).Skip(skip).All(&ads)
	return
}

func (addressService *AddressService) Count() (count int, err error) {
	session := mgo.GetSession()
	//执行完就关闭
	defer session.Close()
	count, err = session.DB(db).C(collectionAddress).Find(nil).Count()
	return
}
