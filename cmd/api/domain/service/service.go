package service

import (
	"crave/hub/cmd/model"
	"crave/hub/cmd/work/cmd/api/infrastructure/cache"
	"crave/hub/cmd/work/cmd/api/infrastructure/repository"
)

type Service struct {
	cache cache.ICache
	repo  repository.IRepository
}

func NewService(cache cache.ICache, repo repository.IRepository) *Service {
	return &Service{cache: cache, repo: repo}
}

func (s *Service) SaveWork(work *model.Work) (*model.Work, error) {
	createdWork, _ := s.repo.Create(work)
	s.cache.Set(createdWork, 0)
	return work, nil
}

func (s *Service) GetWork(key uint16) (*model.Work, error) {
	workCache, exist := s.cache.Get(key)
	if exist {
		return workCache.ToWork(key), nil
	}
	return s.repo.FindById(key)
}

func (s *Service) UpdateStatus(work *model.Work, status model.Status) (*model.Work, error) {
	work.Status = status
	s.repo.Update(work)
	s.cache.Set(work, 0)
	return work, nil
}
