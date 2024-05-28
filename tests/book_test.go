package tests

import (
	"github.com/yuwtennis/household-expense/internal"
	"testing"
)

func TestParseJpy(t *testing.T) {
	expect := 25000
	res := internal.ParseJpy("¥25,000")

	if res != expect {
		t.Fatalf("Currency conversion failed!")
	}
}
