package config

import "github.com/vinoMamba/mangosteen-go/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("APP_REDIS_ADDRESS", "localhost"),
			"port":     config.Env("APP_REDIS_PORT", "6379"),
			"password": config.Env("APP_REDIS_PASSWORD", ""),
			"username": config.Env("APP_REDIS_USER", ""),
		}
	})
}
