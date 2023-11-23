package config

import (
	"log"

	"github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig        AppConfig
	DatabaseConfig   DatabaseConfig
	RedisConfig      RedisConfig
	PrivatePublicKey PrivatePublicKey
}

type AppConfig struct {
	AppHost      string `env:"APP_HOST,required=true"`
	Port         string `env:"APP_PORT,required=true"`
	ReadTimeout  uint32 `env:"READ_TIMEOUT,required=true"`
	IdleTimeout  uint32 `env:"IDLE_TIMEOUT,required=true"`
	WriteTimeout uint32 `env:"WRITE_TIMEOUT,required=true"`
}

type DatabaseConfig struct {
	DbHost     string `env:"DB_HOST,required=true"`
	DbUser     string `env:"DB_USER,required=true"`
	DbPassword string `env:"DB_PASSWORD,required=true"`
	DbName     string `env:"DB_NAME,required=true"`
	DbPort     string `env:"DB_PORT,required=true"`
	Sslmode    string `env:"DB_SSLMODE,required=true"`
	TimeZone   string `env:"DB_TIMEZONE,required=true"`
}
type RedisConfig struct {
	RedisAddr     string `env:"REDIS_ADDRRESS,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDb       int    `env:"REDIS_DB,required=true"`
}
type PrivatePublicKey struct {
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
}

func init() {

	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Println(err)
	}
}

func GetConfig() Config {
	return cfg
}
