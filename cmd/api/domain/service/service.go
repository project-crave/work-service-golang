package service

import (
	"crave/hub/cmd/model"
	"crave/hub/cmd/work/cmd/api/infrastructure/cache"
	"fmt"
	"strings"
)

type Service struct {
	cache cache.ICache
}

func NewService(cache cache.ICache) *Service {
	return &Service{cache: cache}
}

func (s *Service) SaveWork(work *model.Work) (*model.Work, error) {
	createdWork, _ := s.repo.Create(work)
	s.cache.Set(createdWork, 0)
	return work, nil
}

func (s *Service) GetWork(key uint16) (*model.Work, error) {
	workCache, exist := s.cache.Get(key)
	if !exist {
		return nil, fmt.Errorf("👻%s work does not exist", key)
	}
	return workCache.ToWork(key), nil
}
