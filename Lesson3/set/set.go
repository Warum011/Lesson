package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elems ...T) Set[T] {
	set := make(Set[T])
	for _, e := range elems {
		set.Add(e)
	}
	return set

}

func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

func (s Set[T]) Contains(elem T) bool {
	_, found := s[elem]
	return found
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	res := NewSet[T]()
	for elem := range s {
		res.Add(elem)
	}
	for elem := range s2 {
		res.Add(elem)
	}
	return res
}

func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	res := NewSet[T]()
	for elem := range s {
		if s2.Contains(elem) {
			res.Add(elem)
		}
	}
	return res
}

func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	res := NewSet[T]()
	for elem := range s {
		if !s2.Contains(elem) {
			res.Add(elem)
		}
	}
	return res
}
