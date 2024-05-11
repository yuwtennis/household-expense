package internal

type Book struct{}

type MonthlyPayment struct {
	Salary                int
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
	Internet              int
	CableTv               int
	TennisClub            int
	Pilates               int
	Nhk                   int
	CarParkingLot         int
	BicycleParkingLot     int
	CarManagement         int
	CreditCardVisa        int
	CreditCardView        int
	CreditCardMC          int
	BasicLife             int
}

// NewMP is a factory function returns a instance of MonthlyPayment
func NewMP(data [][]interface{}) *MonthlyPayment {
	mp := new(MonthlyPayment)

	// Populate the object
	for _, row := range data {
		if row[1].(string) == "" {
			continue
		}
		switch item := row[1]; item {
		case "月収":
			mp.Salary = row[3].(int)
		case "所得税":
			mp.IncomeTax = row[3].(int)
		case "住民税":
			mp.ResidentTax = row[3].(int)
		case "生命保険":
			mp.LifeInsurance = row[3].(int)
		case "介護保険":
			mp.NursingInsurance = row[3].(int)
		case "雇用保険":
			mp.EmploymentInsurance = row[3].(int)
		case "健康保険":
			mp.HealthInsurance = row[3].(int)
		case "厚生年金":
			mp.WelfarePension = row[3].(int)
		case "貯金":
			mp.Savings = row[3].(int)
		case "証券":
			mp.Securities = row[3].(int)
		case "確定拠出年金":
			mp.Ideco = row[3].(int)
		case "住宅ローン":
			mp.Mortgage = row[3].(int)
		case "管理費用":
			mp.AdministrativeFee = row[3].(int)
		case "積立修繕金":
			mp.RepairFee = row[3].(int)
		case "電気":
			mp.Electricity = row[3].(int)
		case "ガス":
			mp.Gas = row[3].(int)
		case "水道":
			mp.Water = row[3].(int)
		case "電気使用量":
			mp.ElectricityUsageInKwh = row[3].(int)
		case "ガス使用量":
			mp.GasUsageInM3 = row[3].(int)
		case "水道使用量":
			mp.WaterUsageInM3 = row[3].(int)
		case "インターネット":
			mp.Internet = row[3].(int)
		case "CATV":
			mp.CableTv = row[3].(int)
		case "テニススクール":
			mp.CableTv = row[3].(int)
		case "ピラティス":
			mp.Pilates = row[3].(int)
		case "NHK":
			mp.Nhk = row[3].(int)
		case "駐車場":
			mp.CarParkingLot = row[3].(int)
		case "駐輪場":
			mp.BicycleParkingLot = row[3].(int)
		case "自動車維持費":
			mp.CarManagement = row[3].(int)
		case "CREDITCARD_VISA":
			mp.CreditCardVisa = row[3].(int)
		case "CREDITCARD_VIEW":
			mp.CreditCardView = row[3].(int)
		case "CREDITCARD_MC":
			mp.CreditCardMC = row[3].(int)
		case "基本生活費":
			mp.BasicLife = row[3].(int)
		}
	}
	return mp
}
