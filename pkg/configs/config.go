package config

import (
	"fmt"
	"log"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"interview/pkg/logger"
)

var conf = &Config{}

type Config struct {
	// Server config
	Server struct {
		Port    int    `env:"SERVER_PORT,default=8089"`
		Host    string `env:"HOST,default=localhost"`
		GinMode string `env:"GIN_MODE,default=debug"`
	}

	// Database config
	Database struct {
		Host     string `env:"DB_HOST,default=localhost"`
		Port     int    `env:"DB_PORT,default=4001"`
		User     string `env:"DB_USER,default=ic_user"`
		Password string `env:"DB_PASS,default=9xz3jrd8wf"`
		Name     string `env:"DB_NAME,default=ice_db"`
	}

	// Env config
	Env struct {
		Environment string `env:"ENV"`
		LogLevel    string `env:"LOG_LEVEL"`
	}

	// Logging config
	logger *logger.Entry
}

func (c *Config) DBConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}

func Load() {
	lg := logger.General.Component("config")
	err := godotenv.Load(".env.local")
	if err != nil {
		err = godotenv.Load(".env")
	}

	if err != nil {
		lg.Info("Error loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(conf)

	if err != nil {
		log.Fatal(err)
	}

	conf.logger = lg
}

func Get() *Config {
	return conf
}
