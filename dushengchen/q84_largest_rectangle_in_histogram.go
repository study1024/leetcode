package dushengchen

/*
Submission:
	https://leetcode.com/submissions/detail/482442786/
*/

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea(heights)
}

func largestRectangleArea(heights []int) int {
	//dp_righ[i]存放i点的右边+1位置
	dp_righ := make([]int, len(heights))
	//dp_left[i]存放i点的左边-1位置
	dp_left := make([]int, len(heights))
	s := NewSliceStack(10)
	for i := 0; i < len(heights); i++{
		for {
			idx, ok := s.Pop()
			if !ok {
				break
			}
			if heights[idx] > heights[i] {
				dp_left[idx] = i
				continue
			}
			s.Push(idx)
			break
		}
		s.Push(i)
	}
	for {
		idx, ok := s.Pop()
		if !ok {
			break
		}
		dp_left[idx] = len(heights)
	}

	for i := len(heights)-1; i >= 0; i--{
		for {
			idx, ok := s.Pop()
			if !ok {
				break
			}
			if heights[idx] > heights[i] {
				dp_righ[idx] = i
				continue
			}
			s.Push(idx)
			break
		}
		s.Push(i)
	}
	for {
		idx, ok := s.Pop()
		if !ok {
			break
		}
		dp_righ[idx] = -1
	}

	max := 0
	for i := range heights {
		a := heights[i] * (dp_left[i] - dp_righ[i] -1)
		if a > max {
			max = a
		}
	}

	return max
}
