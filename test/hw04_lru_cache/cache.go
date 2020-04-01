package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
	Cap() int // this method must be removed
}

type cacheItem struct {
	// Place your code here
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem),
	}
}

func (lc *lruCache) Set(key string, value interface{}) bool {
	k := Key(key)
	if itm, ok := lc.items[k]; ok { // refresh
		itm.Value = value
		lc.queue.MoveToFront(lc.items[k])
		return true
	}
	// insert
	lc.items[k] = lc.queue.PushFront(value)
	if len(lc.items) > lc.capacity { // remove old record
		// !!! remove old key
	}
	return false
}

func (lc *lruCache) Get(key string) (interface{}, bool) {
	k := Key(key)
	if itm, ok := lc.items[k]; ok { // return value
		lc.queue.MoveToFront(lc.items[k])
		return itm.Value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	return
}

func (lc *lruCache) Cap() int {
	return lc.capacity
}
