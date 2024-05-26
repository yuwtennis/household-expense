package internal

import (
	"encoding/json"
	"fmt"
	"github.com/yuwtennis/household-expense/internal/helpers"
	"strconv"
	"strings"
)

const (
	PaymentBookSheetRange = "B1:F57"
)

type AccountRecord struct {
	Category string `json:"category"`
	Key      string `json:"key"`
	Value    int    `json:"value"`
}

type Book interface{}

// MonthlyAccount is a value object representing record of single month from Balance of Payment
type MonthlyAccount struct {
	Date                  string
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
	AmountLeft            int
}

// NewMP is a factory function returns a instance of MonthlyAccount
// from unstructured data read from Google Sheet
func NewMP(ud [][]interface{}) *MonthlyAccount {
	mp := new(MonthlyAccount)

	// Populate the object
	for _, row := range ud {
		if len(row) < 1 {
			continue
		}

		switch category := row[0]; category {
		case "給料月":
			mp.Date = row[1].(string) + "-01"
		case "剰余金":
			mp.AmountLeft = ParseJpy(row[4].(string))
		}

		switch item := row[1]; item {
		case "月収":
			mp.Salary = ParseJpy(row[4].(string))
		case "所得税":
			mp.IncomeTax = ParseJpy(row[3].(string))
		case "住民税":
			mp.ResidentTax = ParseJpy(row[3].(string))
		case "生命保険":
			mp.LifeInsurance = ParseJpy(row[3].(string))
		case "介護保険":
			mp.NursingInsurance = ParseJpy(row[3].(string))
		case "雇用保険":
			mp.EmploymentInsurance = ParseJpy(row[3].(string))
		case "健康保険":
			mp.HealthInsurance = ParseJpy(row[3].(string))
		case "厚生年金":
			mp.WelfarePension = ParseJpy(row[3].(string))
		case "貯金":
			mp.Savings = ParseJpy(row[3].(string))
		case "証券":
			mp.Securities = ParseJpy(row[3].(string))
		case "確定拠出年金":
			mp.Ideco = ParseJpy(row[3].(string))
		case "住宅ローン":
			mp.Mortgage = ParseJpy(row[3].(string))
		case "管理費用":
			mp.AdministrativeFee = ParseJpy(row[3].(string))
		case "積立修繕金":
			mp.RepairFee = ParseJpy(row[3].(string))
		case "電気":
			mp.Electricity = ParseJpy(row[3].(string))
		case "ガス":
			mp.Gas = ParseJpy(row[3].(string))
		case "水道":
			mp.Water = ParseJpy(row[3].(string))
		case "電気使用量":
			mp.ElectricityUsageInKwh = ParseJpy(row[3].(string))
		case "ガス使用量":
			mp.GasUsageInM3 = ParseJpy(row[3].(string))
		case "水道使用量":
			mp.WaterUsageInM3 = ParseJpy(row[3].(string))
		case "インターネット":
			mp.Internet = ParseJpy(row[3].(string))
		case "CATV":
			mp.CableTv = ParseJpy(row[3].(string))
		case "テニススクール":
			mp.CableTv = ParseJpy(row[3].(string))
		case "ピラティス":
			mp.Pilates = ParseJpy(row[3].(string))
		case "NHK":
			mp.Nhk = ParseJpy(row[3].(string))
		case "駐車場":
			mp.CarParkingLot = ParseJpy(row[3].(string))
		case "駐輪場":
			mp.BicycleParkingLot = ParseJpy(row[3].(string))
		case "自動車維持費":
			mp.CarManagement = ParseJpy(row[3].(string))
		case "CREDITCARD_VISA":
			mp.CreditCardVisa = ParseJpy(row[3].(string))
		case "CREDITCARD_VIEW":
			mp.CreditCardView = ParseJpy(row[3].(string))
		case "CREDITCARD_MC":
			mp.CreditCardMC = ParseJpy(row[3].(string))
		case "基本生活費":
			mp.BasicLife = ParseJpy(row[3].(string))
		case "剰余金":
			mp.AmountLeft = ParseJpy(row[3].(string))
		}
	}
	return mp
}

func (mp *MonthlyAccount) Serialize() []byte {
	bytes, err := json.Marshal(mp)
	helpers.EvaluateErr(err, "Marshaling account records failed.")

	return bytes
}

func (mp *MonthlyAccount) AsHivePartitionLayout() string {
	date := strings.Split(mp.Date, "-")

	return fmt.Sprintf("year=%s/%s.json",
		date[0],
		date[1],
	)
}

// ParseJpy takes string including jpy currency mark at beginning and return as int
func ParseJpy(s string) int {
	res, _ := strconv.Atoi(
		strings.ReplaceAll(
			strings.Replace(s, "¥", "", 1),
			",", ""))
	return res
}
