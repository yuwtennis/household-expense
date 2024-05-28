package tests

import (
	"github.com/yuwtennis/household-expense/internal"
	"testing"
)

func TestMonAsStr(t *testing.T) {
	expect := "Jan"
	if expect != internal.AsMonStr(1) {
		t.Fatalf("Wrong month name.")
	}
}
