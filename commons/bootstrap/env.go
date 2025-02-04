package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	MongoURI       string `mapstructure:"MONGO_URI"`
	MongoDB        string `mapstructure:"MONGO_DB"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
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
