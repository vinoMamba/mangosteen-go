package config

import "github.com/vinoMamba/mangosteen-go/pkg/config"

func init() {
	config.Add("route", func() map[string]interface{} {
		return map[string]interface{}{
			"port":      config.Env("APP_ROUTE_PORT", "3001"),
			"localhost": config.Env("APP_ROUTE_LOCALHOST", "localhost"),
		}
	})
}
