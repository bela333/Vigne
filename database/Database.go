package database

import "github.com/go-redis/redis"

type Database struct {
	Redis *redis.Client
	Identifier string
}

func NewDatabase(identifier, address, password string) *Database {
	d := Database{}
	d.Identifier = identifier
	d.Redis = redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
	})
	return &d
}

func (d *Database) Decorate(value string) string {
	return d.Identifier + ":" + value
}