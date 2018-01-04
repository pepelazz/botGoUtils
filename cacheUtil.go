package goBotUtils

import "time"

type cacheType struct {
	Data        interface{}
	ExpiredTime time.Time
}

var memCacheMap = map[string]cacheType{}

func MemCacheGet(key string) interface{} {
	if res, ok := memCacheMap[key]; ok {
		// проверяем не истекло ли время кэша
		if res.ExpiredTime.After(time.Now()) {
			return res.Data
		} else {
			delete(memCacheMap, key)
		}
	}
	return nil
}

func MemCachePut(key string, duration int, data interface{}) {
	memCacheMap[key] = cacheType{data, time.Now().Add(time.Duration(duration) * time.Second)}
}
