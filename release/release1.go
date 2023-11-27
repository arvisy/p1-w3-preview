package release

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Release1() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			wg.Add(1)
			go func(w string) {
				defer wg.Done()
				mu.Lock()
				wordCount[w]++
				mu.Unlock()
			}(word)
		}
	}

	wg.Wait()

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
