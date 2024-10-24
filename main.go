package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	expenseTracker := NewExpenseTracker()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Personal Expense Tracker ===")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. View Category Summary")
		fmt.Println("4. Generate Visualization")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addExpense(expenseTracker, scanner)
		case "2":
			viewAllExpenses(expenseTracker)
		case "3":
			viewCategorySummary(expenseTracker)
		case "4":
			generateVisualization(expenseTracker)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addExpense(et *ExpenseTracker, scanner *bufio.Scanner) {
	fmt.Print("Enter expense description: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Print("Enter amount: ")
	scanner.Scan()
	amountStr := scanner.Text()
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount. Please enter a valid number.")
		return
	}

	fmt.Print("Enter category (Food, Transport, Entertainment, Bills, Other): ")
	scanner.Scan()
	category := scanner.Text()

	et.AddExpense(description, amount, category)
	fmt.Printf("Expense added successfully!\n")
}

func viewAllExpenses(et *ExpenseTracker) {
	expenses := et.GetAllExpenses()
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n=== All Expenses ===")
	total := 0.0
	for i, expense := range expenses {
		fmt.Printf("%d. %s - $%.2f (%s) - %s\n", 
			i+1, expense.Description, expense.Amount, 
			expense.Category, expense.Date.Format("2006-01-02"))
		total += expense.Amount
	}
	fmt.Printf("\nTotal: $%.2f\n", total)
}

func viewCategorySummary(et *ExpenseTracker) {
	summary := et.GetCategorySummary()
	if len(summary) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n=== Category Summary ===")
	total := 0.0
	for category, amount := range summary {
		fmt.Printf("%s: $%.2f\n", category, amount)
		total += amount
	}
	fmt.Printf("Total: $%.2f\n", total)
}

func generateVisualization(et *ExpenseTracker) {
	fmt.Println("\n=== Expense Visualization ===")
	summary := et.GetCategorySummary()
	
	if len(summary) == 0 {
		fmt.Println("No expenses to visualize.")
		return
	}

	// Simple text-based visualization
	total := 0.0
	for _, amount := range summary {
		total += amount
	}

	for category, amount := range summary {
		percentage := (amount / total) * 100
		bars := int(percentage / 5) // Each bar represents 5%
		bar := strings.Repeat("█", bars)
		fmt.Printf("%-15s: %-20s %.1f%% ($%.2f)\n", 
			category, bar, percentage, amount)
	}
}