// Package powerbill предоставляет функции для расчёта потребления
// электроэнергии, стоимости, применения скидки и формирования отчёта.
package powerbill

import (
	"fmt"
)

// Consumption вычисляет потребление электроэнергии по показаниям счётчика.
// Возвращает ошибку, если текущие показания меньше предыдущих.
func Consumption(prev, curr float64) (float64, error) {
	if prev < 0 || curr < 0 {
		return 0, fmt.Errorf("meter readings cannot be negative")
	}
	if curr < prev {
		return 0, fmt.Errorf("current reading cannot be less than previous")
	}
	return curr - prev, nil
}

// EnergyCost вычисляет стоимость электроэнергии.
// Возвращает ошибку при отрицательных значениях или нулевом тарифе.
func EnergyCost(kwh, tariff float64) (float64, error) {
	if kwh < 0 {
		return 0, fmt.Errorf("kwh cannot be negative")
	}
	if tariff <= 0 {
		return 0, fmt.Errorf("tariff must be positive")
	}
	return kwh * tariff, nil
}

// ApplyDiscount применяет скидку к стоимости через указатель.
// Процент должен быть в диапазоне 0..100.
func ApplyDiscount(cost *float64, percent float64) error {
	if cost == nil {
		return fmt.Errorf("cost pointer is nil")
	}
	if percent < 0 || percent > 100 {
		return fmt.Errorf("invalid discount percent")
	}
	*cost = *cost * (1 - percent/100)
	return nil
}

// FormatEnergyReport формирует строку отчёта по электроэнергии.
func FormatEnergyReport(owner string, kwh, cost float64) (string, error) {
	if owner == "" {
		return "", fmt.Errorf("owner cannot be empty")
	}
	if kwh < 0 || cost < 0 {
		return "", fmt.Errorf("invalid values for report")
	}

	report := fmt.Sprintf(
		"Owner: %-12s | Consumption: %8.2f kWh | Cost: %8.2f",
		owner, kwh, cost,
	)

	return report, nil
}
