package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"exc9/mapred"
)

func main() {
	// Attempt to open the source text
	file, openErr := os.Open("res/meditations.txt")
	if openErr != nil {
		log.Fatalf("unable to access meditations.txt: %v", openErr)
	}
	defer file.Close()

	// Read all lines into slice
	reader := bufio.NewScanner(file)
	var text []string
	for reader.Scan() {
		line := reader.Text()
		text = append(text, line)
	}
	if scanErr := reader.Err(); scanErr != nil {
		log.Fatalf("error while reading file content: %v", scanErr)
	}

	// Execute MapReduce on the collected text
	var mr mapred.MapReduce
	output := mr.Run(text)

	// Convert map → sortable list
	type pair struct {
		word  string
		count int
	}
	entries := make([]pair, 0, len(output))
	for key, val := range output {
		entries = append(entries, pair{word: key, count: val})
	}

	// Order by count (high → low)
	sort.Slice(entries, func(a, b int) bool {
		return entries[a].count > entries[b].count
	})

	// Determine how many elements to display
	limit := 40
	if len(entries) < limit {
		limit = len(entries)
	}

	fmt.Printf("Most common %d terms in Meditations:\n", limit)
	for i := 0; i < limit; i++ {
		fmt.Printf("%3d) %-15s %d\n", i+1, entries[i].word, entries[i].count)
	}
}
