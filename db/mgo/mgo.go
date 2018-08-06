package mgo

import (
	"Neo-Rank/config"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// type Mgo struct {
// 	sessionMongo *mgo.Session
// }

var sessionMongo *mgo.Session

func NewMongo(mgoCon config.Mongodb) {
	// mgoCon := config.GetConfig().Mongodb
	fmt.Printf("config%+v", mgoCon)
	// conn := "127.0.0.1:27017"
	dialInfo := &mgo.DialInfo{
		Addrs:     mgoCon.Addrs,
		Direct:    mgoCon.Direct,
		Timeout:   time.Second * 1,
		Database:  mgoCon.Database,
		Source:    mgoCon.Source,
		Username:  mgoCon.Username,
		Password:  mgoCon.Password,
		PoolLimit: mgoCon.PoolLimit,
	}
	// fmt.Println("conn", conn)
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("连接mongodb失败:%s\n", err)
	}

	sessionMongo = session

}

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func GetSession() *mgo.Session {
	if sessionMongo == nil {
		return nil
	}
	return sessionMongo.Clone()
}
