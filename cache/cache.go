package cache

type Cache struct {
	cacheMap map[string]any
}

func New() *Cache {
	var cacheMap map[string]any = make(map[string]any)
	return &Cache{cacheMap}
}

func (c *Cache) Set(key string, value any) {
	c.cacheMap[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.cacheMap, key)
}

func (c *Cache) Get(key string) any {
	return c.cacheMap[key]
}
