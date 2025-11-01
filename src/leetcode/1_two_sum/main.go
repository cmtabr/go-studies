package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 3, 11, 7}
	target := 9
	result := twoSum(nums, target)
	println(result[0], result[1]) // Output: 0 3
}
