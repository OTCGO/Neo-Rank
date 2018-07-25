package main

import (
	"Neo-Rank/block"
	"Neo-Rank/db/mgo"
	"Neo-Rank/services"
	"Neo-Rank/utils"
	"fmt"
	"math"
)

const (
	configPath = "./config"
	env        = "development"
)

func main() {
	fmt.Println("main.go")
	// cfg, err := config.Load(configPath, env)
	// if err != nil {

	// }

	mgo.NewMongo()

	var b = &block.Block{}

	as := &services.AssetService{}
	list, _ := as.Find(0, 0, nil)

	for _, v := range list {
		// fmt.Println("v", v)
		if v.Type == "nep5" {
			fmt.Println("v", v)
			invoke, err := b.GetNep5Balance(v.AssetId, utils.BigOrLittle(utils.AddressToScripthash("AGwJpXGPowiJfMFAdnrdB1uV92i4ubPANA")))
			if err != nil {

			}
			fmt.Println("v%", invoke)
		}
	}

	fmt.Println(math.Pow10(8))
	// t.Log("list", list)

}
