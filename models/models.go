package models

import "sync"

type User struct {
	Id      string
	Name    string
	Balance float64
}

type SystemUsers struct {
	mu        sync.Mutex
	userStore map[string]User
}

func (s *SystemUsers) Add(user User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.userStore[user.Id] = user
}

func (s *SystemUsers) Get(userId string) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.userStore[userId]
}

func (s *SystemUsers) Delete(userId string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.userStore, userId)
}

func (s *SystemUsers) UpdateBalance(userId string, newBalance float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user := s.userStore[userId]
	user.Balance = newBalance
	s.userStore[userId] = user
}

var Users SystemUsers
