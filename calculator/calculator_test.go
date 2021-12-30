package calculator_test

import (
	"Hello-World/calculator"
	"testing"
)

func TestCalculator(t *testing.T) {
	ans := calculator.Multyply(2, 2)

	if ans != 4 {
		t.Errorf("Multyply(2, 2) = %d; want 4\n", ans)
	}

	ans = calculator.Multyply(10, 10)

	if ans != 100 {
		t.Errorf("Multyply(10, 10) = %d; want 4\n", ans)
	}
}
