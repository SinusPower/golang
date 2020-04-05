package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
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

func (lc *lruCache) Set(key Key, value interface{}) bool {
	if itm, ok := lc.items[key]; ok { // refresh
		refreshed := cacheItem{key, value}
		itm.Value = refreshed
		lc.queue.MoveToFront(lc.items[key])
		return true
	}
	// insert
	lc.items[key] = lc.queue.PushFront(cacheItem{key, value})
	if len(lc.items) > lc.capacity { // remove old record
		old := lc.queue.Back()
		delete(lc.items, old.Value.(cacheItem).key)
		lc.queue.Remove(old)
	}
	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if itm, ok := lc.items[key]; ok { // return value
		lc.queue.MoveToFront(lc.items[key])
		return itm.Value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	lc.items = make(map[Key]*listItem)
}

func (lc *lruCache) GetQueue() List {
	return lc.queue
}
