package main

import "math"

// Deposit описывает банковский депозит с методами для расчёта баланса и процентов.
type Deposit interface {
	// Balance возвращает текущий баланс депозита.
	Balance() float64
	// Interest возвращает сумму начисленных процентов.
	Interest() float64
	// MaturityValue возвращает итоговую сумму к концу срока депозита.
	MaturityValue() float64
}

// SimpleDeposit представляет депозит с простым процентом.
// Проценты начисляются только на начальную сумму вклада.
type SimpleDeposit struct {
	Principal float64 // начальная сумма вклада
	Rate      float64 // годовая процентная ставка (например, 0.05 для 5%)
	Years     float64 // срок депозита в годах
}

// CompoundDeposit представляет депозит со сложным процентом.
// Проценты капитализируются несколько раз в год.
type CompoundDeposit struct {
	Principal          float64 // начальная сумма вклада
	Rate               float64 // годовая процентная ставка
	Years              float64 // срок депозита в годах
	CompoundingPeriods int     // количество периодов капитализации в год (например, 12 для ежемесячной)
}

// FixedDeposit представляет депозит с фиксированной суммой процентов.
type FixedDeposit struct {
	Principal     float64 // начальная сумма вклада
	FixedInterest float64 // фиксированная сумма процентов
}

// DepositPortfolio хранит набор депозитов и позволяет работать с ними через интерфейс Deposit.
type DepositPortfolio struct {
	deposits []Deposit
}

// NewDepositPortfolio создаёт пустой портфель депозитов.
func NewDepositPortfolio() *DepositPortfolio {
	return &DepositPortfolio{}
}

// Balance возвращает текущий баланс депозита с простым процентом.
// Формула: Principal * (1 + Rate * Years)
func (d SimpleDeposit) Balance() float64 {
	// TODO: реализуйте расчёт баланса с простым процентом
	return 0
}

// Interest возвращает сумму начисленных процентов для депозита с простым процентом.
// Формула: Principal * Rate * Years
func (d SimpleDeposit) Interest() float64 {
	// TODO: реализуйте расчёт процентов
	return 0
}

// MaturityValue возвращает итоговую сумму к концу срока для депозита с простым процентом.
// Для простого процента это Principal + Interest.
func (d SimpleDeposit) MaturityValue() float64 {
	// TODO: реализуйте расчёт итоговой суммы
	return 0
}

// Balance возвращает текущий баланс депозита со сложным процентом.
// Для сложного процента текущий баланс равен итоговой сумме (MaturityValue).
func (d CompoundDeposit) Balance() float64 {
	// TODO: реализуйте расчёт баланса (используйте MaturityValue)
	return 0
}

// Interest возвращает сумму начисленных процентов для депозита со сложным процентом.
// Формула: MaturityValue - Principal
func (d CompoundDeposit) Interest() float64 {
	// TODO: реализуйте расчёт процентов
	return 0
}

// MaturityValue возвращает итоговую сумму к концу срока для депозита со сложным процентом.
// Формула: Principal * (1 + Rate/CompoundingPeriods)^(CompoundingPeriods * Years)
// Используйте math.Pow для возведения в степень.
func (d CompoundDeposit) MaturityValue() float64 {
	// TODO: реализуйте расчёт итоговой суммы со сложным процентом
	_ = math.Pow
	return 0
}

// Balance возвращает текущий баланс фиксированного депозита.
// Для фиксированного депозита баланс равен начальной сумме.
func (d FixedDeposit) Balance() float64 {
	// TODO: реализуйте возврат баланса
	return 0
}

// Interest возвращает фиксированную сумму процентов.
func (d FixedDeposit) Interest() float64 {
	// TODO: реализуйте возврат фиксированных процентов
	return 0
}

// MaturityValue возвращает итоговую сумму фиксированного депозита.
// Формула: Principal + FixedInterest
func (d FixedDeposit) MaturityValue() float64 {
	// TODO: реализуйте расчёт итоговой суммы
	return 0
}

// Add добавляет депозит в портфель.
func (p *DepositPortfolio) Add(d Deposit) {
	// TODO: реализуйте добавление депозита
}

// TotalBalance возвращает суммарный баланс всех депозитов в портфеле.
func (p *DepositPortfolio) TotalBalance() float64 {
	// TODO: реализуйте расчёт суммарного баланса
	return 0
}

// MostProfitable возвращает депозит с максимальной доходностью (по Interest()).
// Если портфель пуст, должна быть возвращена nil.
// При равенстве доходности допустимо вернуть любой из депозитов с максимальной доходностью.
func (p *DepositPortfolio) MostProfitable() Deposit {
	// TODO: реализуйте поиск наиболее прибыльного депозита
	return nil
}

func main() {}
