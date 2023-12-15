package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
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

	var wg sync.WaitGroup

	r, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	in := make(chan string)

	go read(r, in)
	counters := make([]map[string]int, runtime.NumCPU())
	for i := 0; i < len(counters); i++ {
		wg.Add(1)
		counters[i] = newCounter()
		go process(in, counters[i], &wg)
	}

	wg.Wait()

	result := merge(counters)
	overall := 0
	for key, value := range result {
		overall += value
		fmt.Printf(key+": %d\n", value)
	}
	result["overall"] = overall
	fmt.Printf("Overall: %d\n", result["overall"])
}

func merge(counters []map[string]int) map[string]int {
	result := newCounter()
	for _, counter := range counters {
		for key, value := range counter {
			result[key] += value
		}
	}
	return result
}

func read(r io.Reader, out chan<- string) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		out <- s.Text()
	}

	close(out)
}

func process(in <-chan string, result map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range in {
		for key := range result {
			result[key] += strings.Count(line, key)
		}
	}
}
