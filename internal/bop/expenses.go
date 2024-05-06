package bop

import "google.golang.org/api/sheets/v4"

// Expense is a struct that represents the expense of balance of payments
type Expense struct {
	MonthDate             string
	IncomeTax             int
	ResidentTax           int
	LifeInsurance         int
	NursingInsurance      int
	EmploymentInsurance   int
	HealthInsurance       int
	WelfarePension        int
	Savings               int
	Securities            int
	Ideco                 int
	Mortgage              int
	AdministrativeFee     int
	RepairFee             int
	Electricity           int
	Gas                   int
	Water                 int
	ElectricityUsageInKwh int
	GasUsageInM3          int
	WaterUsageInM3        int
	internet              int
	CableTv               int
	TennisClub            int
	Pilates               int
	Nhk                   int
	CarParkingLot         int
	BicycleParkingLog     int
	CarManagement         int
	CreditCardVisa        int
	CreditCardView        int
	CreditCardMC          int
	BasicLife             int
}

func NewExpense(data []*sheets.GridData) *Expense {
	// TODO Parse Expenses

	return &Expense{}
}
