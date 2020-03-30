package perfect

import "errors"

// ErrOnlyPositive indicates that the number should be positive
var ErrOnlyPositive error = errors.New("Only positive numbers allowed")

// Classification represents the number's classification
type Classification int

// Different classifications
const (
	_ = iota
	ClassificationPerfect
	ClassificationAbundant
	ClassificationDeficient
)

// Classify checks if a number is perfect, abundant, or deficient.
func Classify(number int64) (classification Classification, err error) {
	if number <= 0 {
		err = ErrOnlyPositive
		return
	}

	aliquotSum := aliquotSum(number)
	switch {
	case aliquotSum == number:
		classification = ClassificationPerfect
	case aliquotSum > number:
		classification = ClassificationAbundant
	case aliquotSum < number:
		classification = ClassificationDeficient
	}

	return
}

func aliquotSum(number int64) int64 {
	sum := int64(0)

	for i := int64(1); i < number/2+1; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	return sum
}
