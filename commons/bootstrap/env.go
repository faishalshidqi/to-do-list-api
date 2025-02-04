package bootstrap

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Env struct {
	ServerAddress                string        `mapstructure:"SERVER_ADDRESS"`
	MongoURI                     string        `mapstructure:"MONGO_URI"`
	MongoDB                      string        `mapstructure:"MONGO_DB"`
	ContextTimeout               int           `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenKey               string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey              string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenExpirationInHour  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRATION_IN_HOUR"`
	RefreshTokenExpirationInHour time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRATION_IN_HOUR"`
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
