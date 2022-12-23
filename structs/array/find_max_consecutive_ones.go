package array

/*
链接 https://leetcode.cn/problems/max-consecutive-ones/description/
给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

示例1:
输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.

示例2:
输入:nums = [1,0,1,1,0,1]
输出：2
*/

func FindMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, num := range nums {
		if num == 0 {
			count = 0
		} else {
			count++
			if count > max {
				max = count
			}
		}
	}
	return max
}
