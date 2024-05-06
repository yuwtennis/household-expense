package bop

import "google.golang.org/api/sheets/v4"

type Income interface {
	GetAmount() int
}

// MonthlySalary represents monthly salary
type MonthlySalary struct {
	MonthDate string
	Amount    int
}

func (m *MonthlySalary) GetAmount() int {
	return m.Amount
}

func NewMonthlySalary(data []*sheets.GridData) *MonthlySalary {
	return &MonthlySalary{}
}
