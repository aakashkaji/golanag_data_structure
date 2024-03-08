// Input: nums = [2,7,11,15], target = 9

package main

import "fmt"

func findIndices(nums []int, target int) [][]int {

	twoD := [][]int{}
	hashMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		x := target - nums[i]

		if ind, ok := hashMap[x]; ok {
			twoD = append(twoD, []int{nums[i], nums[ind]})
		}

		hashMap[nums[i]] = i
	}

	return twoD

}

func main() {
	nums, target := []int{4, 2, 7, 11, 15, 5, 8, -2, 1}, 9
	fmt.Println(findIndices(nums, target))
}
