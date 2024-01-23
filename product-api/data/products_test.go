package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Chai-Latte",
		Price: 1.99,
		SKU:  "abs-abc-xyz",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
