package main

import (
	"Neo-Rank/block"
	"Neo-Rank/db/mgo"
	"Neo-Rank/services"
	"Neo-Rank/utils"
	"fmt"
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

	// fmt.
	for _, v := range list {
		// fmt.Println("v", v)
		if v.Type == "nep5" {
			// fmt.Println("v", v)
			invoke, err := b.GetNep5Balance(v.AssetId, string(utils.BigOrLittle(utils.AddressToScripthash("AGwJpXGPowiJfMFAdnrdB1uV92i4ubPANA"))))
			if err != nil {

			}
			// fmt.Println("v%", invoke)
			if invoke.Result.Stack[0].Value != "" {
				// fmt.Println("Stack", invoke.Result.Stack[0].Value)
				// // s, _ := strconv.ParseFloat(invoke.Result.Stack[0].Value, 64)
				// s, _ := strconv.ParseInt(invoke.Result.Stack[0].Value, 16, 64)
				// balance := float64(s) / math.Pow10(8)

				// fmt.Println("balance", balance)
				// fmt.Println("s%", invoke.Result.Stack[0].Value)
				fmt.Println("AssetId", v.AssetId)
				decimal, err := utils.GetDecimalsCache([]byte(v.AssetId))
				if err != nil {

				}
				fmt.Println("decimal", decimal)

				balance := utils.HexToNumStr(invoke.Result.Stack[0].Value, int(decimal))
				fmt.Println("balance", balance)
			}
		}
	}

	// fmt.Println(math.Pow10(8))
	// t.Log("list", lis
}
