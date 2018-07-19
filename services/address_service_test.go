package services

import (
	"Neo-Rank/db/mgo"
	"testing"
)

func init() {
	mgo.NewMongo()
}
func TestFind(t *testing.T) {
	list, _ := Find(10, 0)

	t.Log("list", list)
}
