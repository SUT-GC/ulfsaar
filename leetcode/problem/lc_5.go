package prob

import "fmt"

func longestPalindrome(s string) string {
	sl := len(s)
	if sl <= 1 {
		return s
	}

	status := make([][]int, sl)
	for i, _ := range status {
		status[i] = make([]int, sl)
	}

	for i := 0; i < sl; i++ {
		for j := i; j < sl; j++ {
			fullStatus(&status, s, i, j)
		}
	}

	return findMaxStr(status, s)
}

func findMaxStr(status [][]int, s string) string {
	maxC := 0
	maxS := ""
	for i := 0; i < len(status); i++ {
		for j := i; j < len(status[i]); j++ {
			if status[i][j] == 1 && j-i+1 > maxC {
				maxC = j - i + 1
				maxS = s[i : j+1]
			}
		}
	}

	return maxS
}

func fullStatus(status *[][]int, s string, i int, j int) {
	if (*status)[i][j] > 0 {
		return
	}

	if j-i == 0 {
		(*status)[i][j] = 1
		return
	}

	if j-i == 1 {
		if s[i] == s[j] {
			(*status)[i][j] = 1
		} else {
			(*status)[i][j] = 2
		}

		return
	}

	fullStatus(status, s, i+1, j-1)
	if (*status)[i+1][j-1] == 1 && s[i] == s[j] {
		(*status)[i][j] = 1
	} else {
		(*status)[i][j] = 2
	}

	return
}

func Test5() {
	fmt.Println(longestPalindrome("abbc"))
}
