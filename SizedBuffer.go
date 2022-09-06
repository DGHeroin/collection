package collection

import (
    "encoding/json"
    "sync"
)

type SizedBuffer struct {
    sync.RWMutex
    list     []interface{}
    capacity int
}

const DefaultCapacity = 8

func NewRingBuffer(cap int) *SizedBuffer {
    ring := &SizedBuffer{}
    if cap < 1 {
        cap = DefaultCapacity
    }
    ring.list = make([]interface{}, 0, cap)
    ring.capacity = cap
    return ring
}
func (r *SizedBuffer) Push(t interface{}) {
    r.Lock()
    defer r.Unlock()
    r.list = append(r.list, t)
}
func (r *SizedBuffer) Pop() interface{} {
    r.Lock()
    defer r.Unlock()
    if len(r.list) == 0 {
        return nil
    }
    p := r.list[0]
    r.list = r.list[1:]
    return p
}
func (r *SizedBuffer) Size() int {
    r.Lock()
    defer r.Unlock()
    return len(r.list)
}
func (r *SizedBuffer) IsEmpty() bool {
    return r.Size() == 0
}
func (r *SizedBuffer) IsFull() bool {
    return r.Size() >= r.capacity
}
func (r *SizedBuffer) ToSlice() (result []interface{}) {
    r.Lock()
    defer r.Unlock()
    for _, v := range r.list {
        result = append(result, v)
    }
    return
}
func (r *SizedBuffer) Foreach(fn func(t interface{}) bool) {
    r.RLock()
    defer r.RUnlock()
    for _, v := range r.list {
        if !fn(v) {
            break
        }
    }
    return
}
func (r *SizedBuffer) ForeachWithRemove(fn func(p interface{}) bool) {
    r.Lock()
    defer r.Unlock()

    i := 0
    for _, x := range r.list {
        if !fn(x) {
            r.list[i] = x
            i++
        }
    }

    for j := i; j < len(r.list); j++ {
        r.list[j] = nil
    }
    r.list = r.list[:i]
    return
}
func (r *SizedBuffer) Clear() {
    r.Lock()
    defer r.Unlock()
    r.list = make([]interface{}, 0, r.capacity)
}
func (r *SizedBuffer) Refill(slice []interface{}) {
    r.Lock()
    defer r.Unlock()

    for _, t := range slice {
        r.list = append(r.list, t)
        if len(r.list) > r.capacity {
            break
        }
    }
}
func (r *SizedBuffer) MarshalJSON() ([]byte, error) {
    var result []interface{}
    r.Foreach(func(t interface{}) bool {
        result = append(result, t)
        return true
    })
    return json.Marshal(result)
}
func (r *SizedBuffer) UnmarshalJSON(b []byte) error {
    var result []interface{}

    err := json.Unmarshal(b, &result)
    if err != nil {
        return err
    }
    r.Refill(result)
    return nil
}
func (r *SizedBuffer) Resize(sz int) {
    r.Lock()
    defer r.Unlock()
    if r.capacity < sz {
        r.capacity = sz
        return
    }
    // not full
    if len(r.list) < sz {
        return
    }
    // truncate
    r.list = r.list[:r.capacity]
}
