package main

import (
	"fmt"

	"Lab_04-variant05/pkg/powerbill"

	"github.com/google/uuid"
)

func main() {

	// UUID из внешнего пакета GitHub
	id := uuid.New()
	fmt.Printf("Report ID: %s\n\n", id)

	prev := 1200.0
	curr := 1355.5
	tariff := 0.18
	discount := 10.0

	kwh, err := powerbill.Consumption(prev, curr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cost, err := powerbill.EnergyCost(kwh, tariff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = powerbill.ApplyDiscount(&cost, discount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	report, err := powerbill.FormatEnergyReport("Dias", kwh, cost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%s\n", report)
}
