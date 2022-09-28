package tool

import (
	"myblog-gin/utils"

	"github.com/go-redis/redis"
)

type RedisStore struct {
	Client *redis.Client
}

var RediStore RedisStore

func InitRedisStore() (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.ReAddr + ":" + utils.RePort,
		Password: utils.RePassWord,
		DB:       utils.ReDb,
	})

	RediStore = RedisStore{Client: client}

	_, err = RediStore.Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// // set
// func (rs *RedisStore) Set(id string, value string) {
// 	err := rs.client.Set(id, value, time.Minute*10).Err()
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// // get
// func (rs *RedisStore) Get(id string, clear bool) string {
// 	val, err := rs.client.Get(id).Result()
// 	if err != nil {
// 		log.Println(err)
// 		return ""
// 	}
// 	if clear {
// 		err := rs.client.Del(id).Err()
// 		if err != nil {
// 			log.Println(err)
// 			return ""
// 		}
// 	}
// 	return val
// }
