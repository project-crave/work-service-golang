package cache

import (
	"crave/hub/cmd/model"
	hubModel "crave/hub/cmd/model"
	"time"
)

type CacheLocal struct {
	data map[uint16]hubModel.WorkCache
}

func NewCache() *CacheLocal {
	return &CacheLocal{
		data: make(map[uint16]hubModel.WorkCache),
	}
}

// Delete implements ICache.
func (c *CacheLocal) Delete(key uint16) error {
	delete(c.data, key)
	return nil
}

// Get implements ICache.
func (c *CacheLocal) Get(key uint16) (model.WorkCache, bool) {
	value, exist := c.data[key]
	return value, exist
}

// Set implements ICache.
func (c *CacheLocal) Set(work *hubModel.Work, ttl time.Duration) {
	c.data[work.Id] = work.ToCache()
}
