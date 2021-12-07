package main

func sortedSquares(nums []int) []int {
	var result []int
	j := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		squaredI := nums[i] * nums[i]
		squaredJ := nums[j] * nums[j]
		if squaredI > squaredJ {
			result = append(result, squaredI)
		} else {
			result = append(result, squaredJ)
			i--
			j--
		}
		if j == i {
			break
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
