package mapred

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("res/meditations.txt")
	if err != nil {
		log.Fatalf("cannot open meditations.txt: %v", err)
	}
	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan error: %v", err)
	}

	var mr MapReduce
	result := mr.Run(lines)

	type kv struct {
		k string
		v int
	}
	var list []kv
	for k, v := range result {
		list = append(list, kv{k, v})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].v > list[j].v
	})

	N := 40
	if len(list) < N {
		N = len(list)
	}

	fmt.Printf("Top %d words:\n", N)
	for i := 0; i < N; i++ {
		fmt.Printf("%3d. %-15s %d\n", i+1, list[i].k, list[i].v)
	}
}
