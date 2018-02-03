/*
共有两个 redis 数据结构

1、String，存放ip的具体信息，key：ip， value：存放 proxyJson 的json 形式
2、SortSet，存放健康度排行，score：健康度，member：ip
*/

package proxy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

var connRedis redis.Conn

// https://godoc.org/github.com/garyburd/redigo/redis#Conn

func dialRedis() {
	var err error
	connRedis, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	// defer connRedis.Close()
}

// Key 如果存在，则覆盖原始数据
func setRString(key, value string) error {
	if len(key) == 0 {
		return errors.New("Key error")
	}
	key = strings.Replace(key, ":", "-", -1)
	_, err := connRedis.Do("SET", key, value)
	return err
}

func getRString(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("Key error")
	}
	key = strings.Replace(key, ":", "-", -1)
	s, err := redis.String(connRedis.Do("GET", key))
	s = strings.Replace(s, "-", ":", -1)
	return fmt.Sprintf("%v", s), err
}

// 如果某个成员已经是有序集的成员，那么更新这个成员的分数值
func setRSortSet(socre int, key string) error {
	if len(key) == 0 {
		return errors.New("Key error")
	}
	_, err := connRedis.Do("ZADD", redisZaddName, socre, key)
	return err
}

func setRSortSets(s map[string]int) error {
	var err error
	var e string
	for k, v := range s {
		err = setRSortSet(v, k)
		if err != nil {
			e = e + fmt.Sprintf("%v:%v", k, err)
			continue
		}
	}
	return errors.New(e)
}

// 递减排序，不获得 分数
func getRSortSet(start, end int) ([]string, error) {
	if start < end {
		t := end
		end = start
		start = t
	}
	if start < 0 {
		start = 0
	}
	members, err := redis.Strings(connRedis.Do("ZREVRANGE", redisZaddName, start, end))
	// members, err := redis.Strings(c.Do("ZREVRANGE", redisZaddName, start, end, "WITHSCORES"))
	return members, err
}

// 获得 SortSet 中的数量，如果不存在也返回0
func getRSortSetNum() int {
	i, _ := redis.Int(connRedis.Do("ZCARD", redisZaddName))
	return i
}
