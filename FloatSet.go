package collection

type FloatSet []float64

func (s *FloatSet) add(str float64) FloatSet {
    for _, val := range *s {
        if str == val {
            return *s
        }
    }
    return append(*s, str)
}
func (s *FloatSet) Add(strs ...float64) FloatSet {
    for _, str := range strs {
        *s = s.add(str)
    }
    return *s
}
func (s *FloatSet) Contains(str float64) bool {
    for _, val := range *s {
        if str == val {
            return true
        }
    }
    return false
}
func (s *FloatSet) Remove(str float64) FloatSet {
    var ks FloatSet
    for _, val := range *s {
        if str == val {
            continue
        }
        ks = append(ks, val)
    }
    *s = ks
    return ks
}
func (s *FloatSet) Count() int {
    return len(*s)
}
func (s *FloatSet) Diff(o FloatSet) FloatSet {
    var ks FloatSet
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
func (s *FloatSet) Inter(o FloatSet) FloatSet {
    var ks FloatSet
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
func (s *FloatSet) Merge(o FloatSet) FloatSet {
    var ks FloatSet
    for _, v1 := range *s {
        ks = ks.Add(v1)
    }
    for _, v2 := range o {
        ks = ks.Add(v2)
    }
    return ks
}
func (s *FloatSet) Sum() float64 {
    sum := float64(0)
    for _, v1 := range *s {
        sum += v1
    }
    return sum
}
func (s *FloatSet) Avg() float64 {
    sum := float64(0)
    for _, v1 := range *s {
        sum += v1
    }
    return sum / float64(len(*s))
}
