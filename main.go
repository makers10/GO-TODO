package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var tasks []Task

type Task struct {
	ID        int
	Title     string
	Completed bool
	DueDate   time.Time
	Deleted   bool
	Category  string
}

func getNextID() int {
	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func main() {
	loadTasks()
	// addTask(title, due)
	// saveTasks()
	showDueAlerts()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List All Tasks")
		fmt.Println("3. Show Pending Tasks")
		fmt.Println("4. Show Completed Tasks")
		fmt.Println("5. Show Deleted Tasks")
		fmt.Println("6. Complete Task")
		fmt.Println("7. Delete Task")
		fmt.Println("8. Edit Task")
		fmt.Println("9. Search Task")
		fmt.Println("10. Undo Delete")
		fmt.Println("11. Exit")
		fmt.Println("12. Statistics Dashboard")
		fmt.Println("13. Today's Tasks")
		fmt.Println("14. Show Tasks By Category")
		fmt.Println("15. Export Tasks")

		fmt.Print("choose: ")
		scanner.Scan()
		var choice int
		fmt.Sscanf(scanner.Text(), "%d", &choice)

		switch choice {
		case 1:
			fmt.Print("Task title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Print("Due Date (YYYY-MM-DD): ")
			scanner.Scan()
			due := strings.TrimSpace(scanner.Text())

			fmt.Print("Category: ")
			scanner.Scan()
			category := strings.TrimSpace(scanner.Text())

			addTask(title, due, category)
			saveTasks()

		case 2:
			listTasks()

			// case 2:
			// listTasks()

		case 3:
			listPendingTasks()

		case 4:
			listCompletedTasks()

		case 5:
			listDeletedTasks()

		case 6:
			var id int
			fmt.Print("Task ID: ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &id)

			completeTask(id)
			saveTasks()

		case 7:
			var id int
			fmt.Print("Task ID: ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &id)

			deleteTask(id)
			saveTasks()

		case 8:
			var id int
			fmt.Print("Task ID: ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &id)

			fmt.Print("New task title: ")
			scanner.Scan()
			newTitle := strings.TrimSpace(scanner.Text())

			if editTask(id, newTitle) {
				saveTasks()
				fmt.Println("Task Updated")
			} else {
				fmt.Println("Task not found")
			}

		case 9:
			fmt.Println("Search Keyword: ")
			scanner.Scan()
			keyword := strings.TrimSpace(scanner.Text())
			searchTask(keyword)

		case 10:
			var id int
			fmt.Println("Task ID to restore: ")
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &id)

			if undoDelete(id) {
				saveTasks()
				fmt.Println("Task restored")
			} else {
				fmt.Println("Task not found or not deleted")
			}

		case 12:
			showStatistics()

		case 13:
			listTodayTasks()

		case 14:
	    fmt.Print("Category: ")
	    scanner.Scan()
	    category := strings.TrimSpace(scanner.Text())

	    listTasksByCategory(category)

		case 15:
	    exportTasks()

		case 11:
			fmt.Println("See you Later")
			return

		}
	}
}

func addTask(title string, dueDate string, category string) {
	due, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		fmt.Println("Invalid date format. Use YYYY-MM-DD")
		return
	}

	task := Task{
		ID:        getNextID(),
		Title:     title,
		Completed: false,
		DueDate:   due,
		Category:  category,
	}

	tasks = append(tasks, task)
}
func editTask(id int, newTitle string) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = newTitle
			return true
		}
	}
	return false
}
func listTasks() {
	for _, task := range tasks {

		if task.Deleted {
			continue
		}

		status := " "
		if task.Completed {
			status = "✓"
		}

		fmt.Printf("[%s] %d - %s (Due: %s)\n",
			status,
			task.ID,
			task.Title,
			task.Category,
			task.DueDate.Format("2006-01-02"),
		)
	}
}
func listPendingTasks() {
	for _, task := range tasks {
		if !task.Completed && !task.Deleted {
			fmt.Printf("[%d] %s (Due: %s, Category: %s)\n",
				task.ID,
				task.Title,
				task.DueDate.Format("2006-01-02"),
				task.Category,
			)
		}
	}
}
func listCompletedTasks() {
	for _, task := range tasks {
		if task.Completed && !task.Deleted {
			fmt.Printf("[✓] %d - %s (Due: %s, Category: %s)\n",
				task.ID,
				task.Title,
				task.DueDate.Format("2006-01-02"),
				task.Category,
			)
		}
	}
}
func listDeletedTasks() {
	for _, task := range tasks {
		if task.Deleted {
			fmt.Printf("[X] %d - %s (Due: %s, Category: %s)\n",
				task.ID,
				task.Title,
				task.DueDate.Format("2006-01-02"),
				task.Category,
			)
		}
	}
}
func listTodayTasks() {
	today := time.Now().Format("2006-01-02")
	found := false

	fmt.Println("\n ===== TODAY'S TASKS =======")

	for _, task := range tasks {
		if task.Deleted {
			continue
		}
		if task.DueDate.Format("2006-01-02") == today {
			status := " "
			if task.DueDate.Format("2006-01-02") == today {

				status = "✓"
			}
			fmt.Println("[%s] %d - %s\n",
				status,
				task.ID,
				task.Title,
			)
			found = true
		}
	}
	if !found {
		fmt.Println("Sad!No Tasks Due Today.")
	}
	fmt.Println("==========================\n")
}
func listTasksByCategory(category string) {
	found := false

	fmt.Printf("\n=== %s Tasks ===\n", category)

	for _, task := range tasks {

		if task.Deleted {
			continue
		}

		if strings.EqualFold(task.Category, category) {

			status := " "
			if task.Completed {
				status = "✓"
			}

			fmt.Printf("[%s] %d - %s (Due: %s)\n",
				status,
				task.ID,
				task.Title,
				task.DueDate.Format("2006-01-02"),
			)

			found = true
		}
	}

	if !found {
		fmt.Println("No tasks found.")
	}
}
func saveTasks() {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
func completeTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return true
		}
	}
	return false
}
func deleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			// tasks = append(tasks[:i], tasks[i+1:]...)
			tasks[i].Deleted = true
			return true
		}
	}
	return false
}
func undoDelete(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id && tasks[i].Deleted {
			tasks[i].Deleted = false
			return true
		}
	}
	return false
}

//	func loadTasks(){
//		data, err := os.ReadFile("tasks.json")
//		if err != nil {
//			return
//		}
//		json.Unmarshal(data, &tasks)
//	  }
func loadTasks() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}
}
func searchTask(keyword string) {
	found := false

	for _, task := range tasks {
		if strings.Contains(
			strings.ToLower(task.Title),
			strings.ToLower(keyword),
		) {
			status := " "
			if task.Completed {
				status = "✓"
			}

			fmt.Printf("[%s] %d - %s\n",
				status,
				task.ID,
				task.Title,
			)

			found = true
		}
	}

	if !found {
		fmt.Println("No matching tasks found")
	}
}
func showDueAlerts() {
	now := time.Now()

	for _, task := range tasks {
		if task.Deleted || task.Completed {
			continue
		}

		if task.DueDate.Before(now) {
			fmt.Printf("OVERDUE: [%d] %s (Due: %s)\n",
				task.ID,
				task.Title,
				task.DueDate.Format("2006-01-02"),
			)
		}
	}
}
func showStatistics() {
	total := 0
	completed := 0
	pending := 0
	deleted := 0

	for _, task := range tasks {
		total++
		if task.Deleted {
			deleted++
		} else if task.Completed {
			completed++
		} else {
			pending++
		}
	}
	fmt.Printf("Statistics:\n")
	fmt.Printf("  Total: %d\n", total)
	fmt.Printf("  Completed: %d\n", completed)
	fmt.Printf("  Pending: %d\n", pending)
	fmt.Printf("  Deleted: %d\n", deleted)
}
func exportTasks() {
	var output strings.Builder

	output.WriteString("===== TASK LIST =====\n\n")

	for _, task := range tasks {

		if task.Deleted {
			continue
		}

		status := "Pending"
		if task.Completed {
			status = "Completed"
		}

		output.WriteString(
			fmt.Sprintf(
				"ID: %d\nTitle: %s\nCategory: %s\nStatus: %s\nDue: %s\n\n",
				task.ID,
				task.Title,
				task.Category,
				status,
				task.DueDate.Format("2006-01-02"),
			),
		)
	}

	err := os.WriteFile(
		"tasks_export.txt",
		[]byte(output.String()),
		0644,
	)

	if err != nil {
		fmt.Println("Error exporting tasks:", err)
		return
	}

	fmt.Println("Tasks exported successfully!")
}
func exportCSV() {
	file, err := os.Create("tasks.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("ID,Title,Category,Status,DueDate\n")

	for _, task := range tasks {

		if task.Deleted {
			continue
		}

		status := "Pending"
		if task.Completed {
			status = "Completed"
		}

		file.WriteString(
			fmt.Sprintf(
				"%d,%s,%s,%s,%s\n",
				task.ID,
				task.Title,
				task.Category,
				status,
				task.DueDate.Format("2006-01-02"),
			),
		)
	}

	fmt.Println("CSV exported successfully!")
}

