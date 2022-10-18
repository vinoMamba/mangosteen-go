package bootstrap

import (
	"fmt"

	"github.com/vinoMamba/mangosteen-go/pkg/config"
	"github.com/vinoMamba/mangosteen-go/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.Get("redis.host"), config.Get("redis.port")),
		config.Get("redis.username"),
		config.Get("redis.password"),
		0,
	)
}
