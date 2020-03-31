package main

func findJudge(N int, trust [][]int) int {
	a := make([]int, N+1)
	for _, v := range trust {
		a[v[0]]--
		a[v[1]]++
	}

	for i := 1; i <= N; i++ {
		if a[i] == N-1 {
			return i
		}
	}
	return -1
}
