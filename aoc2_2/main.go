package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourbasic/radix"
)

func readBoxIds() []string {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

	ids := []string{}
	for scanner.Scan() {
		id := scanner.Text()
		ids = append(ids, id)
	}

	return ids
}

func findDiffing(ids []string) string {

idLoop:
	for i := range ids {
		diffpos := -1
		for c := range ids[i] {
			if ids[i][c] != ids[i+1][c] {
				if diffpos != -1 {
					continue idLoop
				}
				diffpos = c
			}
		}
		if diffpos != -1 {
			return ids[i][:diffpos] + ids[i][diffpos+1:]
		}
	}

	return ""
}

func main() {
	ids := readBoxIds()
	radix.Sort(ids)
	fmt.Println(findDiffing(ids))
}

// quick benchmark main for use with Hyperfine
//
// func main() {
// 	ids := readBoxIds()
// 	for i := 0; i < 25000; i++ {
// 		ids := append([]string{}, ids...)
// 		radix.Sort(ids)
// 		findDiffing(ids)
// 	}
// }
