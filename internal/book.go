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

var (
	PaymentItemMap = map[string]string{
		"Date":                  "給料月",
		"Salary":                "月収",
		"IncomeTax":             "所得税",
		"ResidentTax":           "住民税",
		"LifeInsurance":         "生命保険",
		"NursingInsurance":      "介護保険",
		"LongTermCareInsurance": "介護保険",
		"EmploymentInsurance":   "雇用保険",
		"HealthInsurance":       "健康保険",
		"WelfarePension":        "厚生年金",
		"PensionInsurance":      "年金保険",
		"Savings":               "貯金",
		"Securities":            "証券",
		"Ideco":                 "確定拠出年金",
		"Mortgage":              "住宅ローン",
		"AdministrativeFee":     "管理費用",
		"RepairFee":             "積立修繕金",
		"Electricity":           "電気",
		"Gas":                   "ガス",
		"Water":                 "水道",
		"ElectricityUsageInKwh": "電気使用量",
		"GasUsageInM3":          "ガス使用量",
		"WaterUsageInM3":        "水道使用量",
		"Internet":              "インターネット",
		"CableTv":               "CATV",
		"TennisClub":            "テニススクール",
		"Pilates":               "ピラティス",
		"Nhk":                   "NHK",
		"CarParkingLot":         "駐車場",
		"BicycleParkingLot":     "駐輪場",
		"CarManagement":         "自動車維持費",
		"CreditCardVisa":        "CREDITCARD_VISA",
		"CreditCardView":        "CREDITCARD_VIEW",
		"CreditCardMC":          "CREDITCARD_MC",
		"BasicLife":             "基本生活費",
		"Amountleft":            "剰余金",
	}
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
		case PaymentItemMap["Date"]:
			mp.Date = row[1].(string) + "-01"
		case PaymentItemMap["Amountleft"]:
			mp.AmountLeft = ParseJpy(row[4].(string))
		}

		switch item := row[1]; item {
		case PaymentItemMap["Salary"]:
			mp.Salary = ParseJpy(row[4].(string))
		case PaymentItemMap["IncomeTax"]:
			mp.IncomeTax = ParseJpy(row[3].(string))
		case PaymentItemMap["ResidentTax"]:
			mp.ResidentTax = ParseJpy(row[3].(string))
		case PaymentItemMap["LifeInsurance"]:
			mp.LifeInsurance = ParseJpy(row[3].(string))
		case PaymentItemMap["NursingInsurance"]:
			mp.NursingInsurance = ParseJpy(row[3].(string))
		case PaymentItemMap["EmploymentInsurance"]:
			mp.EmploymentInsurance = ParseJpy(row[3].(string))
		case PaymentItemMap["HealthInsurance"]:
			mp.HealthInsurance = ParseJpy(row[3].(string))
		case PaymentItemMap["WelfarePension"]:
			mp.WelfarePension = ParseJpy(row[3].(string))
		case PaymentItemMap["Savings"]:
			mp.Savings = ParseJpy(row[3].(string))
		case PaymentItemMap["Securities"]:
			mp.Securities = ParseJpy(row[3].(string))
		case PaymentItemMap["Ideco"]:
			mp.Ideco = ParseJpy(row[3].(string))
		case PaymentItemMap["Mortgage"]:
			mp.Mortgage = ParseJpy(row[3].(string))
		case PaymentItemMap["AdministrativeFee"]:
			mp.AdministrativeFee = ParseJpy(row[3].(string))
		case PaymentItemMap["RepairFee"]:
			mp.RepairFee = ParseJpy(row[3].(string))
		case PaymentItemMap["Electricity"]:
			mp.Electricity = ParseJpy(row[3].(string))
		case PaymentItemMap["Gas"]:
			mp.Gas = ParseJpy(row[3].(string))
		case PaymentItemMap["Water"]:
			mp.Water = ParseJpy(row[3].(string))
		case PaymentItemMap["ElectricityUsageInKwh"]:
			mp.ElectricityUsageInKwh = ParseJpy(row[3].(string))
		case PaymentItemMap["GasUsageInM3"]:
			mp.GasUsageInM3 = ParseJpy(row[3].(string))
		case PaymentItemMap["WaterUsageInM3"]:
			mp.WaterUsageInM3 = ParseJpy(row[3].(string))
		case PaymentItemMap["Internet"]:
			mp.Internet = ParseJpy(row[3].(string))
		case PaymentItemMap["CableTv"]:
			mp.CableTv = ParseJpy(row[3].(string))
		case PaymentItemMap["Pilates"]:
			mp.Pilates = ParseJpy(row[3].(string))
		case PaymentItemMap["Nhk"]:
			mp.Nhk = ParseJpy(row[3].(string))
		case PaymentItemMap["CarParkingLot"]:
			mp.CarParkingLot = ParseJpy(row[3].(string))
		case PaymentItemMap["BicycleParkingLot"]:
			mp.BicycleParkingLot = ParseJpy(row[3].(string))
		case PaymentItemMap["CarManagement"]:
			mp.CarManagement = ParseJpy(row[3].(string))
		case PaymentItemMap["CreditCardVisa"]:
			mp.CreditCardVisa = ParseJpy(row[3].(string))
		case PaymentItemMap["CreditCardView"]:
			mp.CreditCardView = ParseJpy(row[3].(string))
		case PaymentItemMap["CreditCardMC"]:
			mp.CreditCardMC = ParseJpy(row[3].(string))
		case PaymentItemMap["BasicLife"]:
			mp.BasicLife = ParseJpy(row[3].(string))
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
