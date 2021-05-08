package cache

import (
	"sync"

	"github.com/sta-golang/go-lib-utils/log"
)

type JudgeResultCache struct {
	data  map[string]string
	mutex sync.RWMutex
}

var judgeResultCache *JudgeResultCache

func GetJudgeResuCache() *JudgeResultCache {
	if judgeResultCache == nil {
		judgeResultCache = &JudgeResultCache{data: make(map[string]string)}
	}
	return judgeResultCache
}

func (cache *JudgeResultCache) Get(ID string) string {
	return cache.data[ID]
}

func (cache *JudgeResultCache) Update(ID string, data string) {
	cache.mutex.RLock()
	cache.data[ID] = data
	cache.mutex.RUnlock()
	log.Infof("JudgeResultCache update k: %s v: %s", ID, data)
}
