package mgo

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// type Mgo struct {
// 	sessionMongo *mgo.Session
// }

var sessionMongo *mgo.Session

func NewMongo() {
	conn := "127.0.0.1:27017"
	// fmt.Println("conn", conn)
	session, err := mgo.Dial(conn)
	//设置连接池个数
	session.SetPoolLimit(300)
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
