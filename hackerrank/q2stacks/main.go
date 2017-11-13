package main

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	next *node
	prev *node
}

func (n *node) String() string {
	return fmt.Sprintf("%s\n", strconv.Itoa(n.data))
}

type stack struct {
	top *node
}

func (s *stack) push(data int) {
	n := &node{data: data}
	temp := s.top
	if temp == nil {
		s.top = n
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		n.prev = temp
		temp.next = n
	}
}

func (s *stack) pop() {
	if s.top == nil {
		fmt.Errorf("Nope!\n")
	} else {
		temp := s.top
		prev := temp
		for temp.next != nil {
			prev = temp
			temp = temp.next
		}
		prev.next = nil
	}
}

func (s *stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Printf("%v", temp)
		temp = temp.next
	}
}

type queue struct {
	s1 *stack
	s2 *stack
}

func (q *queue) dequeue() {
	t := q.s1.top
	for t.next != nil {
		t = t.next
	}

	for t != q.s1.top {
		q.s2.push(t.data)
		t = t.prev
	}

	q.s1 = &stack{}

	t = q.s2.top
	for t != nil {
		q.s1.push(t.data)
		t = t.next
	}

	q.s2 = &stack{}
}

func main() {
	q := &queue{s1: &stack{}, s2: &stack{}}

	q.s1.push(92118642) //
	q.dequeue()
	q.s1.push(107858633)
	q.s1.push(110186788)
	q.s1.push(883309178)
	q.s1.push(430939631) //
	q.s1.display()
	q.dequeue()
	fmt.Printf("\n")
	q.s1.display()
	// q.s1.push(739711408)
	// q.s1.push(803703507)
	// q.s1.push(643797161)
	// q.s1.push(538560826)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(595864615)
	// q.s1.push(490282285)
	// q.s1.push(558095366)
	// q.s1.push(893666727)
	// q.s1.push(595679828)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(99908215)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(333969117)
	// q.s1.push(604624143)
	// q.s1.push(88712599)
	// q.s1.push(224459820)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(153072902)
	// fmt.Printf("%d\n", q.s1.top.data)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.dequeue()
	// q.s1.push(156974087)
	// q.dequeue()
	// q.s1.push(387224973)
	// q.s1.push(154628865)
	// q.s1.push(256130200)
	// q.s1.push(704295204)
	// q.dequeue()
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(928499989)
	// q.dequeue()
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.dequeue()
	// q.s1.push(319507446)
	// q.s1.push(323714081)
	// q.s1.push(772087837)
	// q.s1.push(350417458)
	// q.s1.push(193303587)
	// q.s1.push(213700781)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(565379092)
	// q.s1.push(284925173)
	// q.dequeue()
	// q.s1.push(794176274)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(766929345)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.dequeue()
	// q.s1.push(42983700)
	// q.dequeue()
	// q.s1.push(156768862)
	// q.s1.push(205008057)
	// q.s1.push(93223219)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.dequeue()
	// q.s1.push(17818922)
	// q.s1.push(917626659)
	// q.s1.push(829031126)
	// q.s1.push(346173907)
	// q.s1.push(78065001)
	// q.dequeue()
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.dequeue()
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(565004472)
	// q.s1.push(753139390)
	// q.dequeue()
	// q.s1.push(629441479)
	// q.s1.push(935933973)
	// q.s1.push(650043630)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(158726470)
	// q.s1.push(206429620)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(590439448)
	// q.s1.push(139555053)
	// q.s1.push(78707344)
	// q.s1.push(989497950)
	// q.s1.push(880166047)
	// q.dequeue()
	// q.s1.push(137608768)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(861823671)
	// q.s1.push(625166323)
	// q.s1.push(431218849)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.s1.push(570007363)
	// q.dequeue()
	// fmt.Printf("%d\n", q.s1.top.data)
	// fmt.Printf("%d\n", q.s1.top.data)
	// q.dequeue()
	// q.s1.push(978253366)

	// var nums int
	// fmt.Scanf("%d", &nums)

	// for i := 0; i < nums; i++ {
	// 	var x, y int
	// 	fmt.Scanf("%d %d", &x, &y)

	// 	switch x {
	// 	case 1:
	// 		q.s1.push(y)
	// 	case 2:
	// 		q.dequeue()
	// 	case 3:
	// 		fmt.Printf("%d", q.s1.top.data)
	// 	}
	// }

}
