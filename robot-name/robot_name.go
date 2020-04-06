package robotname

import (
	"math/rand"
	"time"
)

// Robot represents a robot name generator
type Robot struct {
	name string
}

var generatedNames []string

var randomGenerator = rand.New(rand.NewSource(time.Now().Unix()))

// Name returns a robot's name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = generateName()
	}
	return r.name, nil
}

// Reset resets the generated names list
func (r *Robot) Reset() {
	r.name = ""
}

func isNameGiven(name string) bool {
	for _, n := range generatedNames {
		if name == n {
			return true
		}
	}
	return false
}

func generateName() string {
	name := ""
	for {
		name = string(randomGenerator.Intn(26)+65) + string(randomGenerator.Intn(26)+65) + string(randomGenerator.Intn(9)+48) + string(randomGenerator.Intn(9)+48) + string(randomGenerator.Intn(9)+48)
		if !isNameGiven(name) {
			break
		}
	}
	generatedNames = append(generatedNames, name)
	return name
}
