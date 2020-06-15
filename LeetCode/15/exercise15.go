package _5

import "sort"

/**
 * CreatedBy: wlei at 2020/6/12
 * Description:
 * 核心思想； 双指针； 确定一个位置后， 剩余两个数字使用双指针移动遍历所有情况；同时先对数字进行排序确保在指针移动时只需要往一个方向移动
 * 基准数字的选取选择情况不同对处理逻辑影响很大：如果以中间位置的数字作为基准，需要处理更多的边界情况
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := new([][]int)

	if len(nums) < 3 {
		return nil
	}

	for index := 0; index < len(nums)-2; index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		frontPos := index + 1
		endPos := len(nums) - 1
		sum := nums[index] + nums[frontPos] + nums[endPos]

		for frontPos < endPos {
			if sum == 0 {
				*results = append(*results, []int{nums[frontPos], nums[index], nums[endPos]})
				frontPos++
				endPos--
				for frontPos < endPos && nums[frontPos] == nums[frontPos-1] {
					frontPos++
				}
				for frontPos < endPos && nums[endPos] == nums[endPos+1] {
					endPos--
				}
			}
			if sum < 0 {
				frontPos++
			}
			if sum > 0 {
				endPos--
			}
			sum = nums[index] + nums[frontPos] + nums[endPos]
		}
	}
	return *results
}
