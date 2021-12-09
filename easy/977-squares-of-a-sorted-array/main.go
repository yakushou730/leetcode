package main

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	for k := j; k >= 0; k-- {
		squaredI := nums[i] * nums[i]
		squaredJ := nums[j] * nums[j]
		if squaredI > squaredJ {
			result[k] = squaredI
			i++
		} else {
			result[k] = squaredJ
			j--
		}
	}

	return result
}

//func sortedSquares(nums []int) []int {
//	var result []int
//	i := 0
//	j := len(nums) - 1
//	for k := 0; k < len(nums); k++ {
//		squaredI := nums[i] * nums[i]
//		squaredJ := nums[j] * nums[j]
//		if squaredI > squaredJ {
//			result = append(result, squaredI)
//			i++
//		} else {
//			result = append(result, squaredJ)
//			j--
//		}
//	}
//
//	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
//		result[i], result[j] = result[j], result[i]
//	}
//
//	return result
//}
