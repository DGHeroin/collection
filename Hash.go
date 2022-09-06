package collection

import (
    "fmt"
    "strconv"
    "strings"
    "sync"
)

type (
    H     map[string]interface{}
    SafeH struct {
        mu   sync.RWMutex
        data H
    }
)

func (h H) Float64(key string) float64 {
    v, ok := h[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return f64
}
func (h H) Int64(key string) int64 {
    v, ok := h[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return int64(f64)
}
func (h H) Int(key string) int {
    v, ok := h[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return int(f64)
}
func (h H) Bool(key string) bool {
    v, ok := h[key]
    if !ok {
        return false
    }
    str := fmt.Sprintf("%v", v)
    if str == "true" {
        return true
    } else if str == "false" {
        return false
    }
    return false
}
func (h H) String(key string) string {
    return fmt.Sprintf("%v", h[key])
}
func (h H) Interface(key string) interface{} {
    return h[key]
}
func (h H) StringArray(key string) []string {
    var res []string
    str := fmt.Sprintf("%v", h[key])
    tokens := strings.Split(str, ",")
    for _, token := range tokens {
        res = append(res, token)
    }
    return res
}

func NewSafeH() *SafeH {
    return &SafeH{
        data: H{},
    }
}
func (h *SafeH) Float64(key string) float64 {
    h.mu.Lock()
    defer h.mu.Unlock()
    v, ok := h.data[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return f64
}
func (h *SafeH) Int64(key string) int64 {
    h.mu.Lock()
    defer h.mu.Unlock()

    v, ok := h.data[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return int64(f64)
}
func (h *SafeH) Int(key string) int {
    h.mu.Lock()
    defer h.mu.Unlock()

    v, ok := h.data[key]
    if !ok {
        return 0
    }
    str := fmt.Sprintf("%v", v)
    f64, err := strconv.ParseFloat(str, 10)
    if err != nil {
        return 0
    }
    return int(f64)
}
func (h *SafeH) Bool(key string) bool {
    h.mu.Lock()
    defer h.mu.Unlock()

    v, ok := h.data[key]
    if !ok {
        return false
    }
    str := fmt.Sprintf("%v", v)
    if str == "true" {
        return true
    } else if str == "false" {
        return false
    }
    return false
}
func (h *SafeH) String(key string) string {
    h.mu.Lock()
    defer h.mu.Unlock()

    return fmt.Sprintf("%v", h.data[key])
}
func (h *SafeH) Interface(key string) interface{} {
    h.mu.Lock()
    defer h.mu.Unlock()
    return h.data[key]
}
func (h *SafeH) StringArray(key string) []string {
    h.mu.Lock()
    defer h.mu.Unlock()

    var res []string
    str := fmt.Sprintf("%v", h.data[key])
    tokens := strings.Split(str, ",")
    for _, token := range tokens {
        res = append(res, token)
    }
    return res
}
func (h *SafeH) CloneH() H {
    h.mu.Lock()
    defer h.mu.Unlock()
    var result = H{}
    for k, v := range h.data {
        result[k] = v
    }
    return result
}
