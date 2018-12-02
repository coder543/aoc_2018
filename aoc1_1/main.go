package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	freq := int64(0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, e := strconv.ParseInt(scanner.Text(), 10, 64)
		if e != nil {
			panic("bad parse")
		}
		freq += v
	}
	fmt.Println(freq)
}
