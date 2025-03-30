package service

import (
	"crave/hub/cmd/model"
	api "crave/shared/common/api"
)

type IService interface {
	api.IService
	SaveWork(work *model.Work) (*model.Work, error)
	UpdateStatus(work *model.Work, status model.Status) (*model.Work, error)
	GetWork(key uint16) (*model.Work, error)
	PauseWork(workId uint16) error
}
