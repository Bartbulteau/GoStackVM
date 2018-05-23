package vm

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"../stack"
)

func Run(p []string) int {
	var s stack.Stack
	stack.Create(&s)
	running := true
	pc := 0

	var ram [2000]int // for variables
	fp := 0           // frame pointer for local scoped variables
	a, b := 0, 0      // value handlers

	p, labels := Labelprocess(p) // Pre-processing labels

	count := len(ram)
	for i := 0; i < count; i++ { // RAM init
		ram[i] = 0
	}

	for i := 0; i < len(p); i++ {
		println(p[i])
	}

	for running {
		switch p[pc] {
		// Arthmetic ops
		case "add_i32":
			stack.Add(&s)
			break
		case "sub_i32":
			stack.Sub(&s)
			break
		case "mul_i32":
			stack.Mul(&s)
			break

		// Stack operations
		case "const_i32": // pushes constant int on stack
			pc++
			val, err := strconv.Atoi(p[pc])
			if err == nil {
				stack.Push(&s, val)
			} else {
				ErrorHandler("not a valid constant integer :", p[pc])
			}
			break
		case "pop":
			stack.Pop(&s)
			break

		// Variables
		case "gstore":
			pc++
			adress := getInt(p[pc])
			val := stack.Pop(&s)
			ram[adress] = val
			break
		case "gload":
			pc++
			adress := getInt(p[pc])
			stack.Push(&s, ram[adress])
			break
		case "store":
			pc++
			adress := getInt(p[pc])
			val := stack.Pop(&s)
			ram[fp+adress] = val
			break
		case "load":
			pc++
			adress := getInt(p[pc])
			val := stack.Pop(&s)
			ram[fp+adress] = val
			break
		// Logic
		case "lt_i32": // int less than
			b = stack.Pop(&s)
			a = stack.Pop(&s)
			if a < b {
				stack.Push(&s, 0)
			} else {
				stack.Push(&s, 1)
			}
			break

		case "eq_i32": // int less than
			b = stack.Pop(&s)
			a = stack.Pop(&s)
			if a == b {
				stack.Push(&s, 0)
			} else {
				stack.Push(&s, 1)
			}
			break

		// Jumps
		case "jmp":
			pc++
			pc = getInt(p[pc])
			break
		case "jz":
			pc++
			adress := getInt(p[pc])
			if stack.Pop(&s) == 0 {
				pc = adress
			}
			break
		case "jnz":
			pc++
			adress := getInt(p[pc])
			if stack.Pop(&s) != 0 {
				pc = adress
			}
			break

		// Functions / Subroutines
		case "call":
			pc++
			adress := getInt(p[pc])
			pc++
			argc := getInt(p[pc])
			stack.Push(&s, argc)
			stack.Push(&s, fp)
			stack.Push(&s, pc)
			fp = s.Sp
			pc = adress
			break
		case "ret":
			rval := stack.Pop(&s)
			s.Sp = fp
			pc = stack.Pop(&s)
			fp = stack.Pop(&s)
			argc := stack.Pop(&s)
			s.Sp -= argc
			stack.Push(&s, rval)
			break

		// I/O
		case "print_i32":
			fmt.Printf("%d", stack.Pop(&s))
			break

		case "halt":
			running = false
			break

		default:
			if val, ok := labels[p[pc][:len(p[pc])-1]]; ok {
				val++
				val--
			} else {
				ErrorHandler("invalid instruction :", p[pc])
			}
			break
		}
		pc++
	}
	return 0
}

func ErrorHandler(msg string, op string) {
	println("GoVM : Fatal Error :", msg, op)
	os.Exit(1)
}

func getInt(s string) int {
	val, err := strconv.Atoi(s)
	if err == nil {
		return val
	} else {
		ErrorHandler("expected int value :", s)
		return 0
	}
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SpeSplit(s string) []string {
	w := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case ' ', '\n', ';', '\t':
			return true
		}
		return false
	})
	return w
}

func Labelprocess(p []string) ([]string, map[string]int) {
	labels := make(map[string]int)
	for i := 0; i < len(p); i++ {
		if p[i][len(p[i])-1] == ':' {
			labels[p[i][:len(p[i])-1]] = i
		}
	}
	for j := 0; j < len(p); j++ {
		if val, ok := labels[p[j]]; ok {
			p[j] = strconv.Itoa(val)
		}
	}
	return p, labels
}
