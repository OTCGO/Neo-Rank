package services

import (
	"Neo-Rank/db/mgo"
	"testing"
)

func init() {
	mgo.NewMongo()
}
func TestFindAsst(t *testing.T) {

	as := &AssetService{}
	list, _ := as.Find(10, 0, nil)

	t.Log("list", list)
	///list, _ := Find(10, 0)

	// t.Log("list", list)
}
