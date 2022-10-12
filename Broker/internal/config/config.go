package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"mr-l0n3lly/go-broker/pkg/logging"
	"sync"
)

type Config struct {
	Database struct {
		Mongo struct {
			Host string `env:"MONGO_HOST"`
			Port int    `env:"MONGO_PORT"`
		}
		Redis struct {
			Host string `env:"REDIS_HOST"`
			Port int    `env:"REDIS_PORT"`
		}
	}
	SocketServer struct {
		Host string `env:"BROKER_HOST"`
		Port int    `env:"BROKER_TCP_PORT"`
	}
	GrpcServer struct {
		Host string `env:"BROKER_HOST"`
		Port int    `env:"BROKER_GRPC_PORT"`
	}
}

var instance *Config
var once sync.Once

func GetConfiguration() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")

		instance = &Config{}

		err := cleanenv.ReadConfig(".env", instance)

		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Info(err)
		}
	})

	return instance
}
