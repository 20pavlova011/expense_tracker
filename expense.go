package main

import (
	"time"
)

type Expense struct {
	ID          int
	Description string
	Amount      float64
	Category    string
	Date        time.Time
}

type ExpenseTracker struct {
	expenses []Expense
	nextID   int
}

func NewExpenseTracker() *ExpenseTracker {
	return &ExpenseTracker{
		expenses: make([]Expense, 0),
		nextID:   1,
	}
}

func (et *ExpenseTracker) AddExpense(description string, amount float64, category string) {
	expense := Expense{
		ID:          et.nextID,
		Description: description,
		Amount:      amount,
		Category:    category,
		Date:        time.Now(),
	}
	et.expenses = append(et.expenses, expense)
	et.nextID++
}

func (et *ExpenseTracker) GetAllExpenses() []Expense {
	return et.expenses
}

func (et *ExpenseTracker) GetCategorySummary() map[string]float64 {
	summary := make(map[string]float64)
	for _, expense := range et.expenses {
		summary[expense.Category] += expense.Amount
	}
	return summary
}

func (et *ExpenseTracker) GetExpensesByCategory(category string) []Expense {
	var result []Expense
	for _, expense := range et.expenses {
		if expense.Category == category {
			result = append(result, expense)
		}
	}
	return result
}

func (et *ExpenseTracker) GetTotalExpenses() float64 {
	total := 0.0
	for _, expense := range et.expenses {
		total += expense.Amount
	}
	return total
}