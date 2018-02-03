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
	"sync"

	"github.com/garyburd/redigo/redis"
)

var connRedis redis.Conn

var redisMutex sync.Mutex

// https://godoc.org/github.com/garyburd/redigo/redis#Conn

func dialRedis() {
	var err error
	connRedis, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	// defer connRedis.Close()
}

func closeRedis() {
	connRedis.Close()
}

// Key 如果存在，则覆盖原始数据
func setRString(key, value string) error {
	if len(key) == 0 {
		return errors.New("Key error")
	}
	key = strings.Replace(key, ":", "-", -1)
	redisMutex.Lock()
	_, err := connRedis.Do("SET", key, value)
	redisMutex.Unlock()
	return err
}

func getRString(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("Key error")
	}
	key = strings.Replace(key, ":", "-", -1)
	redisMutex.Lock()
	s, err := redis.String(connRedis.Do("GET", key))
	redisMutex.Unlock()
	s = strings.Replace(s, "-", ":", -1)
	return fmt.Sprintf("%v", s), err
}

// 删除 SortSet 中单个元素，成功删除则返回1，失败或者不存在返回0
func deleteRString(key string) int {
	if len(key) == 0 {
		return 0
	}
	key = strings.Replace(key, ":", "-", -1)
	redisMutex.Lock()
	a, _ := redis.Int(connRedis.Do("DEL", key))
	redisMutex.Unlock()
	return a
}

// 如果某个成员已经是有序集的成员，那么更新这个成员的分数值
func setRSortSet(socre int, key string) error {
	if len(key) == 0 {
		return errors.New("Key error")
	}
	redisMutex.Lock()
	_, err := connRedis.Do("ZADD", redisZaddName, socre, key)
	redisMutex.Unlock()
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
	if start > end {
		t := end
		end = start
		start = t
	}
	if start < 0 {
		start = 0
	}
	redisMutex.Lock()
	members, err := redis.Strings(connRedis.Do("ZREVRANGE", redisZaddName, start, end))
	redisMutex.Unlock()
	// members, err := redis.Strings(c.Do("ZREVRANGE", redisZaddName, start, end, "WITHSCORES"))
	return members, err
}

// 获得 SortSet 中的数量，如果不存在也返回0
func getRSortSetNum() int {
	redisMutex.Lock()
	i, _ := redis.Int(connRedis.Do("ZCARD", redisZaddName))
	redisMutex.Unlock()
	return i
}

// 删除 SortSet 中单个元素，成功删除则返回1，失败或者不存在返回0
func deleteRSortSet(key string) int {
	if len(key) == 0 {
		return 0
	}
	redisMutex.Lock()
	a, _ := redis.Int(connRedis.Do("ZREM", redisZaddName, key))
	redisMutex.Unlock()
	return a
}
