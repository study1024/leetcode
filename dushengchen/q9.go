package dushengchen


func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if reverse(x) == x {
        return true
    }
    return false
}
//
////from q7
//func reverse(x int) int {
//   if x == 0 {
//       return 0
//   } else if x < 0 {
//       return -reverse(-x)
//   }
//   res := 0
//   for {
//       res = res * 10+ x % 10
//       x = x / 10
//
//       if x == 0 {
//           break
//       }
//   }
//
//   return res
//}
