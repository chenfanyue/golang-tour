package main

import "fmt"

type set map[string]struct{}

func (s set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}

func (s set) Add(key string) {
	s[key] = struct{}{}
}

func (s set) Remove(key string) error {
	_, ok := s[key]
	if !ok {
		return fmt.Errorf("key %v not exists", key)
	}
	delete(s, key)
	return nil
}

func (s set) Size() int {
	return len(s)
}

func main() {
	s := make(set) // s := set{}
	fmt.Println(s)

	fmt.Println(s.Contains("nothing"))

	s.Add("first")
	s.Add("second")
	s.Add("third")
	fmt.Println(s)
	fmt.Println(s.Size())
	fmt.Println(s.Contains("first"))

	s.Remove("first")
	fmt.Println(s.Contains("first"))
	fmt.Println(s)
}
