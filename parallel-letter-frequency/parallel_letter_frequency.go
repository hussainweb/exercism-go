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
	var freqMapChan = make(chan FreqMap, len(strings))

	for _, s := range strings {
		wg.Add(1)
		go func(str string) {
			freqMapChan <- Frequency(str)
			defer wg.Done()
		}(s)
	}

	go func() {
		wg.Wait()
		close(freqMapChan)
	}()

	fm := FreqMap{}
	for smap := range freqMapChan {
		// Process each value
		for r, c := range smap {
			fm[r] += c
		}
	}

	return fm
}
