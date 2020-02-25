package scale

import (
	"strings"
)

// Scale returns a scale
func Scale(tonic, interval string) []string {
	sharpNotes := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatNotes := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

	// If no interval specified, assume the default
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	// Determine which of the notes we should use depending
	// on the tonic
	var notes []string
	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		notes = sharpNotes
	default:
		notes = flatNotes
	}

	// Find the index of the tonic in our notes
	tonic = strings.Title(tonic)
	foundIndex := -1
	for i := range notes {
		if notes[i] == tonic {
			foundIndex = i
			break
		}
	}

	// If we don't find it, just don't do anything.
	if foundIndex == -1 {
		return nil
	}

	// Begin preparing our scale based on the interval
	var scale []string
	for _, c := range interval {
		scale = append(scale, notes[foundIndex])
		switch c {
		case 'm':
			foundIndex++
		case 'M':
			foundIndex += 2
		case 'A':
			foundIndex += 3
		}
		foundIndex %= len(notes)
	}

	return scale
}
