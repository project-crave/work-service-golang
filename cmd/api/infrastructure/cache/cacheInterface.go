package cache

import (
	"crave/hub/cmd/model"
	"time"
)

type ICache interface {
	Set(value *model.Work, ttl time.Duration) (*model.Work, error)
	Get(key uint16) (*model.Work, bool)
	Delete(key uint16) error
}
