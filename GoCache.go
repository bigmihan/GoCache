package GoCache

import "errors"

type Cache struct {
	date            map[string]interface{}
	NumberOfElement int
}

func NewCache() *Cache {
	return &Cache{date: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.date[key] = value
	c.NumberOfElement = len(c.date)
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.date[key]
	if !ok {
		return value, errors.New("Не бачу")
	}
	c.NumberOfElement = len(c.date)
	return value, nil
}

func (c *Cache) Delete(key string) {

	delete(c.date, key)
	c.NumberOfElement = len(c.date)
}
