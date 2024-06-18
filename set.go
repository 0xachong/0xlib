package xlib

// 集合
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func NewSets[T comparable](list []T) Set[T] {
	ret := make(Set[T])
	for _, v := range list {
		ret.Add(v)
	}
	return ret
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) Set[T] {
	delete(s, v)
	return s
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	ret := make([]T, 0, len(s))
	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	ret := make(Set[T])
	for k := range s {
		ret.Add(k)
	}
	for k := range other {
		ret.Add(k)
	}
	return ret
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	ret := make(Set[T])
	for k := range s {
		if other.Contains(k) {
			ret.Add(k)
		}
	}
	return ret
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	ret := make(Set[T])
	for k := range s {
		if !other.Contains(k) {
			ret.Add(k)
		}
	}
	return ret
}

func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	ret := make(Set[T])
	for k := range s {
		if !other.Contains(k) {
			ret.Add(k)
		}
	}
	for k := range other {
		if !s.Contains(k) {
			ret.Add(k)
		}
	}
	return ret
}
