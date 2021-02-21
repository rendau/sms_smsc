package mem

import (
	"path/filepath"
	"sync"
	"time"
)

type St struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func New() *St {
	return &St{
		data: map[string][]byte{},
	}
}

func (c *St) Get(key string) ([]byte, bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, ok := c.data[key]
	if !ok {
		return nil, false, nil
	}

	return data, true, nil
}

func (c *St) Set(key string, value []byte, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	return nil
}

func (c *St) Del(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)

	return nil
}

func (c *St) Keys(pattern string) []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var ok bool

	resKeys := make([]string, 0, len(c.data))
	for k := range c.data {
		if ok, _ = filepath.Match(pattern, k); ok {
			resKeys = append(resKeys, k)
		}
	}

	return resKeys
}

func (c *St) Clean() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = map[string][]byte{}
}
