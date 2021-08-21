package atcorder

import "testing"

func TestProduct(t *testing.T) {
	if productLogic(1, 1) != "Odd" {
		t.Errorf("expected: %v, results: %v", "Odd", productLogic(1,1))
	}

	if productLogic(2, 3) != "Even" {
		t.Errorf("expected: %v, results: %v", "Even", productLogic(2,3))
	}
}
