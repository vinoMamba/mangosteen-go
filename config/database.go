package config

import "github.com/vinoMamba/mangosteen-go/pkg/config"

func init() {
	config.Add("db", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("APP_DB_HOST", "localhost"),
			"port":     config.Env("APP_DB_PORT", "5432"),
			"user":     config.Env("APP_DB_USER", "postgres"),
			"dbname":   config.Env("APP_DB_NAME", "postgres"),
			"password": config.Env("APP_DB_PASSWORD", "123456"),
		}
	})
}
