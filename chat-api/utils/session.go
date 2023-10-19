package utils

import (
	"sync"
)

type SessionManager struct {
	sessions *sync.Map
}

type Session struct {
	Username string
	IsValid  bool
	Id       string
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: new(sync.Map),
	}
}

func (sm *SessionManager) GetSessionsMap() *sync.Map {
	return sm.sessions
}

func (sm *SessionManager) AddSession(key string, session Session) {
	sm.sessions.Store(key, session)
}

func (sm *SessionManager) InvalidSession(key string) {
	current, ok := sm.sessions.Load(key)
	if ok {
		session := current.(Session)
		session.IsValid = false
		sm.sessions.Store(key, session)
	}
}

func (sm *SessionManager) RemoveSession(key string) {
	sm.sessions.Delete(key)
}
