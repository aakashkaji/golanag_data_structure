package main

import "fmt"

var dp []int

func min(nums ...int) int {

	tmp := nums[0]

	for _, v := range nums {

		if tmp > v {
			tmp = v
		}
	}
	return tmp

}

func solve(indx, n int, days, cost []int) int {

	// base condition

	if indx >= n {
		return 0 // no need to travel more
	}

	if dp[indx] != -1 {
		return dp[indx]
	}

	// one day cost

	oneDay := cost[0] + solve(indx+1, n, days, cost)

	// 7 days cost

	i := indx

	for i < n && days[i] < days[indx]+7 {
		i++
	}

	sevenDays := cost[1] + solve(i, n, days, cost)

	// 30 days cost
	i = indx
	for i < n && days[i] < days[indx]+30 {
		i++
	}

	thirtyDays := cost[2] + solve(i, n, days, cost)

	// calculate minimum

	fmt.Println(oneDay, sevenDays, thirtyDays)

	dp[indx] = min(oneDay, sevenDays, thirtyDays)

	return dp[indx]

}

func main() {

	days := []int{1, 2,4, 6, 7, 8, 20, 35}
	costs := []int{2, 7, 15}

	dp = make([]int, len(days))

	for i := 0; i < len(days); i++ {
		dp[i] = -1
	}

	c := solve(0, len(days), days, costs)
	fmt.Println(c)

	fmt.Println(dp)

}