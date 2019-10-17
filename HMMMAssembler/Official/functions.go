package main

import (
	"fmt"
	"strconv"
)

//TODO: check for if searchList returns nil
func setn(rX string, st string, regs *node) {
	v := searchList(rX, regs)
	n, err := strconv.Atoi(st)
	v.val = n
	check(err)
}

func calln(n int, rX string, regs *node) {
	v := searchList(rX, regs)
	v.val = n
}

func read(rX string, regs *node) {
	v := searchList(rX, regs)
	var i int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scanf("%d", &i)
	check(err)
	v.val = i
}

func addn(rX string, st string, regs *node) {
	v := searchList(rX, regs)
	n, err := strconv.Atoi(st)
	v.val = v.val + n
	check(err)
}

func copy(rX string, rY string, regs *node) {
	v := searchList(rX, regs)
	w := searchList(rY, regs)
	v.val = w.val
}

func neg(rX string, rY string, regs *node) {
	v := searchList(rX, regs)
	w := searchList(rY, regs)
	v.val = -w.val
}

func add(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	a.val = b.val + c.val
}

func sub(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	a.val = b.val - c.val
}

func mul(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	a.val = b.val * c.val
}

func mod(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	if c.val == 0 {
		a.val = 0
		fmt.Println("Division by zero error!")
	} else {
		a.val = b.val % c.val
	}
}

func xor(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	a.val = b.val ^ c.val
}

func and(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	a.val = b.val + c.val
}

func or(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	a.val = b.val | c.val
}

func divide(rX string, rY string, rZ string, regs *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(rZ, regs)
	if c.val == 0 {
		a.val = 0
		fmt.Println("Division by zero error!")
	} else {
		a.val = b.val / c.val
	}
}

func write(rX string, regs *node) {
	a := searchList(rX, regs)
	fmt.Println(a.val)
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func clean(s string) string {
	var a string
	if string(s[len(s)-1]) == "\r" || string(s[len(s)-1]) == "\n" || string(s[len(s)-1]) == " " {
		a = s[0 : len(s)-1]
	} else {
		a = s
	}
	return a
}

func loadr(rX string, rY string, regs *node, mem *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	c := searchList(strconv.Itoa(b.val), mem)
	if c == nil {
		a.val = 0
	} else {
		a.val = c.val
	}
}
func loadn(rX string, n string, regs *node, mem *node) {
	a := searchList(rX, regs)
	c := searchList(n, mem)
	if c == nil {
		a.val = 0
	}
	a.val = c.val
}

func storer(rX string, rY string, regs *node, mem *node) {
	a := searchList(rX, regs)
	b := searchList(rY, regs)
	var s string = strconv.Itoa(b.val)
	c := searchList(s, mem)
	if c == nil {
		c = searchList("---", mem)
	}
	c.val = a.val
	c.name = s
}

func storen(rX string, n string, regs *node, mem *node) {
	a := searchList(rX, regs)
	var s string = strconv.Itoa(a.val)
	c := searchList(s, mem)
	if c == nil {
		c = searchList("---", mem)
	}
	c.val = a.val
	c.name = n
}
