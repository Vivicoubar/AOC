package utils

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) AddAll(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for key := range s {
		values = append(values, key)
	}
	return values
}

func (s Set[T]) Equals(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for key := range s {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	if len(s) > len(other) {
		return false
	}
	for key := range s {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSupersetOf(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := NewSet[T]()
	for key := range s {
		union.Add(key)
	}
	for key := range other {
		union.Add(key)
	}
	return union
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := NewSet[T]()
	for key := range s {
		if other.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := NewSet[T]()
	for key := range s {
		if !other.Contains(key) {
			difference.Add(key)
		}
	}
	return difference
}
