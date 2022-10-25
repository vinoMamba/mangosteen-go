package config

import (
	"fmt"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"github.com/vinoMamba/mangosteen-go/pkg/helpers"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.SetEnvPrefix("appenv")

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	loadEnv(env)
	loadConfig()
}

func loadEnv(envSufix string) {
	envPath := ".env"
	if len(envSufix) > 0 {
		envPath = envPath + "." + envSufix
	}
	viper.SetConfigFile(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Env(envName string, defaultValue string) interface{} {
	if !viper.IsSet(envName) && helpers.IsEmpty(viper.Get(envName)) {
		fmt.Println("env not set: ", envName)
		return defaultValue
	}
	return viper.Get(envName)
}

func Get(path string) string {
	return cast.ToString(viper.Get(path))
}
