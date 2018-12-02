package main

import (
	"bufio"
	"fmt"
	"os"
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

func checksum(id string) (twos int, threes int) {
	chars := map[rune]int{}

	for _, c := range id {
		chars[c] += 1
	}

	for _, count := range chars {
		if count == 2 {
			twos = 1
		} else if count == 3 {
			threes = 1
		}
	}

	return twos, threes
}

func checksums(ids []string) int {
	twos := 0
	threes := 0
	for _, id := range ids {
		two, three := checksum(id)
		twos += two
		threes += three
	}
	return twos * threes
}

func main() {
	ids := readBoxIds()
	checksum := checksums(ids)
	fmt.Println(checksum)
}
