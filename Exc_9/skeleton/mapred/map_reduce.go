package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct{}

// parseLine extracts lowercase words from the input string.
func (m MapReduce) parseLine(src string) []KeyValue {
	reg := regexp.MustCompile(`[^\p{L}]+`)
	cleaned := reg.ReplaceAllString(src, " ")
	parts := strings.Fields(strings.ToLower(cleaned))

	out := make([]KeyValue, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			out = append(out, KeyValue{Key: p, Value: 1})
		}
	}
	return out
}

// combineValues sums all values for a given key.
func (m MapReduce) combineValues(k string, nums []int) KeyValue {
	total := 0
	for _, x := range nums {
		total += x
	}
	return KeyValue{Key: k, Value: total}
}

// --- REQUIRED BY INTERFACE ---

func (m MapReduce) wordCountMapper(text string) []KeyValue {
	return m.parseLine(text)
}

func (m MapReduce) wordCountReducer(key string, values []int) KeyValue {
	return m.combineValues(key, values)
}

// Run performs map → shuffle → reduce.
func (m MapReduce) Run(data []string) map[string]int {
	emit := make(chan KeyValue)

	// MAPPING
	var wgMap sync.WaitGroup
	for _, line := range data {
		wgMap.Add(1)
		go func(s string) {
			defer wgMap.Done()
			for _, kv := range m.wordCountMapper(s) {
				emit <- kv
			}
		}(line)
	}

	go func() {
		wgMap.Wait()
		close(emit)
	}()

	// SHUFFLING
	groups := make(map[string][]int)
	for kv := range emit {
		groups[kv.Key] = append(groups[kv.Key], kv.Value)
	}

	// REDUCING
	result := make(map[string]int)
	var wgRed sync.WaitGroup
	var lock sync.Mutex

	for key, vals := range groups {
		wgRed.Add(1)
		go func(k string, v []int) {
			defer wgRed.Done()
			kv := m.wordCountReducer(k, v)
			lock.Lock()
			result[kv.Key] = kv.Value
			lock.Unlock()
		}(key, vals)
	}

	wgRed.Wait()
	return result
}
