package cache

import (
	"crave/hub/cmd/model"
	"time"
)

type ICache interface {
	Set(value *model.Work, ttl time.Duration)
	Get(key uint16) (model.WorkCache, bool)
	Delete(key uint16) error
}
