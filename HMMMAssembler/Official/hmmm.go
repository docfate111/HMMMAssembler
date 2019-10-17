package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//pointer to 1st register
	regs := &node{0, "r1", nil}
	mem := &node{0, "---", nil}
	Stac := &stack{}
	//create registers 1 through 15
	createRegs(regs, "r")
	createRegs(mem, "m")
	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("HMMM shell")
		fmt.Println("---------------------")
		var emp []string
		for {
			fmt.Print("ASSEMBLE!> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			text = clean(text)
			if strings.Compare("help", text) == 0 {
				fmt.Println("Any instruction from HMMM is valid except jumps")
				fmt.Println("Type quit or exit to leave the shell")
			}
			if strings.Compare("exit", text) == 0 || strings.Compare("quit", text) == 0 {
				break
			}
			if strings.Compare("avengers", text) == 0 {
				fmt.Printf("Easter egg found. Now go to https://www.youtube.com/watch?v=T5Xeat06R28")
				break
			}
			eval(strings.Split(text, " "), regs, mem, Stac, emp, 0)
		}
	} else {
		//read file to array of strings
		file, err := ioutil.ReadFile(os.Args[1])
		check(err)
		var fileName string = string(file)
		var sA []string = strings.Split(fileName, "\n")
		var strArr []string
		//check for last line being counted this causes serious errors
		//in future make sure that the last line has no whitespace before running the program
		if string(sA[len(sA)-1][0]) == " " || string(sA[len(sA)-1][0]) == "\t" || string(sA[len(sA)-1][0]) == "\n" || string(sA[len(sA)-1][0]) == "\r" {
			strArr = remove(sA, sA[len(sA)-1])
		} else {
			strArr = sA
		}
		var x int = 0
		if string(strArr[0][0]) == "1" {
			x++
		}
		for i := x; i < len(strArr); i++ {
			strArr[i] = clean(strArr[i])
		}
		var s int = 0
		if string(strArr[0][0]) == "0" || string(strArr[0][0]) == "1" {
			s++
		}
		for i := x; i < len(strArr); i++ {
			fmt.Println(strings.Split(strArr[i], " "))
			eval(strings.Split(strArr[i], " "), regs, mem, Stac, strArr, s)
		}
		printList(regs)
		printList(mem)
	}
}
func eval(cmd []string, regs *node, mem *node, s *stack, st []string, shift int) int {
	switch string(cmd[(0 + shift)][0]) {
	case "l":
		if string(cmd[0+shift][4]) == "r" {
			loadr(cmd[1+shift], cmd[2+shift], regs, mem)
		} else {
			loadn(cmd[1+shift], cmd[2+shift], regs, mem)
		}
	case "o":
		or(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
	case "x":
		xor(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
	case "s":
		if string(cmd[0+shift][1]) == "u" {
			sub(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
		} else if string(cmd[0+shift][1]) == "e" {
			setn(cmd[1+shift], cmd[2+shift], regs)
		} else if (string(cmd[0+shift][1]) == "t") && (string(cmd[0+shift][5]) == "r") {
			storer(cmd[1+shift], cmd[2+shift], regs, mem)
		} else {
			storen(cmd[1+shift], cmd[2+shift], regs, mem)
		}
	case "a":
		if strings.Compare("and", cmd[0+shift]) == 0 {
			and(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
		} else {
			if strings.Compare("addn", cmd[0+shift]) == 0 {
				addn(cmd[1+shift], cmd[2+shift], regs)
			} else {
				add(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
			}
		}
	case "m":
		if strings.Compare("mov", cmd[0+shift]) == 0 {
			copy(cmd[1+shift], cmd[2+shift], regs)
		} else {
			if strings.Compare("mul", cmd[0+shift]) == 0 {
				mul(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
			} else {
				mod(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
			}
		}
	case "c":
		//calln rX N 	Copy the next address into rX and then jump to mem. addr. N 	call
		if string(cmd[0+shift][1]) == "a" {
			c2, err := strconv.Atoi(cmd[2+shift])
			check(err)
			calln(c2, cmd[1], regs)
			c, err := strconv.Atoi(cmd[2+shift])
			check(err)
			var i int = c
			for i < len(st) {
				eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
				i++
			}
		} else {
			copy(cmd[1+shift], cmd[2+shift], regs)
		}
	case "d":
		divide(cmd[1+shift], cmd[2+shift], cmd[3+shift], regs)
	case "n":
		if strings.Compare("neg", cmd[0+shift]) == 0 {
			neg(cmd[1+shift], cmd[2+shift], regs)
		} //else: nop
	case "e":
		return 2
	case "h":
		return 2
	case "r":
		read(cmd[1+shift], regs)
	case "w":
		write(cmd[1+shift], regs)
	case "p":
		if string(cmd[0+shift][1]) == "u" {
			n := searchList(cmd[1+shift], regs)
			push(s, n.val)
		} else {
			n := searchList(cmd[1+shift], regs)
			n.val = pop(s)
		}
	case "j":
		//jumpn N 	Set program counter to address N 	None
		//jumpr rX 	Set program counter to address in rX 	jump
		if string(cmd[0+shift][1]) == "u" {
			if string(cmd[0+shift][4]) == "n" {
				c, err := strconv.Atoi(cmd[1+shift])
				check(err)
				var i int = c
				for i < len(st) {
					eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
					i++
				}
			} else {
				c := searchList(cmd[1+shift], regs)
				var i int = c.val
				for i < len(st) {
					eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
					i++
				}
			}
		} else {
			/*
				jeqzn rX N 	If rX == 0, then jump to line N 	jeqz
				jnezn rX N 	If rX != 0, then jump to line N 	jnez
				jgtzn rX N 	If rX > 0, then jump to line N 	jgtz
				jltzn rX N 	If rX < 0, then jump to line N 	jltz
			*/
			c := searchList(cmd[1+shift], regs)
			N, err := strconv.Atoi(cmd[2+shift])
			check(err)
			switch string(cmd[0+shift][1]) {
			case "e":
				if c.val == 0 {
					var i int = N
					for i < len(st) {
						eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
						i++
					}
				}
			case "n":
				if c.val != 0 {
					var i int = N
					for i < len(st) {
						eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
						i++
					}
				}
			case "g":
				if c.val > 0 {
					var i int = N
					for i < len(st) {
						eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
						i++
					}
				}
			case "l":
				if c.val < 0 {
					var i int = N
					for i < len(st) {
						eval(strings.Split(st[i], " "), regs, mem, s, st, shift)
						i++
					}
				}
			}

		}
	default:
		fmt.Println("Error: opcode " + cmd[0+shift] + " not found\n")
	}
	return 0
}
