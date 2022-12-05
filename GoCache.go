package GoCache

type Cache struct {
	date map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{date: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.date[key] = value
}
