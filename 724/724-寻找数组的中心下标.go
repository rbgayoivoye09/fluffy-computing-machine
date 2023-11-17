/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心下标
 */
package main

// import "fmt"

// @lc code=start
func pivotIndex(nums []int) int {

	index := -1

	for i := 0; i < len(nums); i++ {
		leftsum := 0
		rightsum := 0

		for l := 0; l < i; l++ {
			leftsum += nums[l]
		}

		// fmt.Printf("leftsum: %v nums %v\n", leftsum, nums[:i])

		for r := i; r < len(nums); r++ {
			rightsum += nums[r]
		}
		rightsum -= nums[i]
		// fmt.Printf("rightsum: %v nums %v\n", rightsum, nums[i:])

		if rightsum == leftsum {
			index = i
			break
		}
	}

	return index
}

// @lc code=end
