// +build cache_debug

package lru

import . "go-cache"
import "time"

func (c *LRUCache) Get(key string) (object CacheObject, err error) {
	if len(key) == 0 {
		return nil, EmptyKey
	}
	start := time.Now()
	cdb, err := c.get(key)
	t := time.Since(start)
	c.AddAccessTime(t.Nanoseconds())
	return c.GetOrFetch(key, cdb, err)
}