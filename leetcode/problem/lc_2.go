package prob

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toStringNum() string {
	if l == nil {
		return ""
	}

	return fmt.Sprintf("%s%d", (*l).Next.toStringNum(), (*l).Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0);
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if (*l1).Next == nil && (*l2).Next == nil {
		sumBit := (*l1).Val + (*l2).Val + carry
		if sumBit > 9 {
			return &ListNode{sumBit % 10, addTwoNumbersWithCarry(&ListNode{0, nil}, &ListNode{0, nil}, 1)}
		}
		return &ListNode{sumBit, nil}
	}

	if (*l1).Next == nil {
		(*l1).Next = &ListNode{0, nil}
	}

	if (*l2).Next == nil {
		(*l2).Next = &ListNode{0, nil}
	}

	sumBit := (*l1).Val + (*l2).Val + carry
	if sumBit > 9 {
		return &ListNode{sumBit % 10, addTwoNumbersWithCarry((*l1).Next, (*l2).Next, 1)}
	} else {
		return &ListNode{sumBit, addTwoNumbersWithCarry((*l1).Next, (*l2).Next, 0)}
	}
}

func Test2() {
	l1 := ListNode{5, nil}
	l2 := ListNode{5, nil}
	fmt.Println((&l1).toStringNum())
	fmt.Println((&l2).toStringNum())

	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(l3.toStringNum())
}
