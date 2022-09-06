package collection

type IntSet []int64

func (s *IntSet) add(str int64) IntSet {
    for _, val := range *s {
        if str == val {
            return *s
        }
    }
    return append(*s, str)
}
func (s *IntSet) Add(strs ...int64) IntSet {
    for _, str := range strs {
        *s = s.add(str)
    }
    return *s
}
func (s *IntSet) Contains(str int64) bool {
    for _, val := range *s {
        if str == val {
            return true
        }
    }
    return false
}
func (s *IntSet) Remove(str int64) IntSet {
    var ks IntSet
    for _, val := range *s {
        if str == val {
            continue
        }
        ks = append(ks, val)
    }
    *s = ks
    return ks
}
func (s *IntSet) Count() int {
    return len(*s)
}
func (s *IntSet) Diff(o IntSet) IntSet {
    var ks IntSet
    for _, v1 := range *s {
        isMatch := false
        for _, v2 := range o {
            if v2 == v1 {
                isMatch = true
                break
            }
        }
        if !isMatch {
            ks = append(ks, v1)
        }
    }
    return ks
}
func (s *IntSet) Inter(o IntSet) IntSet {
    var ks IntSet
    for _, v1 := range *s {
        isMatch := false
        for _, v2 := range o {
            if v2 == v1 {
                isMatch = true
                break
            }
        }
        if isMatch {
            ks = append(ks, v1)
        }
    }
    return ks
}
func (s *IntSet) Merge(o IntSet) IntSet {
    var ks IntSet
    for _, v1 := range *s {
        ks = ks.Add(v1)
    }
    for _, v2 := range o {
        ks = ks.Add(v2)
    }
    return ks
}
func (s *IntSet) Sum() int64 {
    sum := int64(0)
    for _, v1 := range *s {
        sum += v1
    }
    return sum
}
func (s *IntSet) Avg() int64 {
    sum := int64(0)
    for _, v1 := range *s {
        sum += v1
    }
    return sum / int64(len(*s))
}
