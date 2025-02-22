package repository

import (
	"crave/hub/cmd/model"
	"crave/shared/common/api"
)

type IRepository interface {
	api.IRepository
	Create(*model.Work) (*model.Work, error)
	GetLastIndex() (uint, error)
}
