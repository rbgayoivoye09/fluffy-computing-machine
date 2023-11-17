/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var nums []int

	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	sort.Ints(nums)

	if len(nums)%2 == 0 {

		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}

// @lc code=end
