package prob

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.Trim(s, " "), " ")
	return len(ss[len(ss)-1])
}

func Test58() {
	fmt.Println(lengthOfLastWord("a "))
}
