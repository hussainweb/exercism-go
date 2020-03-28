package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type runeCount struct {
	r rune
	c int
}

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
	var runeCounterChan = make([]chan runeCount, len(strings))

	for i, s := range strings {
		runeCounterChan[i] = make(chan runeCount, 25)
		go func(str string, runeCounterChan chan runeCount) {
			smap := Frequency(str)
			for r, c := range smap {
				runeCounterChan <- runeCount{r: r, c: c}
			}
			close(runeCounterChan)
		}(s, runeCounterChan[i])
	}

	var wg sync.WaitGroup
	var merged = make(chan runeCount, 100)
	for i := range strings {
		wg.Add(1)
		go func(i int) {
			for rc := range runeCounterChan[i] {
				merged <- rc
			}
			defer wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	fm := FreqMap{}
	for rc := range merged {
		// Process each value
		fm[rc.r] += rc.c
	}

	return fm
}
