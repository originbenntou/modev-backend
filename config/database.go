package config

import (
	"os"
)

const (
	ConnMaxLifetimeSec = 120
	MaxOpenConns       = 100
	MaxIdleConns       = 100
)

const (
	DefaultDBName = "modev"
)

type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Name     string
}

func loadDatabaseConfig(conf *Config) {
	if conf.Env == "test" {
		conf.Database.Host = os.Getenv("TEST_DB_HOST")
		conf.Database.User = os.Getenv("TEST_DB_USER")
		conf.Database.Password = os.Getenv("TEST_DB_PASSWORD")
		conf.Database.Port = 3307
		conf.Database.Name = DefaultDBName
	} else {
		conf.Database.Host = os.Getenv("DB_HOST")
		conf.Database.User = os.Getenv("DB_USER")
		conf.Database.Password = os.Getenv("DB_PASSWORD")
		conf.Database.Port = 3306
		conf.Database.Name = DefaultDBName
	}
}
