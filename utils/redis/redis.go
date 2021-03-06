package redis

import (
	"Neo-Rank/config"

	"github.com/go-redis/redis"
)

var client *redis.Client
var InitErr error

func NewRedis(mgoRedis config.Redis) *RedisCon {
	client, InitErr = redisInit(mgoRedis)
	return &RedisCon{}
}

type RedisCon struct {
}

func redisInit(mgoRedis config.Redis) (client *redis.Client, err error) {

	client = redis.NewClient(&redis.Options{
		Addr:     mgoRedis.Addr,
		Password: mgoRedis.Password, // no password set
		DB:       mgoRedis.DB,       // use default DB
	})
	// switch config.GetEnvironment() {
	// case "DEVELOPMENT":
	// 	client = redis.NewClient(&redis.Options{
	// 		Addr:     config.RedisAddrDev,
	// 		Password: "", // no password set
	// 		DB:       0,  // use default DB
	// 	})
	// 	break
	// case "STAGING":
	// 	client = redis.NewFailoverClient(&redis.FailoverOptions{
	// 		MasterName:    config.MasterNameStaging,
	// 		SentinelAddrs: config.SentinelAddrsStaging,
	// 	})
	// 	break

	// case "PRODUCTION":
	// 	client = redis.NewFailoverClient(&redis.FailoverOptions{
	// 		MasterName:    config.MasterNamePro,
	// 		SentinelAddrs: config.SentinelAddrsPro,
	// 	})
	// 	break
	// }
	_, err = client.Ping().Result()

	return

}

func (*RedisCon) Get(key string) (string, error) {
	var value string
	var err error
	if InitErr != nil {
		return value, InitErr
	}

	value, err = client.Get(key).Result()
	return value, err
}

func (*RedisCon) Set(key string, value string) (ok bool, err error) {

	if InitErr != nil {
		return false, InitErr
	}

	err = client.Set(key, value, 0).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (*RedisCon) ZAdd(key string, score float64, member string) (ok bool, err error) {
	err = client.ZAdd(key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (*RedisCon) Sub(c string) (content string, err error) {
	pubsub := client.Subscribe(c)
	defer pubsub.Close()
	msg, err := pubsub.ReceiveMessage()

	// fmt.Println(msg.Channel, msg.Payload)
	return msg.Payload, err
	// fmt.Println(msg.Channel, msg.Payload)
}
