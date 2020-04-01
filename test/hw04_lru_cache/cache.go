package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	GetQueue() List // this function must be removed!
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
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
		refreshed := cacheItem{k, value}
		itm.Value = refreshed
		lc.queue.MoveToFront(lc.items[k])
		return true
	}
	// insert
	lc.items[k] = lc.queue.PushFront(cacheItem{k, value})
	if len(lc.items) > lc.capacity { // remove old record
		old := lc.queue.Back()
		delete(lc.items, old.Value.(cacheItem).key)
		lc.queue.Remove(old)
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

func (lc *lruCache) GetQueue() List {
	return lc.queue
}
