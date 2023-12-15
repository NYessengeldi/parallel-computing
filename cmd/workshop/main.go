package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func newCounter() map[string]int {
	return map[string]int{"explain": 0, "mistaken": 0, "ex": 0, "e": 0, "blabla": 0}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()

	r, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	result := newCounter()
	s := bufio.NewScanner(r)
	for s.Scan() {
		for key := range result {
			result[key] += strings.Count(s.Text(), key)
		}
	}
	overall := 0
	for key, value := range result {
		overall += value
		fmt.Printf(key+": %d\n", value)
	}
	result["overall"] = overall
	fmt.Printf("Overall: %d\n", result["overall"])
}
