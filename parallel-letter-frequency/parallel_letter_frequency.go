package letter

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
	var freqMapChan = make(chan FreqMap, len(strings))

	for _, s := range strings {
		go func(str string) {
			freqMapChan <- Frequency(str)
		}(s)
	}

	fm := FreqMap{}
	for range strings {
		smap := <-freqMapChan
		for r, c := range smap {
			fm[r] += c
		}
	}

	return fm
}
