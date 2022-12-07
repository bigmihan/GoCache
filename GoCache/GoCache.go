package GoCache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	date            map[string]itemCache
	CleanupInterval time.Duration
	mu              sync.RWMutex
}

type itemCache struct {
	value          interface{}
	timeDeleteUnix int64
}

func (c *Cache) Cleanup() {

	for {
		c.cleanupDate()
		<-time.After(c.CleanupInterval)
	}
}

func (c *Cache) cleanupDate() {
	if c.date == nil || len(c.date) == 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	jobs := make(chan string)
	timeNowUnix := time.Now().Unix()
	//results:=make(chan itemCache)
	go c.cleanWorker(jobs, timeNowUnix)

	for key := range c.date {
		jobs <- key

	}
	close(jobs)
}

func (c *Cache) cleanWorker(jobs <-chan string, timeNowUnix int64) {
	for key := range jobs {
		itemCache := c.date[key]
		if itemCache.timeDeleteUnix < timeNowUnix {
			delete(c.date, key) //
		}
	}

}

func NewCache(CleanupInterval time.Duration, startCleanup bool) *Cache {
	c := Cache{date: make(map[string]itemCache),

		CleanupInterval: CleanupInterval,
	}

	if startCleanup {

		go c.Cleanup()
	}

	return &c
}

func (c *Cache) Set(key string, value interface{}, DurationLive time.Duration) {
	c.date[key] = itemCache{
		value:          value,
		timeDeleteUnix: time.Now().Add(DurationLive).Unix(),
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	itemCache, ok := c.date[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key %s not found", key))
	}
	if itemCache.timeDeleteUnix < time.Now().Unix() {
		return nil, errors.New(fmt.Sprintf("key %s is outdated", key))
	}
	return itemCache.value, nil
}

func (c *Cache) Delete(key string) bool {

	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.date[key]

	if ok {
		delete(c.date, key)
	}

	return ok
}

func (c *Cache) CountOfElement() int {
	return len(c.date)
}
