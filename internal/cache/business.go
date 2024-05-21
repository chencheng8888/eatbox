package cache

import (
	"eat_box/internal/model/swagger"
	"eat_box/pkg/errcode"
	"encoding/json"
	"fmt"
	"time"
)

const Expire = time.Hour * 24 * 7

func (cache *Cache) SetBusinessData(data swagger.BusinessData) *errcode.Error {
	key := GetBusinessListKey(data.CurrentPage)
	content, err := json.Marshal(&data)
	if err != nil {
		return errcode.ToJSONError
	}
	//先检验是否存在
	_, err = cache.redisdb.Exists(key).Result()
	if err != nil {
		return errcode.ErrRedisGet
	}
	errset := cache.redisdb.Set(key, content, Expire).Err()
	if errset != nil {
		return errcode.ErrRedisSet
	}
	return errcode.Success
}
func (cache *Cache) GetBusinessData(page int64) (swagger.BusinessData, *errcode.Error) {
	key := GetBusinessListKey(page)
	val, err := cache.redisdb.Get(key).Result()
	if err != nil {
		return swagger.BusinessData{}, errcode.NotFound
	} else {
		data := swagger.BusinessData{}
		if err := json.Unmarshal([]byte(val), &data); err != nil {
			//t.Error(target)
			return swagger.BusinessData{}, errcode.JSONUnmarshalError
		}
		return data, errcode.Success
	}
}
func (cache *Cache) DeleteBusinessData(page int64) *errcode.Error {
	key := GetBusinessListKey(page)
	_, err := cache.redisdb.Exists(key).Result()
	if err != nil {
		return errcode.ErrRedisGet
	}
	err = cache.redisdb.Del(key).Err()
	if err != nil {
		return errcode.ErrRedisDel
	}
	return errcode.Success
}
func GetBusinessListKey(page int64) string {
	key := fmt.Sprintf("BusinessList_page:%d", page)
	return key
}
