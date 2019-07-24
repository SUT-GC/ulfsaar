package prob

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s;
	}

	board := make([][]string, numRows)
	for i, _ := range board {
		board[i] = make([]string, len(s))
	}

	i, j := 0, 0
	down := true
	for _, c := range s {
		board[i][j] = string(c)

		if i < numRows-1 && down {
			i++
		} else if i > 0 {
			down = false
			j++
			i--
		} else {
			down = true
			i++
		}
	}

	result := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if len(board[i][j]) > 0 {
				result += board[i][j]
			}
		}
	}
	return result
}

func Test6() {
	fmt.Println(convert("AB", 1))
}
