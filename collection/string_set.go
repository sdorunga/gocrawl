package collection

import (
	"sync"
)

type StringSet struct {
	items map[string]bool
	lock  *sync.Mutex
}

func NewSet() StringSet {
	return StringSet{make(map[string]bool), &sync.Mutex{}}
}

func (set *StringSet) Add(item string) {
	defer set.lock.Unlock()
	set.lock.Lock()
	set.items[item] = true
}

func (set *StringSet) Contains(item string) bool {
	defer set.lock.Unlock()
	set.lock.Lock()
	return set.items[item]
}
