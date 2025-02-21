package repository

import "crave/hub/cmd/model"

type IRepository interface {
	Create(*model.Work) (*model.Work, error)
}
