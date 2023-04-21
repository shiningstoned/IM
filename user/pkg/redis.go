package pkg

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RUserManager struct {
	client *redis.Client
}

func NewRUserManager(c *redis.Client) *RUserManager {
	return &RUserManager{
		client: c,
	}
}

func (r *RUserManager) AddFriend(myUuid, friendUuid string) (bool, error) {
	if flag := r.CheckFriend(myUuid, friendUuid); flag {
		return false, nil //已经是好友
	}
	err := r.client.SAdd(context.Background(), myUuid, friendUuid).Err()
	if err != nil {
		return false, err
	}
	err = r.client.SAdd(context.Background(), friendUuid, myUuid).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RUserManager) DelFriend(myUuid, friendUuid string) (bool, error) {
	if flag := r.CheckFriend(myUuid, friendUuid); !flag {
		return false, nil //不是好友
	}
	err := r.client.SRem(context.Background(), myUuid, friendUuid).Err()
	if err != nil {
		return false, err
	}
	err = r.client.SRem(context.Background(), friendUuid, myUuid).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RUserManager) CheckFriend(myUuid, friendUuid string) bool {
	exist, _ := r.client.SIsMember(context.Background(), myUuid, friendUuid).Result()
	if !exist {
		return false //用户不存在
	}
	return true //用户存在
}
