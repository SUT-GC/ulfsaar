package prob

import "fmt"

type Set struct {
	m map[interface{}]int
}

func (s *Set) contain(i interface{}) bool {
	if _, ok := s.m[i]; ok {
		return true
	}

	return false
}

func (s *Set) add(i interface{}) bool {
	if (*s).m == nil {
		(*s).m = map[interface{}]int{}
	}

	if _, ok := s.m[i]; ok {
		return false
	}

	s.m[i] = 1

	return true
}

func (s *Set) clear() {
	(*s).m = map[interface{}]int{}
}

func (s *Set) size() int {
	return len((*s).m)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength := 0

	for i := 0; i < len(s); i++ {
		set := Set{}
		nowLength := 0
		for j := i; j < len(s); j++ {
			if set.contain(s[j]) {
				break
			}

			set.add(s[j])
			nowLength++

			if nowLength > maxLength {
				maxLength = nowLength
			}
		}
	}

	return maxLength
}

func Test3() {
	c := lengthOfLongestSubstring("avabc")

	fmt.Println(c)
}
