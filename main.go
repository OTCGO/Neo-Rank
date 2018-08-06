package main

import (
	"Neo-Rank/block"
	"Neo-Rank/db/mgo"
	"Neo-Rank/services"
	"Neo-Rank/utils"
	"Neo-Rank/utils/redis"
	"fmt"
	"math"
	"time"

	"github.com/pkg/errors"
)

const (
	configPath = "./config"
	env        = "development"
)

var rd *redis.RedisCon

const maxRoutineNum int = 10

var pageCount int = 50 //每页数量
var offset int = 0

var total int = 0 // 总页数

func main() {
	fmt.Println("main.go")
	start := time.Now()

	mgo.NewMongo()
	rd = redis.NewRedis()

	GetCount()

	fmt.Printf("总页数:%d", total)
	// cfg, err := config.Load(configPath, env)
	// if err != nil {

	// }

	//两个channel，一个用来放置工作项，一个用来存放处理结果。

	jobs := make(chan int, 100)

	results := make(chan int, 100)

	// 开启三个线程，也就是说线程池中只有3个线程，实际情况下，我们可以根据需要动态增加或减少线程。

	for w := 1; w <= maxRoutineNum; w++ {

		go worker(w, jobs, results)

	}

	// 添加9个任务后关闭Channel

	// channel to indicate that's all the work we have.

	for j := 1; j <= total; j++ {

		jobs <- j

	}

	close(jobs)

	//获取所有的处理结果

	for a := 1; a <= total; a++ {

		<-results

	}

	//记录时间
	elapsed := time.Since(start)
	fmt.Println("App elapsed: ", elapsed)
}

//这个是工作线程，处理具体的业务逻辑，将jobs中的任务取出，处理后将处理结果放置在results中。

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		//fmt.Println("worker", id, "processing job", j)
		offset = pageCount * (j - 1)
		Start(offset, pageCount)
		results <- j
	}

}

func Start(offset int, count int) {

	fmt.Println("offset", offset)
	fmt.Println("count", count)

	addressService := &services.AddressService{}
	addList, err := addressService.Find(count, offset, nil)
	if err != nil {
		errors.New("addressService error")
	}

	for _, v := range addList {
		// fmt.Println('k', k)
		// fmt.Println('v', v.Address)
		Banlance(v.Address)
	}
}

func Banlance(address string) {

	var b = &block.Block{}

	as := &services.AssetService{}
	match := make(map[string]interface{})
	match["type"] = "nep5"

	list, _ := as.Find(match)

	// fmt.
	for _, v := range list {
		// fmt.Println("v", v)
		if v.Type == "nep5" {
			// fmt.Println("v", v)
			invoke, err := b.GetNep5Balance(v.AssetId, string(utils.BigOrLittle(utils.AddressToScripthash(address))))
			if err != nil {
				errors.New("GetNep5Balance error")
			}
			// fmt.Println("v%", invoke)
			if invoke.Result.Stack[0].Value != "" {
				// fmt.Println("Stack", invoke.Result.Stack[0].Value)
				// // s, _ := strconv.ParseFloat(invoke.Result.Stack[0].Value, 64)
				// s, _ := strconv.ParseInt(invoke.Result.Stack[0].Value, 16, 64)
				// balance := float64(s) / math.Pow10(8)

				// fmt.Println("balance", balance)
				// fmt.Println("s%", invoke.Result.Stack[0].Value)
				// fmt.Println("AssetId", v.AssetId)
				decimal, err := utils.GetDecimalsCache([]byte(v.AssetId))
				if err != nil {
					errors.New("decimal error")
				}
				// fmt.Println("decimal", decimal)

				balance := utils.HexToNumStr(invoke.Result.Stack[0].Value, int(decimal))
				// fmt.Println("balance", balance)
				if balance != 0 {
					// fmt.Println("balance", balance)
					_, err = rd.ZAdd(string(v.AssetId[2:]), balance, address)
					if err != nil {
						errors.New("redis error")
					}
				}
			}
		}
	}
}

func GetCount() {
	addressService := &services.AddressService{}
	counts, err := addressService.Count()
	if err != nil {
		errors.New("GetCount error")
	}
	fmt.Println("counts", counts)

	total = int(math.Ceil(float64(counts) / float64(pageCount))) //page总数

	return
}
