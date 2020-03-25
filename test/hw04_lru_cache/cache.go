package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
	Cap() int // this method must be removed
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*cacheItem
}

type cacheItem struct {
	key   Key
	value interface{}
	link  *listItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem),
	}
}

func (lc *lruCache) Set(key string, value interface{}) bool {
	k := Key(key)
	if itm, ok := lc.items[k]; ok { // refresh
		itm.value = value
		lc.queue.MoveToFront(itm.link)
		return true
	}
	// insert
	link := lc.queue.PushFront(value)
	lc.items[k] = &cacheItem{
		key:   k,
		value: value,
		link:  link,
	}
	return false
}

func (lc *lruCache) Get(key string) (interface{}, bool) {
	k := Key(key)
	if itm, ok := lc.items[k]; ok { // return value
		lc.queue.MoveToFront(itm.link)
		return itm.value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	return
}

func (lc *lruCache) Cap() int {
	return lc.capacity
}
