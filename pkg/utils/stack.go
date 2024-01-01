package utils

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(pin string) {
	*s = append(*s, pin)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Size() int {
	return len(*s)
}

func Reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(Reverse(input[1:]), input[0])
}
