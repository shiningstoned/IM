package pkg

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	client *redis.Client
}

func NewRedisManager(c *redis.Client) *RedisManager {
	return &RedisManager{
		client: c,
	}
}

func (r *RedisManager) JoinGroup(groupUuid, myUuid string) error {
	err := r.client.SAdd(context.Background(), groupUuid+":members", myUuid).Err()
	if err != nil {
		klog.Fatalf("join group failed: %s", err.Error())
	}
	err = r.client.SAdd(context.Background(), myUuid+":group", groupUuid).Err()
	if err != nil {
		klog.Fatalf("join group failed: %s", err.Error())
	}
	return nil
}

func (r *RedisManager) GetGroupMem(groupUuid string) ([]string, error) {
	members, err := r.client.SMembers(context.Background(), groupUuid+":members").Result()
	if err != nil {
		klog.Fatalf("get group members failed: %s", err.Error())
	}
	return members, nil
}

func (r *RedisManager) IsGroupMem(groupUuid, myUuid string) bool {
	flag, err := r.client.SIsMember(context.Background(), groupUuid, myUuid).Result()
	if err != nil {
		klog.Fatalf("check group member failed: %s", err.Error())
	}
	return flag
}
