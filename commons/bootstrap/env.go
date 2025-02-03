package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	MongoURI string `mapstructure:"MONGO_URI"`
	MongoDB  string `mapstructure:"MONGO_DB"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}
	return &env
}
