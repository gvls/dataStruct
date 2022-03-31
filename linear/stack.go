package linear

// package linear

// Stack
type Stack struct {
	Data   [10]int // Stack
	Top    int     // point to end element + 1
	Bottom int     // point to first element
}

func (s *Stack) IsFull() bool {
	return s.Top-s.Bottom == len(s.Data)
}

func (s *Stack) IsNull() bool {
	return s.Top == s.Bottom
}

func (s *Stack) Add(e int) {
	if s.IsFull() {
		return
	}

	s.Data[s.Top] = e
	s.Top++
}

func (s *Stack) Get() int {
	if s.IsNull() {
		return -1
	}

	s.Top--
	return s.Data[s.Top]
}

func (s *Stack) Length() int {
	return s.Top - s.Bottom
}
