package database

import (
	"github.com/bela333/Vigne/errors"
	"github.com/go-redis/redis"
)

type Config struct {
	Database *Database
}

func (d *Database) Config() (*Config, error) {
	config := Config{}
	config.Database = d
	//Getting hasConfig
	hasConfig, err := d.Redis.Get(d.Decorate("hasConfig")).Int()
	if hasConfig != 1 || err == redis.Nil {
		return nil, errors.NoConfig
	}
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (d *Database) CreateConfig() error {
	//Set hasConfig
	err := d.Redis.Set(d.Decorate("hasConfig"), 1, 0).Err()
	if err != nil {
		return err
	}
	err = d.Redis.HMSet(d.Decorate("config"), map[string]interface{}{
		"token": "Bot 123456789.abcdEFGH",
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (config *Config) Token() string {
	return config.Database.Redis.HGet(config.Database.Decorate("config"), "token").Val()
}