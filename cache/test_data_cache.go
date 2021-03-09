package cache

import (
	"judger/model"
	"sync"
)

type TestDataCache struct {
	data  map[uint][]model.ProblemTestData
	mutex sync.RWMutex
}

var cache *TestDataCache

func GetTestDataCache() *TestDataCache {
	if cache == nil {
		cache = &TestDataCache{data: make(map[uint][]model.ProblemTestData)}
	}
	return cache
}

func (cache *TestDataCache) Get(problemID uint) []model.ProblemTestData {
	return cache.data[problemID]
}

func (cache *TestDataCache) Update(problemID uint, data []model.ProblemTestData) {
	cache.mutex.RLock()
	cache.data[problemID] = data
	cache.mutex.RUnlock()
}
