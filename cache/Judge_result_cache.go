package cache

import (
	"judger/model"
	"sync"

	"github.com/sta-golang/go-lib-utils/log"
)

type JudgeResultCache struct {
	data  map[string]*model.JudgeStatusData
	mutex sync.RWMutex
}

var judgeResultCache *JudgeResultCache

func GetJudgeResuCache() *JudgeResultCache {
	if judgeResultCache == nil {
		judgeResultCache = &JudgeResultCache{data: make(map[string]*model.JudgeStatusData)}
	}
	return judgeResultCache
}

func (cache *JudgeResultCache) Get(ID string) *model.JudgeStatusData {
	return cache.data[ID]
}

func (cache *JudgeResultCache) Update(ID string, data *model.JudgeStatusData) {
	cache.mutex.RLock()
	cache.data[ID] = data
	cache.mutex.RUnlock()
	log.Infof("JudgeResultCache update k: %s v: %s", ID, data)
}
