package main

// https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1); i > 0; i-- {
		if m <= 0 {
			nums1[i-1] = nums2[n-1]
			n--
			continue
		}
		if n <= 0 {
			nums1[i-1] = nums1[m-1]
			m--
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[i-1] = nums1[m-1]
			m--
		} else {
			nums1[i-1] = nums2[n-1]
			n--
		}
	}
}
