/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心下标
 */
package main

// import "fmt"

// @lc code=start
func pivotIndex2(nums []int) int {

	index := -1

	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i, v := range nums {
		sum -= v
		if leftSum == sum {
			index = i
			break
		}
		leftSum += v
	}

	return index
}

// @lc code=end
