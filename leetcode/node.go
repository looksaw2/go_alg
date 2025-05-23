package leetcode

type Node struct {
	Val  int
	Next *Node
}

func hasCycle(head *Node) bool {
	slow := head
	for fast := head; fast.Next != nil && fast.Next.Next != nil; fast = fast.Next.Next {
		if fast == slow {
			return true
		}
		slow = slow.Next
	}
	return false
}

func detecCycle(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return nil
	}
	slow := head
	for fast := head; fast != nil && fast.Next != nil && fast.Next.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func getLinkNodeLength(head *Node) int {
	i := 0
	for node := head; node != nil; node = node.Next {
		i++
	}
	return i
}

func getIntersectionNode(headA *Node, headB *Node) *Node {
	Alen := getLinkNodeLength(headA)
	Blen := getLinkNodeLength(headB)
	if Alen > Blen {
		for i := Alen - Blen; i > 0; i-- {
			headA = headA.Next
		}
		for node := headB; node != nil; node = node.Next {
			if node == headA {
				return node
			}
			node = node.Next
			headA = headA.Next
		}
	}
	if Alen < Blen {
		for i := Blen - Alen; i > 0; i-- {
			headB = headB.Next
		}
		for node := headA; node != nil; node = node.Next {
			if node == headB {
				return node
			}
			node = node.Next
			headB = headB.Next
		}
	} else {
		for headA != nil && headB != nil {
			if headA == headB {
				return headA
			}
			headA = headA.Next
			headB = headB.Next
		}
	}
	return nil

}
