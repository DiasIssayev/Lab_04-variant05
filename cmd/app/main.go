package main

import (
	"fmt"
	"os"

	"Lab_04-variant05/pkg/powerbill"

	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
)

func main() {

	// Генерация ID отчёта
	id := uuid.New()
	fmt.Printf("Report ID: %s\n\n", id)

	prev := 1200.0
	curr := 1355.5
	tariff := 0.18
	discount := 10.0

	// Расчёт потребления
	kwh, err := powerbill.Consumption(prev, curr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Расчёт стоимости
	cost, err := powerbill.EnergyCost(kwh, tariff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Применение скидки
	err = powerbill.ApplyDiscount(&cost, discount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Формирование текстового отчёта
	report, err := powerbill.FormatEnergyReport("Dias", kwh, cost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(report)
	fmt.Println()

	// Таблица
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Owner", "Consumption (kWh)", "Cost"})

	table.Append([]string{
		"Dias",
		fmt.Sprintf("%.2f", kwh),
		fmt.Sprintf("%.2f", cost),
	})

	table.Render()
}
