package cache

type Cache struct { // кастомный тип
	data map[string]any
}

func New() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (any, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
