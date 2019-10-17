package main

import "fmt"

type node struct {
	val  int
	name string
	next *node
}

type stack struct {
	nums []int
}

func push(s *stack, n int) {
	s.nums = append(s.nums, n)
}
func pop(s *stack) int {
	i := s.nums[len(s.nums)-1]
	s.nums = s.nums[0 : len(s.nums)-1]
	return i
}
func printList(regs *node) {
	for p := regs; p != nil; p = p.next {
		fmt.Println(p.name+" = ", p.val)
	}
}

func addNodeEnd(node, regs *node) *node {
	if regs == nil {
		return regs
	}
	for p := regs; p != nil; p = p.next {
		if p.next == nil {
			p.next = node
			return regs
		}
	}
	return regs
}

func searchList(s string, regs *node) *node {
	if regs == nil {
		return nil
	}
	for p := regs; p != nil; p = p.next {
		if p.name == s {
			return p
		}
	}
	return nil
}
func createRegs(regs *node, a string) *node {
	var names [13]string
	if a == "r" {
		names = [13]string{"r2", "r3", "r4", "r5", "r6", "r7", "r8", "r9", "r10", "r11", "r12", "r13", "r14"}
	} else {
		names = [13]string{"---", "---", "---", "---", "---", "---", "---", "---", "---", "---", "---", "---", "---"}
	}
	var index int = 0
	r2 := &node{0, names[index], nil}
	index++
	r3 := &node{0, names[index], nil}
	index++
	r4 := &node{0, names[index], nil}
	index++
	r5 := &node{0, names[index], nil}
	index++
	r6 := &node{0, names[index], nil}
	index++
	r7 := &node{0, names[index], nil}
	index++
	r8 := &node{0, names[index], nil}
	index++
	r9 := &node{0, names[index], nil}
	index++
	r10 := &node{0, names[index], nil}
	index++
	r11 := &node{0, names[index], nil}
	index++
	r12 := &node{0, names[index], nil}
	index++
	r13 := &node{0, names[index], nil}
	index++
	r14 := &node{0, names[index], nil}
	index++
	regs = addNodeEnd(r2, regs)
	regs = addNodeEnd(r3, regs)
	regs = addNodeEnd(r4, regs)
	regs = addNodeEnd(r5, regs)
	regs = addNodeEnd(r6, regs)
	regs = addNodeEnd(r7, regs)
	regs = addNodeEnd(r8, regs)
	regs = addNodeEnd(r9, regs)
	regs = addNodeEnd(r10, regs)
	regs = addNodeEnd(r11, regs)
	regs = addNodeEnd(r12, regs)
	regs = addNodeEnd(r13, regs)
	regs = addNodeEnd(r14, regs)
	return regs
}
