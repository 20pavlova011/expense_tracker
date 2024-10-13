package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

var expenses []Expense
var nextID = 1

func main() {
	fmt.Println("=== Personal Expense Tracker ===")
	
	// Load existing data
	loadExpenses()
	
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. View Expenses by Category")
		fmt.Println("4. Show Summary")
		fmt.Println("5. Show Data Visualization")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
		
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch input {
		case "1":
			addExpense()
		case "2":
			viewAllExpenses()
		case "3":
			viewByCategory()
		case "4":
			showSummary()
		case "5":
			showVisualization()
		case "6":
			saveExpenses()
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addExpense() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
	if err != nil {
		fmt.Println("Invalid amount!")
		return
	}
	
	fmt.Print("Enter category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)
	
	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	
	expense := Expense{
		ID:          nextID,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now(),
	}
	
	expenses = append(expenses, expense)
	nextID++
	fmt.Printf("Expense added successfully! (ID: %d)\n", expense.ID)
}

func viewAllExpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}
	
	fmt.Printf("\n%-5s %-10s %-15s %-20s %-12s\n", "ID", "Amount", "Category", "Description", "Date")
	fmt.Println(strings.Repeat("-", 70))
	
	total := 0.0
	for _, expense := range expenses {
		fmt.Printf("%-5d $%-9.2f %-15s %-20s %-12s\n", 
			expense.ID, expense.Amount, expense.Category, 
			expense.Description, expense.Date.Format("2006-01-02"))
		total += expense.Amount
	}
	
	fmt.Printf("\nTotal: $%.2f\n", total)
}

func viewByCategory() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}
	
	categoryMap := make(map[string][]Expense)
	for _, expense := range expenses {
		categoryMap[expense.Category] = append(categoryMap[expense.Category], expense)
	}
	
	fmt.Println("\nExpenses by Category:")
	for category, categoryExpenses := range categoryMap {
		total := 0.0
		for _, expense := range categoryExpenses {
			total += expense.Amount
		}
		fmt.Printf("- %s: $%.2f (%d expenses)\n", category, total, len(categoryExpenses))
	}
}

func showSummary() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}
	
	total := 0.0
	categoryTotals := make(map[string]float64)
	
	for _, expense := range expenses {
		total += expense.Amount
		categoryTotals[expense.Category] += expense.Amount
	}
	
	fmt.Printf("\n=== Expense Summary ===\n")
	fmt.Printf("Total Expenses: $%.2f\n", total)
	fmt.Printf("Number of Expenses: %d\n", len(expenses))
	fmt.Printf("Average per Expense: $%.2f\n", total/float64(len(expenses)))
	
	fmt.Println("\nBy Category:")
	for category, amount := range categoryTotals {
		percentage := (amount / total) * 100
		fmt.Printf("- %s: $%.2f (%.1f%%)\n", category, amount, percentage)
	}
}

func showVisualization() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}
	
	categoryTotals := make(map[string]float64)
	total := 0.0
	
	for _, expense := range expenses {
		categoryTotals[expense.Category] += expense.Amount
		total += expense.Amount
	}
	
	fmt.Println("\n=== Data Visualization ===")
	fmt.Println("Expense Distribution by Category:")
	
	for category, amount := range categoryTotals {
		percentage := (amount / total) * 100
		bars := int(percentage / 2) // Each bar represents 2%
		fmt.Printf("%-15s: %s %.1f%% ($%.2f)\n", 
			category, 
			strings.Repeat("█", bars), 
			percentage, 
			amount)
	}
	
	// Monthly trend visualization
	fmt.Println("\nMonthly Spending Trend:")
	monthlyTotals := make(map[string]float64)
	for _, expense := range expenses {
		monthYear := expense.Date.Format("2006-01")
		monthlyTotals[monthYear] += expense.Amount
	}
	
	for month, amount := range monthlyTotals {
		fmt.Printf("%s: $%.2f\n", month, amount)
	}
}

func saveExpenses() {
	// In a real application, you'd save to a file or database
	// For this demo, we'll just print a message
	fmt.Println("Expenses saved (in memory only for this demo)")
}

func loadExpenses() {
	// In a real application, you'd load from a file or database
	// For this demo, we'll add some sample data
	if len(expenses) == 0 {
		expenses = []Expense{
			{ID: 1, Amount: 45.50, Category: "Food", Description: "Groceries", Date: time.Now().AddDate(0, 0, -5)},
			{ID: 2, Amount: 25.00, Category: "Transport", Description: "Bus pass", Date: time.Now().AddDate(0, 0, -3)},
			{ID: 3, Amount: 80.00, Category: "Entertainment", Description: "Movie tickets", Date: time.Now().AddDate(0, 0, -1)},
		}
		nextID = 4
	}
}