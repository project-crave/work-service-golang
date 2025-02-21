package controller

import (
	"crave/hub/cmd/work/cmd/api/domain/service"
)

type Controller struct {
	svc service.IService
}

func NewController(svc service.IService) *Controller {
	return &Controller{svc: svc}
}
