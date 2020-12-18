package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, 1000)  // d[i] min cost before leaving i
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i]
	}
	// We can reach top from either n - 1, or n - 2
	return min(dp[n - 1], dp[n - 2])
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}