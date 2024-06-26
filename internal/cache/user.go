package cache

import (
	"eat_box/internal/model"
	"eat_box/pkg/errcode"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

// UserDetailDuration  设置过期时间
const UserDetailDuration = time.Minute * 5

func GetUserCacheName(id string) string {
	return "user" + id
}
func (cache *Cache) SetOneUser(user model.User) *errcode.Error {

	key := GetUserCacheName(user.ID)
	content, err := json.Marshal(&user)
	if err != nil {
		return errcode.ToJSONError
	}
	_, err = cache.redisdb.Exists(key).Result()
	if err != nil {
		return errcode.ErrRedisGet
	}
	err = cache.redisdb.Set(key, content, UserDetailDuration).Err()
	if err != nil {
		return errcode.ErrRedisSet
	}
	return errcode.Success
}
func (cache *Cache) GetUserFromCache(id string) (model.User, *errcode.Error) {
	key := GetUserCacheName(id)
	val, err := cache.redisdb.Get(key).Result()
	if err == redis.Nil || err != nil {
		return model.User{}, errcode.NotFound
	} else {
		user := model.NewUser()
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			//t.Error(target)
			return model.User{}, errcode.JSONUnmarshalError
		}
		return user, errcode.Success
	}
}
func (cache *Cache) DeleteOneUser(id string) *errcode.Error {
	key := GetUserCacheName(id)
	// 检查键是否存在
	_, err := cache.redisdb.Exists(key).Result()
	if err != nil {
		return errcode.ErrRedisGet
	}

	// 如果键存在，则执行删除操作

	err = cache.redisdb.Del(key).Err()
	if err != nil {
		return errcode.ErrRedisDel
	}
	return errcode.Success
}
