package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type ExpenseTracker struct {
	Expenses []Expense `json:"expenses"`
	NextID   int       `json:"next_id"`
}

func main() {
	tracker := LoadExpenses()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Personal Expense Tracker ===")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. View Expenses by Category")
		fmt.Println("4. Show Summary")
		fmt.Println("5. Generate Visualization")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addExpense(&tracker, scanner)
		case "2":
			viewAllExpenses(tracker)
		case "3":
			viewByCategory(tracker, scanner)
		case "4":
			showSummary(tracker)
		case "5":
			generateVisualization(tracker)
		case "6":
			SaveExpenses(tracker)
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addExpense(tracker *ExpenseTracker, scanner *bufio.Scanner) {
	fmt.Print("Enter amount: ")
	scanner.Scan()
	amountStr := scanner.Text()
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount!")
		return
	}

	fmt.Print("Enter category: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Print("Enter description: ")
	scanner.Scan()
	description := scanner.Text()

	expense := Expense{
		ID:          tracker.NextID,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now(),
	}

	tracker.Expenses = append(tracker.Expenses, expense)
	tracker.NextID++
	SaveExpenses(*tracker)
	fmt.Printf("Expense added successfully! (ID: %d)\n", expense.ID)
}

func viewAllExpenses(tracker ExpenseTracker) {
	if len(tracker.Expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Printf("\n%-5s %-10s %-15s %-20s %-12s\n", "ID", "Amount", "Category", "Description", "Date")
	fmt.Println(strings.Repeat("-", 70))
	for _, expense := range tracker.Expenses {
		fmt.Printf("%-5d $%-9.2f %-15s %-20s %-12s\n",
			expense.ID,
			expense.Amount,
			expense.Category,
			expense.Description,
			expense.Date.Format("2006-01-02"))
	}
}

func viewByCategory(tracker ExpenseTracker, scanner *bufio.Scanner) {
	fmt.Print("Enter category to filter: ")
	scanner.Scan()
	category := scanner.Text()

	var filtered []Expense
	total := 0.0
	for _, expense := range tracker.Expenses {
		if strings.EqualFold(expense.Category, category) {
			filtered = append(filtered, expense)
			total += expense.Amount
		}
	}

	if len(filtered) == 0 {
		fmt.Printf("No expenses found in category: %s\n", category)
		return
	}

	fmt.Printf("\nExpenses in category '%s':\n", category)
	fmt.Printf("%-5s %-10s %-20s %-12s\n", "ID", "Amount", "Description", "Date")
	fmt.Println(strings.Repeat("-", 55))
	for _, expense := range filtered {
		fmt.Printf("%-5d $%-9.2f %-20s %-12s\n",
			expense.ID,
			expense.Amount,
			expense.Description,
			expense.Date.Format("2006-01-02"))
	}
	fmt.Printf("\nTotal spent in %s: $%.2f\n", category, total)
}

func showSummary(tracker ExpenseTracker) {
	if len(tracker.Expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	categoryTotals := make(map[string]float64)
	totalSpent := 0.0

	for _, expense := range tracker.Expenses {
		categoryTotals[expense.Category] += expense.Amount
		totalSpent += expense.Amount
	}

	fmt.Printf("\n=== Expense Summary ===\n")
	fmt.Printf("Total Expenses: %d\n", len(tracker.Expenses))
	fmt.Printf("Total Amount: $%.2f\n\n", totalSpent)
	fmt.Println("Spending by Category:")
	fmt.Println(strings.Repeat("-", 30))
	for category, amount := range categoryTotals {
		percentage := (amount / totalSpent) * 100
		fmt.Printf("%-15s: $%-8.2f (%.1f%%)\n", category, amount, percentage)
	}
}

func generateVisualization(tracker ExpenseTracker) {
	if len(tracker.Expenses) == 0 {
		fmt.Println("No expenses to visualize.")
		return
	}

	categoryTotals := make(map[string]float64)
	for _, expense := range tracker.Expenses {
		categoryTotals[expense.Category] += expense.Amount
	}

	fmt.Println("\n=== Spending Visualization ===")
	fmt.Println("(Each '#' represents 5% of total spending)\n")

	totalSpent := 0.0
	for _, amount := range categoryTotals {
		totalSpent += amount
	}

	for category, amount := range categoryTotals {
		percentage := (amount / totalSpent) * 100
		bars := int(percentage / 5)
		if bars == 0 && percentage > 0 {
			bars = 1
		}
		fmt.Printf("%-15s: %s (%.1f%%)\n", category, strings.Repeat("#", bars), percentage)
	}
}