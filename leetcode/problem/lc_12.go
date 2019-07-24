package prob

import "fmt"

var numPool = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	900:  "CM",
	1000: "M",
	2000: "MM",
	3000: "MMM",
}

func intToRoman(num int) string {
	result := ""

	result += numPool[num/1000*1000]
	num = num % 1000
	result += numPool[num/100*100]
	num = num % 100
	result += numPool[num/10*10]
	num = num % 10
	result += numPool[num]

	return result
}

func Test12() {
	fmt.Println(intToRoman(1994))
}
