package main

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	if a > b {
		return a-b <= eps
	}
	return b-a <= eps
}

func TestSimpleDeposit_BalanceAndInterest(t *testing.T) {
	d := SimpleDeposit{
		Principal: 1000.0,
		Rate:      0.05,
		Years:     2.0,
	}

	wantInterest := 1000.0 * 0.05 * 2.0
	if got := d.Interest(); !almostEqual(got, wantInterest, 1e-9) {
		t.Fatalf("SimpleDeposit.Interest: got %v, want %v", got, wantInterest)
	}

	wantBalance := 1000.0 * (1 + 0.05*2.0)
	if got := d.Balance(); !almostEqual(got, wantBalance, 1e-9) {
		t.Fatalf("SimpleDeposit.Balance: got %v, want %v", got, wantBalance)
	}

	wantMaturity := 1000.0 + wantInterest
	if got := d.MaturityValue(); !almostEqual(got, wantMaturity, 1e-9) {
		t.Fatalf("SimpleDeposit.MaturityValue: got %v, want %v", got, wantMaturity)
	}
}

func TestCompoundDeposit_MaturityValue(t *testing.T) {
	d := CompoundDeposit{
		Principal:          1000.0,
		Rate:               0.05,
		Years:              2.0,
		CompoundingPeriods: 12,
	}

	wantMaturity := 1000.0 * math.Pow(1+0.05/12, 12*2.0)
	if got := d.MaturityValue(); !almostEqual(got, wantMaturity, 1e-6) {
		t.Fatalf("CompoundDeposit.MaturityValue: got %v, want %v", got, wantMaturity)
	}

	wantInterest := wantMaturity - 1000.0
	if got := d.Interest(); !almostEqual(got, wantInterest, 1e-6) {
		t.Fatalf("CompoundDeposit.Interest: got %v, want %v", got, wantInterest)
	}

	if got := d.Balance(); !almostEqual(got, wantMaturity, 1e-6) {
		t.Fatalf("CompoundDeposit.Balance: got %v, want %v", got, wantMaturity)
	}
}

func TestFixedDeposit_Values(t *testing.T) {
	d := FixedDeposit{
		Principal:    1000.0,
		FixedInterest: 150.0,
	}

	if got := d.Balance(); got != 1000.0 {
		t.Fatalf("FixedDeposit.Balance: got %v, want %v", got, 1000.0)
	}

	if got := d.Interest(); got != 150.0 {
		t.Fatalf("FixedDeposit.Interest: got %v, want %v", got, 150.0)
	}

	wantMaturity := 1000.0 + 150.0
	if got := d.MaturityValue(); got != wantMaturity {
		t.Fatalf("FixedDeposit.MaturityValue: got %v, want %v", got, wantMaturity)
	}
}

func TestDepositPortfolio_TotalBalanceAndMostProfitable(t *testing.T) {
	portfolio := NewDepositPortfolio()

	simple := SimpleDeposit{Principal: 1000.0, Rate: 0.05, Years: 1.0}
	compound := CompoundDeposit{Principal: 2000.0, Rate: 0.06, Years: 1.0, CompoundingPeriods: 12}
	fixed := FixedDeposit{Principal: 500.0, FixedInterest: 100.0}

	portfolio.Add(simple)
	portfolio.Add(compound)
	portfolio.Add(fixed)

	total := portfolio.TotalBalance()
	wantTotal := simple.Balance() + compound.Balance() + fixed.Balance()
	if !almostEqual(total, wantTotal, 1e-6) {
		t.Fatalf("DepositPortfolio.TotalBalance: got %v, want %v", total, wantTotal)
	}

	mostProfitable := portfolio.MostProfitable()
	if mostProfitable == nil {
		t.Fatalf("MostProfitable: got nil, want non-nil deposit")
	}

	simpleInterest := simple.Interest()
	compoundInterest := compound.Interest()
	fixedInterest := fixed.Interest()

	maxInterest := math.Max(math.Max(simpleInterest, compoundInterest), fixedInterest)
	if got := mostProfitable.Interest(); !almostEqual(got, maxInterest, 1e-6) {
		t.Fatalf("MostProfitable.Interest: got %v, want %v", got, maxInterest)
	}
}

func TestDepositPortfolio_MostProfitableOnEmpty(t *testing.T) {
	portfolio := NewDepositPortfolio()
	if got := portfolio.MostProfitable(); got != nil {
		t.Fatalf("MostProfitable on empty portfolio: got %v, want nil", got)
	}
}

