package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type SessionManager struct {
	client *redis.Client
}

type Session struct {
	Username string
	IsValid  bool
	Id       string
}

func NewSessionManager(client *redis.Client) *SessionManager {
	return &SessionManager{
		client: client,
	}
}

func (sm *SessionManager) AddSession(key string, session Session) {
	marshal, _ := json.Marshal(session)
	sm.client.Set(key, marshal, 0)
}

func (sm *SessionManager) GetSession(key string) (*Session, error) {
	sessionMap, err := sm.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	if len(sessionMap) == 0 {
		return nil, fmt.Errorf("Session not found")
	}

	var session Session
	err = json.Unmarshal([]byte(sessionMap), &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (sm *SessionManager) InvalidSession(key string) {
	session, _ := sm.GetSession(key)
	session.IsValid = false
	marshal, _ := json.Marshal(session)
	sm.client.Set(key, marshal, 0)
}

func (sm *SessionManager) RemoveSession(key string) {
	sm.client.Del(key)
}
