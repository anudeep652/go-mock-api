package api

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Rc struct {
	redisClient *redis.Client
}

func initDb() *Rc {
	// initialize redis client
	var rdb Rc
	rdb.redisClient = redis.NewClient(&redis.Options{
		Addr:     "go-mock-api_redis_1:6379", // server can access redis container by it's hostname
		Password: "",                         // no password set
		DB:       0,                          // use default DB
	})
	fmt.Println(rdb.redisClient)
	return &rdb
}

func (rc *Rc) SignUp(username string, password string, ctx context.Context) error {
	st := rc.redisClient.JSONSet(ctx, "user:"+username, ".", User{Name: username, Pass: hashPassword(password)})
	if st.Err() != nil {
		fmt.Println(st.Err())
		return st.Err()
	}
	return nil
}

func (rc *Rc) GetAllUsers() {
	st := rc.redisClient.Keys(context.Background(), "user:*")
	fmt.Println(st.Val())
}
