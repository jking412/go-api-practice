package config

import (
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"go-api-practice/helpers"
	"os"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	viper = viperlib.New()

	viper.SetConfigType("env")

	viper.AddConfigPath(".")

	viper.SetEnvPrefix("appenv")

	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	loadEnv(env)
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := envPath + envSuffix
		if _, err := os.Stat(filePath); err != nil {
			envPath = filePath
		}
	}

	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

func GetStringMapString(path string, defaultValue ...interface{}) map[string]string {
	return cast.ToStringMapString(internalGet(path, defaultValue...))
}
