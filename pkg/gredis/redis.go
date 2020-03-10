package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/iprologue/myBlog/pkg/setting"
	"log"
	"time"
)

var RedisConn *redis.Pool

func SetUp() error {
	RedisConn = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := conn.Do("AUTH", setting.RedisSetting.Password); err != nil {
					err := conn.Close()
					if err != nil {
						log.Println(err)
					}
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         setting.RedisSetting.MaxIdle,
		MaxActive:       setting.RedisSetting.MaxActive,
		IdleTimeout:     setting.RedisSetting.IdleTimeout,
		Wait:            false,
		MaxConnLifetime: 0,
	}

	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return nil
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}


func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
