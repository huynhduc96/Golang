package cache

import (
	"social/internal/constant"
	"strconv"

	"github.com/go-redis/redis"
)

type SessionModel struct {
	redis *redis.Client
}

func CreateSessionRepository(cache *redis.Client) *SessionModel {
	return &SessionModel{
		redis: cache,
	}
}

func (session *SessionModel) CreateSession(sessionID string, userId int) error {

	err := session.redis.Set(sessionID, userId, constant.Ttl).Err()

	return err
}

func (session *SessionModel) GetUserIdFromSession(sessionID string) (int, error) {

	value, err := session.redis.Get(sessionID).Result()
	userId, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
