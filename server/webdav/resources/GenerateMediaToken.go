package resources

import (
	"sync"

	"github.com/google/uuid"
)

type CacheToken struct {
	MediaCache map[string]string
	cacheMutex sync.RWMutex
}

func GenerateToken(ct *CacheToken, path string) string {
	token := uuid.New().String()

	ct.cacheMutex.Lock()
	ct.MediaCache[token] = path
	ct.cacheMutex.Unlock()

	return token
}
