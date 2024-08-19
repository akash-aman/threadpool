package features

import "fmt"

type PrimeCheckTask struct {
	Number int
	Status bool
}

func (pct *PrimeCheckTask) Process() {
	for i := 2; i < pct.Number; i++ {
		if pct.Number%i == 0 {
			pct.Status = false
			break
		} else if i == pct.Number-1 {
			pct.Status = true
			break
		}
	}
	if pct.Status {
		fmt.Printf("Number %d is prime\n", pct.Number)
	} else {
		fmt.Printf("Number %d is not prime\n", pct.Number)
	}
	pct.Status = false
}
