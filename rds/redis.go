package rds

import (
	"errors"
	"log"

	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Fatalln("error connection redis: ", err)
			}
			return c, err
		},
	}
}

func Ping(c redis.Conn) error {
	pong, err := c.Do("PING")
	if err != nil {
		return err
	}

	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}

	log.Println("PING response: ", s)
	return nil
}

func Set(c redis.Conn, key string, value string) error {
	_, err := c.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func Get(c redis.Conn, key string) (string, error) {
	res, err := redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		return "", errors.New("Key does not exist")
	} else if err != nil {
		return "", err
	}
	return res, nil
}
