package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float64 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * (InterestRate(balance) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// We could use Compound Interest formula to get the answer in this way:
	// int(math.Ceil(math.Log(targetBalance/balance) / math.Log(1.0+(InterestRate(balance)/100))))
	//
	// However, this won't account for the fact that the interest rate will change in some cases.
	// So, we have to do this with a loop.
	years := 0
	thisYearBalance := balance
	for thisYearBalance < targetBalance {
		thisYearBalance += Interest(thisYearBalance)
		years++
	}
	return years
}
