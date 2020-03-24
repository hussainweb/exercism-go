package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in parallel.
func ConcurrentFrequency(strings []string) FreqMap {
	var wg sync.WaitGroup
	sc := safeCounter{}
	sc.v = make(FreqMap)

	for _, s := range strings {
		wg.Add(1)
		go func(str string) {
			smap := Frequency(str)
			for r, c := range smap {
				sc.Add(r, c)
			}
			defer wg.Done()
		}(s)
	}

	wg.Wait()

	return sc.v
}

type safeCounter struct {
	v   FreqMap
	mux sync.Mutex
}

// Add increases the value for the counter for the given key.
func (c *safeCounter) Add(key rune, count int) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key] += count
	c.mux.Unlock()
}
