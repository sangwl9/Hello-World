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
}
