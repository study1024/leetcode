package dushengchen

//func quick_search(nums []int, target int) int {
//    fmt.Println("quick_search", nums, target)
//    start := 0
//    end := len(nums)-1
//    for {
//        mid := (start+end) / 2
//        if nums[mid] == target {
//            return mid
//        }
//        if nums[start] == target {
//            return start
//        }
//        if nums[end] == target {
//            return end
//        }
//        if nums[mid] > target {
//            end = mid-1
//        } else {
//            start = mid+1
//        }
//        if start > end {
//            return -1
//        }
//    }
//}

/**
    二分查找,sortedArray 升序
    如果找到，返回（pos，ok）
    如果没有找到，返回（lookingFor在sortedArray排位，false）
 */
func binarySearch(sortedArray []int, lookingFor int) (int, bool) {
    if len(sortedArray) == 0 {
        return -1, false
    }
    var low int = 0
    var high int = len(sortedArray) - 1
    if sortedArray[low] < lookingFor {
        return low-1, false
    }
    if sortedArray[high] > lookingFor {
        return high+1, false
    }

    for low <= high {
        mid := low + (high - low)/2
        if sortedArray[mid] == lookingFor {
            return mid, true
        } else if sortedArray[mid] > lookingFor {
            high = mid -1
        } else {
            low = mid + 1
        }
    }
    if sortedArray[low] > lookingFor {
        return low-1, false
    }
    return low, false
}

func simpleSearch(arr []int, lookingFor int) int {
    for i, v := range arr {
        if v == lookingFor {
            return i
        }
    }
    return -1
}