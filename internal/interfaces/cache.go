package interfaces

import "time"

type Cache interface {
	Set(key string, value []byte, expiration time.Duration) error
	Get(key string) ([]byte, bool, error)
}
