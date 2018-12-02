package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scan(freq int64, freqs map[int64]int) int64 {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, e := strconv.ParseInt(scanner.Text(), 10, 64)
		if e != nil {
			panic("bad parse")
		}
		freqs[freq] = 1
		freq += v
		if _, ok := freqs[freq]; ok {
			fmt.Println("found it", freq)
			os.Exit(0)
		}
	}

	return freq
}

func main() {
	freq := int64(0)
	freqs := map[int64]int{}
	for {
		freq = scan(freq, freqs)
	}
}
