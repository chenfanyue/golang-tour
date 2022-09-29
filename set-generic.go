package main

import "fmt"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Contains(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s Set[T]) Remove(key T) error {
	_, ok := s[key]
	if !ok {
		return fmt.Errorf("key %#v not exists", key)
	}
	delete(s, key)
	return nil
}

func (s Set[T]) Size() int {
	return len(s)
}

func main() {
	s := make(Set[int])
	fmt.Println(s)

	fmt.Println(s.Contains(0))

	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println(s)
	fmt.Println(s.Size())
	fmt.Println(s.Contains(1))

	s.Remove(1)
	fmt.Println(s.Contains(1))
	fmt.Println(s)
}
