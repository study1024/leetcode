package dushengchen

/*
Submission:
    https://leetcode.com/submissions/detail/481198437/
*/
func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	if len(nums) == 0 {
		return result
	}
	endIndexMap := map[int]int{}
	endIndexMap[0] = -1
	for l := 1; l < len(nums); l++ {
		for j := len(result)-1; j >= 0 && len(result[j]) == l - 1; j-- {
			for i := endIndexMap[j]+1; i < len(nums); i++ {
				x := make([]int, l)
				copy(x, result[j])
				x[l-1] = nums[i]
				result = append(result, x)
				endIndexMap[len(result)-1] = i
			}
		}
	}
	result = append(result, nums)
	return result
}
//
//func main() {
//	fmt.Println(subsets([]int{1,2,3}))
//}