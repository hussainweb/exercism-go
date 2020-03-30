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
	if number == 1 {
		classification = ClassificationDeficient
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
	sum := int64(1)

	for i := int64(2); i*i <= number; i++ {
		if number%i == 0 {
			sum += i
			if number != i*i {
				// Add the second factor as well, unless it's
				// a proper square, which means the factor is already
				// added.
				sum += number / i
			}
		}
	}

	return sum
}
