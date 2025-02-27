package cache

import (
	hubModel "crave/hub/cmd/model"
	"time"
)

type CacheLocal struct {
	data map[uint16]*hubModel.Work
}

func NewCache() *CacheLocal {
	return &CacheLocal{
		data: make(map[uint16]*hubModel.Work),
	}
}

// Delete implements ICache.
func (c *CacheLocal) Delete(key uint16) error {
	delete(c.data, key)
	return nil
}

// Get implements ICache.
func (c *CacheLocal) Get(key uint16) (*hubModel.Work, bool) {
	value, exist := c.data[key]
	if !exist || value == nil {
		return nil, false
	}
	return value, exist
}

// Set implements ICache.
func (c *CacheLocal) Set(work *hubModel.Work, ttl time.Duration) (*hubModel.Work, error) {
	existingWork, ok := c.data[work.Id]
	if ok {
		*existingWork = *work
	}
	if !ok {
		c.data[work.Id] = work
	}
	return c.data[work.Id], nil
}
