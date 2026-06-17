package main
import ( "fmt"
"encoding/json"
"os"
"bufio"
"strings"

)
var tasks []Task

type Task struct {
	ID int
	Title string
	Completed bool

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

func main(){
	loadTasks()
	scanner := bufio.NewScanner(os.Stdin)

 for {
	fmt.Println("\nTodo App")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Complete Tasks")
	fmt.Println("4. Delete task")
	fmt.Println("5. Edit Task")
	fmt.Println("6. Search Task")
    fmt.Println("7. Exit")

    fmt.Print("choose: ")
    scanner.Scan()
	var choice int
	fmt.Sscanf(scanner.Text(), "%d", &choice)

	switch choice {
	case 1:
	    fmt.Print("Task title: ")
		scanner.Scan()
		title := strings.TrimSpace(scanner.Text())

		addTask(title)
		saveTasks()

	case 2:
		listTasks()

	case 3:
		var id int
		fmt.Print("Task ID: ")
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &id)

		completeTask(id)
		saveTasks()

	case 4:
		var id int
		fmt.Print("Task ID: ")
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &id)

		deleteTask(id)
		saveTasks()

	case 5:
		var id int
		fmt.Print("Task ID: ")
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &id)

		fmt.Print("New task title: ")
		scanner.Scan()
		newTitle := strings.TrimSpace(scanner.Text())

		if editTask(id, newTitle){
        saveTasks()
		fmt.Println("Task Updated")
		} else {
			fmt.Println("Task not found")
		}


	case 6:
		fmt.Println("Search Keyword: ")
		scanner.Scan()
		keyword:= strings.TrimSpace(scanner.Text())
		searchTask(keyword)

	case 7:
		fmt.Println("See you Later")
		return
	}
  }
}

func addTask(title string){
	task := Task{
		// ID: len(tasks) + 1,
		ID:getNextID(),
		Title: title,
		Completed: false,
	}
	tasks = append(tasks, task)
}
func editTask(id int, newTitle string) bool{
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = newTitle
			return true
		}
// 		if editTask(id, newTitle) {
// 	saveTasks()
// 	fmt.Println("Task updated")
// } else {
// 	fmt.Println("Task not found")
// }
	}
	return false
}
func listTasks(){
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("[%s] %d - %s\n",
	status,
    task.ID,
    task.Title)
	}
}
func saveTasks(){
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	err = os.WriteFile("tasks.json", data,0644)
	if err != nil {
		fmt.Println("Error writing file:",err )
	}
	// os.WriteFile("tasks.json", data, 0644)
}
func completeTask(id int) bool {
	for i:= range tasks {
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
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
func loadTasks(){
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, &tasks)
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

