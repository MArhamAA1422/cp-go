// I AM A MUSLIM
// URL: https://codeforces.com/contest/2106/problem/C

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	var testCases int = 1
	testCases = ri()
	
	for TC := 1; TC <= testCases; TC++ {
		n := ri()
		k := ri()
		var a = make([]int, n, n)
		var b = make([]int, n, n)
		mn, mx := 1000000000, 0
		for i := 0; i < n; i++ {
			a[i] = ri()
			if mn > a[i] {
				mn = a[i]
			}
			if mx < a[i] {
				mx = a[i]
			}
		}
		prev := -1
		ans := 1
		for i := 0; i < n; i++ {
			b[i] = ri()
			if b[i] != -1 {
				cur := a[i] + b[i]
				if prev == -1 {
					prev = cur
				} else if prev != cur {
					ans = 0
				}
			}
		}
		if prev == -1 {
			fmt.Println(mn+k - mx+1)
		} else {
			for i := 0; i < n; i++ {
				if b[i] == -1 {
					v := prev - a[i]
					if v < 0 || v > k {
						ans = 0
						break
					}
				}
			}
			fmt.Println(ans)
		}
	}
}

// lis05 io
var stdout = bufio.NewWriter(os.Stdout)
var stdin = bufio.NewReader(os.Stdin)
var inputString string
var inputStringPtr int

// prints space-separated args and adds a newline
func println(args ...any) { fmt.Fprintln(stdout, args...) }

// reads a string from stdin
func readString() string {
	str, _ := stdin.ReadString('\n')
	return strings.Trim(str, "\n\r")
}

// reads a string from stdin if necesarry
func readStringIfNecessary() {
	if inputStringPtr == len(inputString) {
		inputString = readString()
		inputStringPtr = 0
	}
}

// skips all spaces in the input until it meets a non-space character
func skipSpaces() {
	readStringIfNecessary()
	for {
		readStringIfNecessary()
		for ; inputStringPtr < len(inputString); inputStringPtr++ {
			if inputString[inputStringPtr] == ' ' {
				continue
			} else {
				break
			}
		}
		if inputStringPtr != len(inputString) {
			break
		}
	}
}

// if n > 0, reads at most n non-space characters.
// if n = 0 reads a full non-space token.
// if n = -1 reads a full string with spaces included.
func nextToken(n int) string {
	if n == -1 {
		for inputStringPtr == len(inputString) {
			readStringIfNecessary()
		}
		start := inputStringPtr
		inputStringPtr = len(inputString)
		return inputString[start:]
	}

	if n == 0 {
		n = math.MaxInt
	}

	skipSpaces()
	start := inputStringPtr
	cnt := 0
	for inputStringPtr < len(inputString) {
		if inputString[inputStringPtr] == ' ' {
			break
		} else {
			cnt++
			inputStringPtr++
			if cnt == n {
				break
			}
		}
	}
	return inputString[start : start+cnt]
}

func readSingle(i any) {
	var token string

	switch v := i.(type) {
	case *int:
		token = nextToken(0)
		*v, _ = strconv.Atoi(token)
	case *int64:
		token = nextToken(0)
		*v, _ = strconv.ParseInt(token, 10, 64)
	case *float64:
		token = nextToken(0)
		*v, _ = strconv.ParseFloat(token, 64)
	case *string:
		token = nextToken(0)
		*v = token
	}
}

// reads a full slice
func readSlice[T any](arg []T) {
	for i := range arg {
		readSingle(&arg[i])
	}
}

// reads many arguments with readSingle()
func readMany(args ...any) {
	for _, val := range args {
		readSingle(val)
	}
}

// reads and returns an int
func ri() int {
	var x int
	readSingle(&x)
	return x
}

// reads and returns an int64
func ri64() int64 {
	var x int64
	readSingle(&x)
	return x
}

// reads and returns a uint
func ru() uint {
	var x uint
	readSingle(&x)
	return x
}

// reads and returns a float64
func rf64() float64 {
	var x float64
	readSingle(&x)
	return x
}

// reads and returns a string of non-space characters
func rstr() string {
	var x string
	readSingle(&x)
	return x
}

// reads and returns a single byte(character) from stdin
func rb() byte {
	return nextToken(1)[0]
}

// reads and returns a full string from stdin
func rfullString() string {
	return nextToken(-1)
}

// reads and returns a slice of length n
func rs[T any](n int) []T {
	res := make([]T, n)
	readSlice(res)
	return res
}
