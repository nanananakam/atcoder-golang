package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIo() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) PrintLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *Io) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *Io) PrintIntLn(a []int) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func (io *Io) PrintStringLn(a []string) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func Log(name string, value interface{}) {
	fmt.Fprintf(os.Stderr, "%s=%+v\n", name, value)
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pow(a, b int) (result int64) {
	result = 1
	for i := int64(0); i < int64(b); i++ {
		result = (result * int64(a)) % 1000000007
	}
	return
}

var rexp *regexp.Regexp
var memo map[string]int64

func rec(last3 string, n int) int64 {
	if n <= 0 {
		return 1
	}
	key := fmt.Sprintf("%s%d", last3, n)
	memoValue, exits := memo[key]
	if exits {
		return memoValue
	}
	sum := int64(0)
	for _, c := range []string{"A", "C", "G", "T"} {
		last4 := last3 + c
		if !rexp.Match([]byte(last4)) {
			sum += rec(last4[1:4], n-1)
		}
	}
	sum = sum % 1000000007
	memo[key] = sum
	return sum
}

func main() {
	io := NewIo()
	defer io.Flush()
	n := io.NextInt()
	/*
		ダメなパターンは
		AGC
		A?GC
		GAC
		ACG
		AG?C
		が含まれるもの
	*/
	sum := int64(0)
	rexp = regexp.MustCompile("(AGC|GAC|ACG|A.GC|AG.C)")
	memo = map[string]int64{}
	for _, c1 := range []string{"A", "C", "G", "T"} {
		for _, c2 := range []string{"A", "C", "G", "T"} {
			for _, c3 := range []string{"A", "C", "G", "T"} {
				str := c1 + c2 + c3
				if !strings.Contains(str, "AGC") &&
					!strings.Contains(str, "GAC") &&
					!strings.Contains(str, "ACG") {
					sum += rec(str, n-3)
				}
			}
		}
	}
	io.PrintLn(sum % 1000000007)
}
