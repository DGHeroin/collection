package collection

import (
    "sync"
    "time"
)

type Cache struct {
    mutex    sync.RWMutex
    items    map[string]*item
    close    chan struct{}
    onRemove func(key string, value interface{})
}
type item struct {
    key      string
    data     interface{}
    expire   time.Time
    duration time.Duration
}

func NewCache(GCDuration time.Duration) *Cache {
    c := &Cache{
        items:    map[string]*item{},
        close:    make(chan struct{}),
        onRemove: func(key string, value interface{}) {},
    }
    if GCDuration > 0 {
        go func() {
            ticker := time.NewTicker(GCDuration)
            defer ticker.Stop()
            for {
                select {
                case <-ticker.C:
                    c.GC()
                case <-c.close:
                    return
                }
            }
        }()
    }
    return c
}
func (c *Cache) Count() int64 {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    return int64(len(c.items))
}

func (c *Cache) GC() {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    for _, v := range c.items {
        if !v.IsTimeout() {
            continue
        }
        c.onRemove(v.key, v.data)
        delete(c.items, v.key)
    }
}
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    it, ok := c.items[key]
    if !ok {
        return nil, false
    }
    c.addItemDuration(it, it.duration)
    return it.data, true
}
func (c *Cache) addItemDuration(it *item, duration time.Duration) {
    if duration > 0 {
        it.expire = time.Now().Add(duration)
    }
}
func (c *Cache) Set(key string, value interface{}, duration time.Duration) *Cache {
    it := &item{
        key:      key,
        data:     value,
        duration: duration,
    }
    c.addItemDuration(it, duration)
    c.mutex.Lock()
    defer c.mutex.Unlock()
    // overwrite
    if val, ok := c.items[key]; ok {
        c.onRemove(val.key, val.data)
    }
    c.items[key] = it
    return c
}
func (c *Cache) Remove(k string) interface{} {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    v, ok := c.items[k]
    if !ok {
        return nil
    }
    c.onRemove(v.key, v.data)
    delete(c.items, k)
    return v.data
}
func (c *Cache) GetAllItems() []*item {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    var result []*item
    for _, it := range c.items {
        result = append(result, it)
    }
    return result
}
func (c *Cache) Foreach(f func(key string, value interface{}) bool) {
    allItems := c.GetAllItems()
    for _, it := range allItems {
        if !f(it.key, it.data) {
            break
        }
    }
}
func (c *Cache) Close() {
    c.close <- struct{}{}
    c.items = make(map[string]*item)
}
func (it *item) IsTimeout() bool {
    if it.duration == 0 {
        return false
    }
    return time.Now().After(it.expire)
}
func (c *Cache) OnRemove(fn func(key string, value interface{})) {
    c.onRemove = fn
}
