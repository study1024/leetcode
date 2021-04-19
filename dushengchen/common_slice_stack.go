package dushengchen


type SliceStack struct {
	top int
	raw []int
}

func NewSliceStack(cap int) *SliceStack {
	if cap <= 0 {
		cap = 10
	}
	return &SliceStack{top: -1, raw: make([]int, 0, cap)}
}

func (s *SliceStack) Push(element int) {
	s.top++
	if s.top < len(s.raw) {
		s.raw[s.top] = element
	} else {
		s.raw = append(s.raw, element)
	}
}

func (s *SliceStack) GetTop() (element int, ok bool){
	if s.top < 0 {
		return 0, false
	}
	return s.raw[s.top], true
}

func (s *SliceStack) Pop() (element int, ok bool){
	if s.top < 0 {
		return 0, false
	}
	element = s.raw[s.top]
	s.top--
	ok = true
	return
}
