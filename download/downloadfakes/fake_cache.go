// This file was generated by counterfeiter
package downloadfakes

import (
	"sync"

	"github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/download/cache"
)

type FakeCache struct {
	EntryStub        func(Url string) cache.CacheEntry
	entryMutex       sync.RWMutex
	entryArgsForCall []struct {
		Url string
	}
	entryReturns struct {
		result1 cache.CacheEntry
	}
	entryReturnsOnCall map[int]struct {
		result1 cache.CacheEntry
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCache) Entry(Url string) cache.CacheEntry {
	fake.entryMutex.Lock()
	ret, specificReturn := fake.entryReturnsOnCall[len(fake.entryArgsForCall)]
	fake.entryArgsForCall = append(fake.entryArgsForCall, struct {
		Url string
	}{Url})
	fake.recordInvocation("Entry", []interface{}{Url})
	fake.entryMutex.Unlock()
	if fake.EntryStub != nil {
		return fake.EntryStub(Url)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.entryReturns.result1
}

func (fake *FakeCache) EntryCallCount() int {
	fake.entryMutex.RLock()
	defer fake.entryMutex.RUnlock()
	return len(fake.entryArgsForCall)
}

func (fake *FakeCache) EntryArgsForCall(i int) string {
	fake.entryMutex.RLock()
	defer fake.entryMutex.RUnlock()
	return fake.entryArgsForCall[i].Url
}

func (fake *FakeCache) EntryReturns(result1 cache.CacheEntry) {
	fake.EntryStub = nil
	fake.entryReturns = struct {
		result1 cache.CacheEntry
	}{result1}
}

func (fake *FakeCache) EntryReturnsOnCall(i int, result1 cache.CacheEntry) {
	fake.EntryStub = nil
	if fake.entryReturnsOnCall == nil {
		fake.entryReturnsOnCall = make(map[int]struct {
			result1 cache.CacheEntry
		})
	}
	fake.entryReturnsOnCall[i] = struct {
		result1 cache.CacheEntry
	}{result1}
}

func (fake *FakeCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.entryMutex.RLock()
	defer fake.entryMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCache) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cache.Cache = new(FakeCache)
